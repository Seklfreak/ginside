package ginside

import (
	"math/rand"
	"strconv"
	"time"
)

var (
	// the gall base url
	baseURL = "http://gall.dcinside.com"
	// shows recommended posts for a board
	boardPath = func(id string, page int, recommended bool) string {
		path := baseURL + "/board/lists/?id=" + id + "&page=" + strconv.Itoa(page)
		if recommended {
			path += "&exception_mode=recommend"
		}
		return path
	}
	// shows recommended posts for a minor board
	boardMinorPath = func(id string, page int, recommended bool) string {
		path := baseURL + "/mgallery/board/lists/?id=" + id + "&page=" + strconv.Itoa(page)
		if recommended {
			path += "&exception_mode=recommend"
		}
		return path
	}
	// the format used by dcinside
	dateFormat         = "2006.01.02 15:04:05"
	dateFormatAlt      = "2006/01/02 15:04:05"
	dateFormatShort    = "2006.01.02"
	dateFormatShortAlt = "2006/01/02"
	dateLocation, _    = time.LoadLocation("Asia/Seoul")

	headerReferer = "http://gall.dcinside.com/"
	headerAccept  = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8"

	// list of random user agents
	userAgents = []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:56.0) Gecko/20100101 Firefox/56.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
	}

	// returns a random user agent from the list of user agents
	randomUserAgent = func() string {
		rand.Seed(time.Now().Unix())
		return userAgents[rand.Intn(len(userAgents))]
	}
)
