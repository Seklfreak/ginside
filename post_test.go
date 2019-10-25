package ginside

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func TestGInside_PostDetails(t *testing.T) {
	g := NewGInside(&http.Client{
		Timeout: 1 * time.Minute,
	})

	postDetails, err := g.PostDetails(context.Background(), "https://gall.dcinside.com/board/view/?id=stone&no=326054")
	if err != nil {
		t.Fatal("BoardPosts returned an error: " + err.Error())
	}

	spew.Dump(postDetails)
}
