package ginside

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

// Post contains information about a single dcinside post
type Post struct {
	ID      string
	Subject string
	Title   string
	Author  string
	Date    time.Time
	Hits    int
	Votes   int
	URL     string
}

// BoardMinorPosts returns the posts from the first page of a dcgall minor board
func (g *GInside) BoardMinorPosts(ctx context.Context, id string, recommended bool) (posts []Post, err error) {
	return boardPostsWithPath(
		ctx, g.httpClient, boardMinorPath(id, 1, recommended),
	)
}

// BoardPosts returns the posts from the first page of a dcgall board
func (g *GInside) BoardPosts(ctx context.Context, id string, recommended bool) (posts []Post, err error) {
	return boardPostsWithPath(
		ctx, g.httpClient, boardPath(id, 1, recommended),
	)
}

// boardPostsWithPath returns the posts from the first page of a dcgall board  at the given path
func boardPostsWithPath(ctx context.Context, client *http.Client, path string) (posts []Post, err error) {
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Referer", headerReferer)
	req.Header.Set("User-Agent", randomUserAgent())
	req.Header.Set("Accept", headerAccept)
	req = req.WithContext(ctx)

	// do http request
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close() // nolint: errcheck
	if res.StatusCode != 200 {
		return nil, errors.New("unexpected status code: " + strconv.Itoa(res.StatusCode))
	}

	// parse html
	mainDoc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	mainDocHTML, err := mainDoc.Html()
	if err != nil {
		return nil, err
	}

	// follow JavaScript forwarding if necessary
	parts := forwardRegex.FindStringSubmatch(mainDocHTML)
	if len(parts) >= 2 {
		return boardPostsWithPath(ctx, client, parts[1])
	}

	// parse posts
	entries := mainDoc.Find(".gall_list tbody tr")

	// parse post content
	for _, entry := range entries.Nodes {
		entryNode := goquery.NewDocumentFromNode(entry)

		subjectHTML, err := entryNode.Find("td.gall_subject").Html()
		if err != nil {
			return nil, err
		}

		noticeID := entryNode.Find("td.gall_num").Text()
		subject := entryNode.Find("td.gall_subject").Text()
		title := entryNode.Find("td.gall_tit a").First().Text()

		// skip pinned
		if strings.Contains(subjectHTML, "<b>") {
			continue
		}
		_, err = strconv.Atoi(noticeID)
		if err != nil {
			continue
		}

		link, ok := entryNode.Find("td.gall_tit a").Attr("href")
		if !ok {
			return nil, errors.New("unable to find link")
		}
		if !strings.HasPrefix(link, "http") {
			link = baseURL + link
		}
		if strings.Contains(link, "javascript:;") {
			continue
		}
		parsedLink, err := url.Parse(link)
		if err != nil {
			return nil, err
		}
		// remove page and exception_mode from final url
		newQueries := parsedLink.Query()
		newQueries.Del("page")
		newQueries.Del("exception_mode")
		parsedLink.RawQuery = newQueries.Encode()

		author := entryNode.Find("td.gall_writer .nickname").Text()
		var date time.Time
		dateText, ok := entryNode.Find("td.gall_date").Attr("title")
		if ok && dateText != "" {
			date, err = time.ParseInLocation(dateFormat, dateText, dateLocation)
			if err != nil {
				date, err = time.ParseInLocation(dateFormatAlt, dateText, dateLocation)
				if err != nil {
					date, err = time.ParseInLocation(dateFormatAlt2, dateText, dateLocation)
					if err != nil {
						return nil, err
					}
				}
			}
		} else {
			dateText := entryNode.Find("td.gall_date").Text()
			if dateText != "" {
				date, err = time.ParseInLocation(dateFormatShort, dateText, dateLocation)
				if err != nil {
					date, err = time.ParseInLocation(dateFormatShortAlt, dateText, dateLocation)
					if err != nil {
						date, err = time.ParseInLocation(dateFormatShortAlt2, dateText, dateLocation)
						if err != nil {
							date, err = time.ParseInLocation(dateFormatShortAlt3, dateText, dateLocation)
							if err != nil {
								return nil, err
							}
						}
					}
				}
			}
		}
		hitsText := entryNode.Find("td.gall_count").Eq(0).Text()
		hits, _ := strconv.Atoi(hitsText)
		votesText := entryNode.Find("td.gall_recommend").Eq(1).Text()
		votes, _ := strconv.Atoi(votesText)

		// add to list of posts
		posts = append(posts, Post{
			ID:      noticeID,
			Subject: subject,
			Title:   title,
			Author:  author,
			Date:    date,
			Hits:    hits,
			Votes:   votes,
			URL:     parsedLink.String(),
		})
	}

	return posts, nil
}
