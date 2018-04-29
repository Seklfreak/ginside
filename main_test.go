package ginside

import (
	"strconv"
	"testing"
)

func TestBoardRecommendedPosts(t *testing.T) {
	posts, err := BoardRecommendedPosts("stone")
	if err != nil {
		t.Fatal("BoardRecommendedPosts returned an error: " + err.Error())
	}

	if len(posts) < 10 {
		t.Fatal("BoardRecommendedPosts returned too few posts: ", strconv.Itoa(len(posts)))
	}
}
