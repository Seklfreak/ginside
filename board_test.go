package ginside

import (
	"strconv"
	"testing"
)

func TestBoardRecommendedPosts(t *testing.T) {
	posts, err := BoardRecommendedPosts("pledis")
	if err != nil {
		t.Fatal("BoardRecommendedPosts returned an error: " + err.Error())
	}

	if len(posts) < 10 {
		t.Fatal("BoardRecommendedPosts returned too few posts: ", strconv.Itoa(len(posts)))
	}
}

func TestBoardMinorRecommendedPosts(t *testing.T) {
	posts, err := BoardMinorRecommendedPosts("sis")
	if err != nil {
		t.Fatal("BoardRecommendedPosts returned an error: " + err.Error())
	}

	if len(posts) < 10 {
		t.Fatal("BoardRecommendedPosts returned too few posts: ", strconv.Itoa(len(posts)))
	}
}

func TestBoardAllPosts(t *testing.T) {
	posts, err := BoardAllPosts("pledis")
	if err != nil {
		t.Fatal("BoardAllPosts returned an error: " + err.Error())
	}

	if len(posts) < 10 {
		t.Fatal("BoardAllPosts returned too few posts: ", strconv.Itoa(len(posts)))
	}
}

func TestBoardMinorAllPosts(t *testing.T) {
	posts, err := BoardMinorAllPosts("sis")
	if err != nil {
		t.Fatal("BoardMinorAllPosts returned an error: " + err.Error())
	}

	if len(posts) < 10 {
		t.Fatal("BoardMinorAllPosts returned too few posts: ", strconv.Itoa(len(posts)))
	}
}
