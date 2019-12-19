package seapi

import (
	"context"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
)

func newZhiHuTopic(t Transport) ZhiHuTopic {
	return func(Query string, o ...func(request *ZhiHuTopicRequest)) (response *Response, err error) {
		r := &ZhiHuTopicRequest{
			Query: Query,
			host:  defaultZhiHuHost,
		}
		for _, f := range o {
			f(r)
		}
		return r.Do(context.TODO(), t)
	}
}

type ZhiHuTopic func(Query string, o ...func(request *ZhiHuTopicRequest)) (*Response, error)

type ZhiHuTopicRequest struct {
	Query  string
	host   string
	format string
}

func (Z *ZhiHuTopicRequest) formatUrl(path string) {
	Z.host = defaultZhiHuHost
	Z.format = fmt.Sprintf("%s/%s", Z.host, path)
}
func (Z *ZhiHuTopicRequest) parse(response *http.Response) ([]ZhiHuTopicResult, error) {
	doc, _ := goquery.NewDocumentFromReader(response.Body)
	var results []ZhiHuTopicResult
	doc.Find(".ContentItem-head").Each(func(i int, selection *goquery.Selection) {
		var result ZhiHuTopicResult
		title := selection.Find(".ContentItem-title")
		result.Title = title.Find("div a").Text()
		link, _ := title.Find("div a").Attr("href")
		result.Link = fmt.Sprintf("https:%s", link)

		meta := selection.Find(".ContentItem-meta")
		result.Text = meta.Find("div div").First().Text()
		numbers := meta.Find("div div").Last().Find("a")
		result.Concern.Number = numbers.Eq(0).Text()
		subLink0, _ := numbers.Eq(0).Attr("href")
		result.Concern.SubLink = fmt.Sprintf(defaultZhiHuHost + subLink0)
		result.Question.Number = numbers.Eq(1).Text()
		subLink1, _ := numbers.Eq(1).Attr("href")
		result.Concern.SubLink = fmt.Sprintf(defaultZhiHuHost + subLink1)
		result.Hot.Number = numbers.Last().Text()
		subLink2, _ := numbers.Last().Attr("href")
		result.Concern.SubLink = fmt.Sprintf(defaultZhiHuHost + subLink2)

		results = append(results, result)
	})
	return results, nil
}
func (Z *ZhiHuTopicRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
	)
	method = http.MethodGet
	path.WriteString("search")
	Z.formatUrl(path.String())

	searchUrl := SearchUrl{Q: Z.format}
	searchUrl.Url()
	searchUrl.Add(defaultZhiHuQueryKey, Z.Query)
	searchUrl.Add(defaultZhiHuTypeKey, defaultZhiHuTypeValue)
	log.Println("ZhiHu Topic Url", searchUrl.Format())
	req, _ := http.NewRequest(method, searchUrl.Format(), nil)
	response, e := transport.Perform(req)
	if e != nil {
		log.Println("transport perform", e.Error())
		return nil, errors.New("transport perform")
	}
	results, e := Z.parse(response)
	if e != nil {
		log.Println("zhihu parse fail")
		return nil, errors.New("zhihu parse fail")
	}
	newRes := newResponse(results, response)
	return newRes, nil

}

type ZhiHuTopicResult struct {
	Title   string
	Text    string
	Link    string
	Concern struct {
		Number  string
		SubLink string
	}
	Question struct {
		Number  string
		SubLink string
	}
	Hot struct {
		Number  string
		SubLink string
	}
}
type ZhiHuTopicResults []ZhiHuTopicRequest
