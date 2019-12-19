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

func newWeiBoUser(t Transport) WeiBoUser {
	return func(name string, o ...func(*WeiBoUserRequest)) (response *Response, err error) {
		var r = &WeiBoUserRequest{
			Query: name,
			host:  defaultWBHost,
		}
		for _, f := range o {
			f(r)
		}
		return r.Do(context.TODO(), t)
	}
}

type WeiBoUser func(query string, o ...func(*WeiBoUserRequest)) (*Response, error)

type WeiBoUserRequest struct {
	Query  string
	host   string
	format string
}

func (W *WeiBoUser) WithUser(name string) func(request *WeiBoUserRequest) {
	return func(request *WeiBoUserRequest) {
		request.Query = name
	}
}

func (W *WeiBoUserRequest) formatUrl(path string) {
	W.host = defaultWBHost
	W.format = fmt.Sprintf("%s/%s", W.host, path)
}

func (W *WeiBoUserRequest) parse(response *http.Response) []WeiBoUserResponse {
	doc, e := goquery.NewDocumentFromReader(response.Body)
	if e != nil {
		log.Println("goquery new document fail", e.Error())
		return nil
	}
	var results []WeiBoUserResponse
	doc.Find(".info").Each(func(i int, selection *goquery.Selection) {
		var result WeiBoUserResponse
		href, _ := selection.Find("div").First().Find("a").First().Attr("href")
		result.Blog = fmt.Sprintf("https:%s", strings.TrimSpace(href))
		result.Name = strings.TrimSpace(selection.Find("div").First().Find("a").First().Text())
		local := strings.TrimSpace(selection.Find("p").Eq(0).Text())
		replacer := strings.NewReplacer("\n", "", "个人主页", "", " ", "")
		result.Local = replacer.Replace(local)
		result.Description = strings.TrimSpace(selection.Find("p:contains(简介)").Text())
		if selection.Find("p").Find("span").Size() == 3 {
			numbers := selection.Find("p")
			result.Following = numbers.Find("span").Eq(0).Text()
			result.Follower = numbers.Find("span").Eq(1).Text()
			result.PublishNumber = numbers.Find("span").Eq(2).Text()

		}
		if selection.Find("p").Eq(1).Find("span").Size() == 0 {
			result.Content = strings.TrimSpace(selection.Find("p").Eq(1).Text())
		}
		tags := selection.Find("p:contains(标签)")
		tags.Find("a").Each(func(i int, selection *goquery.Selection) {
			result.Tags = append(result.Tags, strings.TrimSpace(selection.Text()))
		})
		result.Jobs = selection.Find("p:contains(职业信息)").Find("a").Text()
		//log.Println(fmt.Sprintf("%#v", result))
		results = append(results, result)
	})
	return results
}

func (W *WeiBoUserRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
	)
	method = http.MethodGet
	path.WriteString("user")

	W.formatUrl(path.String())
	log.Println("user url", W.format)
	u, e := url.Parse(W.format)
	if e != nil {
		log.Println(e)
		return nil, e
	}
	query := u.Query()
	query.Set(defaultQueryKey, W.Query)
	query.Set(defaultReferKey, defaultReferUser)
	u.RawQuery = query.Encode()
	req, e := http.NewRequest(method, u.String(), nil)
	if e != nil {
		log.Println(e)
		return nil, e
	}
	req.Host = defaultWeiBoHost
	req.Header.Add("User-Agent", defaultWeiBoUserAgent)
	response, e := transport.Perform(req)
	if e != nil {
		log.Println(e)
		return nil, e
	}
	results := W.parse(response)
	if len(results) == 0 {
		log.Println("No Result")
		return nil, errors.New("no result")
	}
	newResponse := newResponse(results, response)
	return newResponse, nil

}

type WeiBoUserResponse struct {
	Blog          string
	Name          string
	Local         string
	Description   string
	Content       string
	Follower      string   // 粉丝
	Following     string   // 关注
	PublishNumber string   // 微博数目
	Tags          []string // 标签
	Jobs          string   // 职业信息
}

type WeiBoUserResponses []WeiBoUserResponse
