package seapi

import (
	"context"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func newWeiBoTopic(t Transport) WeiBoTopic {
	return func(query string, o ...func(request *WeiBoTopicRequest)) (response *Response, err error) {
		r := &WeiBoTopicRequest{Host: defaultWBHost, Query: query}
		for _, f := range o {
			f(r)
		}
		return r.Do(context.TODO(), t)
	}
}

type WeiBoTopic func(query string, o ...func(request *WeiBoTopicRequest)) (*Response, error)

type WeiBoTopicRequest struct {
	Query  string
	Host   string
	format string
}

func (W *WeiBoTopic) WithQuery(query string) func(*WeiBoTopicRequest) {
	return func(request *WeiBoTopicRequest) {
		request.Query = query
	}
}

func (W *WeiBoTopicRequest) formatUrl(path string) {
	W.format = fmt.Sprintf("%s/%s", W.Host, path)
}
func (W *WeiBoTopicRequest) parse(response *http.Response) ([]WeiBoTopicResult, error) {
	doc, _ := goquery.NewDocumentFromReader(response.Body)
	var results []WeiBoTopicResult
	doc.Find(".info").Each(func(i int, selection *goquery.Selection) {
		var result WeiBoTopicResult
		href, _ := selection.Find("div a").Attr("href")
		result.Link = strings.TrimSpace(href)
		result.Topic = strings.TrimSpace(selection.Find("div a").Text())
		result.Description = selection.Find("p").Eq(0).Text()
		numbers := selection.Find("p").Last().Text()
		numberList := strings.Split(numbers, " ")
		result.Discussion = numberList[0]
		result.Read = numberList[1]
		results = append(results, result)
	})
	return results, nil
}
func (W *WeiBoTopicRequest) Do(ctx context.Context, t Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
	)
	method = http.MethodGet
	path.WriteString("topic")
	W.formatUrl(path.String())
	log.Println("topic url", W.format)
	u, _ := url.Parse(W.format)
	query := u.Query()
	query.Set(defaultQueryKey, W.Query)
	query.Set(defaultReferKey, defaultReferTopic)
	query.Set("pagetype", "topic")
	query.Set("topic", "1")
	u.RawQuery = query.Encode()

	req, _ := http.NewRequest(method, u.String(), nil)
	req.Header.Add("User-Aget", defaultWeiBoUserAgent)
	req.Header.Add("Host", defaultWeiBoHost)
	response, e := t.Perform(req)
	if e != nil {
		log.Println("transport perform fail", e.Error())
		return nil, errors.New("transport perform fail")
	}
	var results []WeiBoTopicResult
	results, e = W.parse(response)
	if e != nil {
		log.Println("weibo topic parse fail")
		return nil, errors.New("weibo topic parse fail")
	}
	newResponse := newResponse(results, response)
	return newResponse, nil
}

type WeiBoTopicResult struct {
	Topic       string
	Link        string
	Description string
	Discussion  string
	Read        string
}

type WeiBoTopicResults []WeiBoTopicResult
