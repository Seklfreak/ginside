package ginside

import (
	"strconv"
	"testing"
)

func TestBoardRecommendedPosts(t *testing.T) {
	posts, err := BoardPosts("pledis", true)
	if err != nil {
		t.Fatal("BoardPosts returned an error: " + err.Error())
	}

	if len(posts) < 10 {
		t.Fatal("BoardPosts returned too few posts: ", strconv.Itoa(len(posts)))
	}
}

func TestBoardMinorRecommendedPosts(t *testing.T) {
	posts, err := BoardMinorPosts("sis", true)
	if err != nil {
		t.Fatal("BoardMinorPosts returned an error: " + err.Error())
	}

	if len(posts) < 10 {
		t.Fatal("BoardMinorPosts returned too few posts: ", strconv.Itoa(len(posts)))
	}
}

func TestBoardAllPosts(t *testing.T) {
	posts, err := BoardPosts("pledis", false)
	if err != nil {
		t.Fatal("BoardPosts returned an error: " + err.Error())
	}

	if len(posts) < 10 {
		t.Fatal("BoardPosts returned too few posts: ", strconv.Itoa(len(posts)))
	}
}

func TestBoardMinorAllPosts(t *testing.T) {
	posts, err := BoardMinorPosts("sis", false)
	if err != nil {
		t.Fatal("BoardMinorPosts returned an error: " + err.Error())
	}

	if len(posts) < 10 {
		t.Fatal("BoardMinorPosts returned too few posts: ", strconv.Itoa(len(posts)))
	}
}
