package ginside

import (
	"context"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func TestBoardRecommendedPosts(t *testing.T) {
	g := NewGInside(&http.Client{
		Timeout: 1 * time.Minute,
	})

	posts, err := g.BoardPosts(context.Background(), "pledis", true)
	if err != nil {
		t.Fatal("BoardPosts returned an error: " + err.Error())
	}

	if len(posts) < 10 {
		t.Fatal("BoardPosts returned too few posts: ", strconv.Itoa(len(posts)))
	}
}

func TestBoardMinorRecommendedPosts(t *testing.T) {
	g := NewGInside(&http.Client{
		Timeout: 1 * time.Minute,
	})

	posts, err := g.BoardMinorPosts(context.Background(), "sis", true)
	if err != nil {
		t.Fatal("BoardMinorPosts returned an error: " + err.Error())
	}

	if len(posts) < 10 {
		t.Fatal("BoardMinorPosts returned too few posts: ", strconv.Itoa(len(posts)))
	}
}

func TestBoardAllPosts(t *testing.T) {
	g := NewGInside(&http.Client{
		Timeout: 1 * time.Minute,
	})

	posts, err := g.BoardPosts(context.Background(), "pledis", false)
	if err != nil {
		t.Fatal("BoardPosts returned an error: " + err.Error())
	}

	if len(posts) < 10 {
		t.Fatal("BoardPosts returned too few posts: ", strconv.Itoa(len(posts)))
	}
}

func TestBoardMinorAllPosts(t *testing.T) {
	g := NewGInside(&http.Client{
		Timeout: 1 * time.Minute,
	})

	posts, err := g.BoardMinorPosts(context.Background(), "sis", false)
	if err != nil {
		t.Fatal("BoardMinorPosts returned an error: " + err.Error())
	}

	if len(posts) < 10 {
		t.Fatal("BoardMinorPosts returned too few posts: ", strconv.Itoa(len(posts)))
	}
}
