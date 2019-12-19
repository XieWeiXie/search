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

func newWeiBoPassage(t Transport) WeiBoPassage {
	return func(query string, o ...func(request *WeiBoPassageRequest)) (response *Response, err error) {
		r := &WeiBoPassageRequest{
			Query: query,
			host:  defaultWBHost,
		}
		for _, f := range o {
			f(r)
		}
		return r.Do(context.TODO(), t)
	}
}

type WeiBoPassage func(query string, o ...func(request *WeiBoPassageRequest)) (*Response, error)

type WeiBoPassageRequest struct {
	Query  string
	host   string
	format string
}

func (W *WeiBoPassageRequest) formatUrl(path string) {
	W.host = defaultWBHost
	W.format = fmt.Sprintf("%s/%s", W.host, path)
}

func (W *WeiBoPassageRequest) parse(response *http.Response) ([]WeiBoPassageResult, error) {
	doc, _ := goquery.NewDocumentFromReader(response.Body)
	var results []WeiBoPassageResult
	replacer := func(v string) string {
		rp := strings.NewReplacer("\n", "", "\t", "", " ", "")
		return rp.Replace(v)
	}
	doc.Find(".card-article-a").Each(func(i int, selection *goquery.Selection) {
		var result WeiBoPassageResult
		result.Title = strings.TrimSpace(selection.Find("h3").Text())
		href, _ := selection.Find("h3 a").Attr("href")
		result.Link = strings.TrimSpace(href)
		detail := selection.Find("div.detail")
		result.Detail = detail.Find("p.txt").Text()
		numbers := detail.Find("div.act ul.s-fr")
		result.Share = numbers.Find("li").First().Text()
		result.Like = replacer(numbers.Find("li").Last().Text())
		authorMessage := detail.Find("div.act div span")
		result.Date = authorMessage.Last().Text()
		result.Author = authorMessage.First().Find("a").First().Text()
		hrefLink, _ := authorMessage.First().Find("a").First().Attr("href")
		result.AuthorLink = fmt.Sprintf("https:%s", hrefLink)

		results = append(results, result)

	})
	return results, nil
}

func (W *WeiBoPassageRequest) Do(ctx context.Context, t Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
	)
	method = http.MethodGet
	path.WriteString("article")
	W.formatUrl(path.String())
	searchUrl := SearchUrl{Q: W.format}
	searchUrl.Url()
	searchUrl.Add(defaultQueryKey, W.Query)
	searchUrl.Add(defaultReferKey, defaultReferPassage)

	req, _ := http.NewRequest(method, searchUrl.Format(), nil)
	req.Header.Add("User-Agent", defaultWeiBoUserAgent)
	response, e := t.Perform(req)
	if e != nil {
		log.Println("transport perform fail", e.Error())
		return nil, errors.New("transport perform fail")
	}
	results, e := W.parse(response)
	if e != nil {
		log.Println("weibo passage parse fail")
		return nil, errors.New("transport perform fail")
	}
	newRes := newResponse(results, response)
	return newRes, nil

}

type WeiBoPassageResult struct {
	Title      string
	Link       string
	Detail     string
	Author     string
	AuthorLink string
	Date       string
	Share      string
	Like       string
}

type WeiBoPassageResults []WeiBoPassageResult
