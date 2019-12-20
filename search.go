package search

import (
	"github.com/wuxiaoxiaoshen/search/seapi"
	"github.com/wuxiaoxiaoshen/search/setransport"
	"net/http"
)

type ClientConfig struct {
	Query string `json:"query"`
}

func NewClientConfig(query string) *ClientConfig {
	return &ClientConfig{
		Query: query,
	}
}

type Client struct {
	*seapi.API
	Transport seapi.Transport
}

func NewClient(cfg ClientConfig) *Client {
	transport := setransport.NewClient(cfg.Query)
	return &Client{
		Transport: transport,
		API:       seapi.New(transport),
	}
}
func (C Client) Perform(request *http.Request) (*http.Response, error) {
	return C.Transport.Perform(request)
}
