package setransport

import "net/http"

type Client struct {
	Host      string `json:"host"`
	Query     string `json:"query"`
	transport *CacheTransport
}

func NewClient(host string, query string) *Client {
	return &Client{
		Host:      host,
		Query:     query,
		transport: newCacheTransport(),
	}
}

func (C *Client) newRequest() *http.Request {
	return nil
}

func (C *Client) Perform(r *http.Request) (*http.Response, error) {
	// 真实的数据处理，实现了 Transport 接口
	// 更底层的网络请求由 transport 来实现，先读取缓存，否则发起请求
	return C.transport.RoundTrip(r)
}
