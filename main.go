package ginside

import (
	"net/http"
	"strconv"
	"time"

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

// BoardRecommendedPosts returns the posts from the first page of a dcgall board
func BoardRecommendedPosts(id string) (posts []Post, err error) {
	// setup http request
	client := &http.Client{
		Timeout: time.Second * 30,
	}
	req, err := http.NewRequest("GET", boardRecommendedPath(id, 1), nil)
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
		noticeID := entryNode.Find("td.t_notice").Text()
		title := entryNode.Find("td.t_subject a").Text()
		link, ok := entryNode.Find("td.t_subject a").Attr("href")
		if !ok {
			return nil, errors.New("unable to find link")
		}
		author, ok := entryNode.Find("td.t_writer").Attr("user_name")
		if !ok {
			return nil, errors.New("unable to find author")
		}
		var date time.Time
		dateText, ok := entryNode.Find("td.t_date").Attr("title")
		if ok {
			date, err = time.Parse(dateFormat, dateText)
			if err != nil {
				return nil, err
			}
		} else {
			dateText := entryNode.Find("td.t_date").Text()
			date, err = time.Parse(dateFormatShort, dateText)
			if err != nil {
				return nil, err
			}
		}
		if date.IsZero() {
			return nil, errors.New("unable to find date")
		}
		hitsText := entryNode.Find("td.t_hits").Eq(0).Text()
		hits, err := strconv.Atoi(hitsText)
		if err != nil {
			return nil, err
		}
		votesText := entryNode.Find("td.t_hits").Eq(1).Text()
		votes, err := strconv.Atoi(votesText)
		if err != nil {
			return nil, err
		}

		// skip announcements
		if noticeID == "공지" {
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
			URL:    baseURL + link,
		})
	}

	return posts, nil
}
