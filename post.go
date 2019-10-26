package ginside

import (
	"context"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type PostDetails struct {
	ID          string
	Title       string
	Author      string
	Date        time.Time
	URL         string
	ContentHTML string
	Attachments []PostAttachment
}

type PostAttachment struct {
	URL      string
	Filename string
}

func (g *GInside) PostDetails(ctx context.Context, link string) (*PostDetails, error) {
	resp, err := makeRequest(ctx, g.httpClient, link)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // nolint: errcheck

	// parse html
	mainDoc, err := goquery.NewDocumentFromReader(resp.Body)
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
		return g.PostDetails(ctx, parts[1])
	}

	var details PostDetails

	details.ID, _ = mainDoc.Find("input[name=gallery_no]").First().Attr("value")
	details.Title = mainDoc.Find(".title_subject").First().Text()
	details.Author, _ = mainDoc.Find(".gall_writer").First().Attr("data-nick")
	details.Date, _ = parseDate(mainDoc.Find(".gall_date").First().Text())
	details.URL = link

	postHTML, _ := mainDoc.Find(".writing_view_box").First().Html()
	postHTML = strings.Replace(strings.Replace(postHTML, "\n", "", -1), "\t", "", -1)
	details.ContentHTML = postHTML

	attachments := mainDoc.Find("ul.appending_file li")
	var attachment PostAttachment
	attachments.Each(func(_ int, selection *goquery.Selection) {
		attachment = PostAttachment{}
		attachment.URL, _ = selection.Find("a").Attr("href")
		attachment.Filename = selection.Find("a").Text()

		if attachment.URL != "" {
			details.Attachments = append(details.Attachments, attachment)
		}
	})

	return &details, nil
}
