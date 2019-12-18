package seapi

import (
	"context"
	"net/http"
)

type SearchApi interface {
	Do(ctx context.Context, transport Transport) (*Response, error)
}

type Transport interface {
	Perform(*http.Request) (*http.Response, error)
}
