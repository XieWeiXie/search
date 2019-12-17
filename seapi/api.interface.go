package seapi

import (
	"context"
	"io"
	"net/http"
)

type SearchApi interface {
	Do(ctx context.Context, transport Transport) (*Response, error)
	Run(ctx context.Context, transport Transport) (*Response, error)
}

type Response struct {
	StatusCode int
	Header     http.Header
	Body       io.ReadCloser
}

type Transport interface {
	Perform(*http.Request) (*http.Response, error)
}
