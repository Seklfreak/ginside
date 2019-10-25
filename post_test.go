package ginside

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func TestGInside_PostDetails(t *testing.T) {
	g := NewGInside(&http.Client{
		Timeout: 1 * time.Minute,
	})

	postDetails, err := g.PostDetails(context.Background(), "https://gall.dcinside.com/board/view/?id=stone&no=326054")
	if err != nil {
		t.Fatal("PostDetails returned an error: " + err.Error())
	}

	if postDetails.URL == "" {
		t.Fatal("PostDetails URL is empty")
	}
}
