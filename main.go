package ginside

import (
	"net/http"
	"strconv"
	"time"

	"strings"

	"net/url"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
)

// Post contains information about a single dcinside post
type Post struct {
	ID     string
	Title  string
	Author string
	Date   time.Time
	Hits   int
	Votes  int
	URL    string
}

func BoardMinorRecommendedPosts(id string) (posts []Post, err error) {
	return boardRecommendedPostsWithPath(boardMinorRecommendedPath(id, 1))
}

// BoardRecommendedPosts returns the posts from the first page of a dcgall board
func BoardRecommendedPosts(id string) (posts []Post, err error) {
	return boardRecommendedPostsWithPath(boardRecommendedPath(id, 1))
}

// BoardRecommendedPosts returns the posts from the first page of a dcgall board
func boardRecommendedPostsWithPath(path string) (posts []Post, err error) {
	// setup http request
	client := &http.Client{
		Timeout: time.Second * 30,
	}
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Referer", headerReferer)
	req.Header.Set("User-Agent", randomUserAgent())
	req.Header.Set("Accept", headerAccept)

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

	// parse posts
	entries := mainDoc.Find(".list_tbody .tb")

	// parse post content
	for _, entry := range entries.Nodes {
		entryNode := goquery.NewDocumentFromNode(entry)
		//fmt.Println(entryNode.Html())
		noticeID := entryNode.Find("td.t_notice").Text()
		title := entryNode.Find("td.t_subject a").Text()
		// remove [n] after title
		if strings.Contains(title, "[") {
			parts := strings.Split(title, "[")
			title = strings.Join(parts[0:len(parts)-1], "")
		}

		link, ok := entryNode.Find("td.t_subject a").Attr("href")
		if !ok {
			return nil, errors.New("unable to find link")
		}
		link = baseURL + link
		parsedLink, err := url.Parse(link)
		if err != nil {
			return nil, err
		}
		// remove page and exception_mode from final url
		newQueries := parsedLink.Query()
		newQueries.Del("page")
		newQueries.Del("exception_mode")
		parsedLink.RawQuery = newQueries.Encode()

		author, ok := entryNode.Find("td.t_writer").Attr("user_name")
		if !ok {
			return nil, errors.New("unable to find author")
		}
		var date time.Time
		dateText, ok := entryNode.Find("td.t_date").Attr("title")
		if ok {
			date, err = time.ParseInLocation(dateFormat, dateText, dateLocation)
			if err != nil {
				return nil, err
			}
		} else {
			dateText := entryNode.Find("td.t_date").Text()
			date, err = time.ParseInLocation(dateFormatShort, dateText, dateLocation)
			if err != nil {
				return nil, err
			}
		}
		if date.IsZero() {
			return nil, errors.New("unable to find date")
		}
		hitsText := entryNode.Find("td.t_hits").Eq(0).Text()
		hits, _ := strconv.Atoi(hitsText)
		votesText := entryNode.Find("td.t_hits").Eq(1).Text()
		votes, _ := strconv.Atoi(votesText)

		// skip announcements and news
		if noticeID == "공지" || noticeID == "뉴스" {
			continue
		}

		// add to list of posts
		posts = append(posts, Post{
			ID:     noticeID,
			Title:  title,
			Author: author,
			Date:   date,
			Hits:   hits,
			Votes:  votes,
			URL:    parsedLink.String(),
		})
	}

	return posts, nil
}
