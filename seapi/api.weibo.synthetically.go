package seapi

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func newWeiBoSynthetically() func(request WeiBoSyntheticallyRequest) {
	return func(request WeiBoSyntheticallyRequest) {
		request.formatUrl = "https://s.weibo.com/weibo"
	}
}

type WeiBoSynthetically func(o ...func(request WeiBoSyntheticallyRequest))

type WeiBoSyntheticallyRequest struct {
	Url       url.Values  // 为了构造请求参数
	header    http.Header // 为了头部信息
	formatUrl string      // 完整路径
	ctx       context.Context
}

func (W *WeiBoSyntheticallyRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
	)
	method = http.MethodGet

	if W.Url == nil {
		log.Println("url values empty")
		return nil, errors.New("url values should not be empty")
	}
	for k, v := range W.Url {
		path.WriteString(fmt.Sprintf("%s=%s", k, v))
	}
	path.WriteString("&Refer=SWeibo_box")

	W.formatUrl += path.String()

	req, e := http.NewRequest(method, W.formatUrl, nil)
	if e != nil {
		log.Println("http NewRequest fail ", e.Error())
		return nil, e
	}
	response, e := transport.Perform(req)
	if e != nil {
		log.Println("transport Perform fail ", e.Error())
		return nil, e
	}
	newResponse := Response{
		Body:       response.Body,
		Header:     response.Header,
		StatusCode: response.StatusCode,
	}
	return &newResponse, nil

}

func (W *WeiBoSyntheticallyRequest) WithUrlParams(key, value string) func(*WeiBoSyntheticallyRequest) {
	return func(request *WeiBoSyntheticallyRequest) {
		request.Url.Add(key, value)
	}
}

func (W *WeiBoSyntheticallyRequest) WithHeaderParams(key, value string) func(request *WeiBoSyntheticallyRequest) {
	return func(request *WeiBoSyntheticallyRequest) {
		request.header.Add(key, value)
	}
}
func (W *WeiBoSyntheticallyRequest) WithContext(key, value string) func(request *WeiBoSyntheticallyRequest) {
	return func(request *WeiBoSyntheticallyRequest) {
		request.ctx = context.WithValue(W.ctx, key, value)
	}
}
