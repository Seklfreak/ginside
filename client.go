package ginside

import (
	"net/http"
)

type GInside struct {
	httpClient *http.Client
}

func NewGInside(httpClient *http.Client) *GInside {
	return &GInside{
		httpClient: httpClient,
	}
}
