package ginside

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

func makeRequest(ctx context.Context, client *http.Client, link string) (*http.Response, error) {
	req, err := http.NewRequest("GET", link, nil)
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

	if res.StatusCode != 200 {
		return nil, errors.New("unexpected status code: " + strconv.Itoa(res.StatusCode))
	}

	return res, nil
}

func parseDate(dateText string) (time.Time, error) {
	date, err := time.ParseInLocation(dateFormat, dateText, dateLocation)
	if err == nil {
		return date, nil
	}

	date, err = time.ParseInLocation(dateFormatAlt, dateText, dateLocation)
	if err == nil {
		return date, nil
	}

	date, err = time.ParseInLocation(dateFormatAlt2, dateText, dateLocation)
	if err == nil {
		return date, nil
	}

	date, err = time.ParseInLocation(dateFormatShort, dateText, dateLocation)
	if err == nil {
		return date, nil
	}

	date, err = time.ParseInLocation(dateFormatShortAlt, dateText, dateLocation)
	if err == nil {
		return date, nil
	}

	date, err = time.ParseInLocation(dateFormatShortAlt2, dateText, dateLocation)
	if err == nil {
		return date, nil
	}

	date, err = time.ParseInLocation(dateFormatShortAlt3, dateText, dateLocation)
	if err == nil {
		return date, nil
	}

	return time.Time{}, fmt.Errorf("unable to parse date: %s", dateText)
}
