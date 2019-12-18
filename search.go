package search

import (
	"github.com/wuxiaoxiaoshen/search/seapi"
	"github.com/wuxiaoxiaoshen/search/setransport"
	"net/http"
)

var (
	defaultZhiHu  = ""
	defaultWeiBo  = "https://s.weibo.com"
	defaultWeChat = ""
)

type ClientConfig struct {
	Host  string `json:"name"`
	Query string `json:"query"`
}

func NewClientConfig(host string, query string) *ClientConfig {
	return &ClientConfig{
		Host:  host,
		Query: query,
	}
}

type Client struct {
	*seapi.API
	Transport seapi.Transport
}

func NewClient(cfg ClientConfig) *Client {
	transport := setransport.NewClient(cfg.Host, cfg.Query)
	return &Client{
		Transport: transport,
		API:       seapi.New(transport),
	}
}
func (C Client) Perform(request *http.Request) (*http.Response, error) {
	return C.Transport.Perform(request)
}
