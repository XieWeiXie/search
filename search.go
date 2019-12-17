package search

import (
	"context"
	"github.com/wuxiaoxiaoshen/search/seapi"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	url     url.Values
	request *http.Request
}

func (C Client) Do(ctx context.Context, api seapi.SearchApi) {
	api.Do(ctx, C)
}
func (C Client) Perform(r *http.Request) (*http.Response, error) {

}
func (C *Client) newRequest(method string, url url.Values, body io.Reader) *http.Request {
	return C.request
}
