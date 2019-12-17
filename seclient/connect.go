package seclient

import (
	"github.com/wuxiaoxiaoshen/search/seapi"
	"net/http"
)

var (
	defaultZhiHu  = ""
	defaultWeiBo  = ""
	defaultWeChat = ""
)

type ClientConfig struct {
	Name string `json:"name"`
}

type Client struct {
	*seapi.API
	Transport seapi.Transport
}

func NewClient(cfg ClientConfig) *Client {
	return &Client{}
}
func (C Client) Perform(request *http.Request) (*http.Response, error) {
	return C.Transport.Perform(request)
}
