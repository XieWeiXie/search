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

func newWeiBoVideo(t Transport) WeiBoVideo {
	return func(query string, o ...func(request *WeiBoVideoRequest)) (response *Response, err error) {
		r := &WeiBoVideoRequest{
			Query: query,
			host:  defaultWBHost,
		}
		for _, f := range o {
			f(r)
		}
		return r.Do(context.TODO(), t)
	}
}

type WeiBoVideo func(query string, o ...func(request *WeiBoVideoRequest)) (*Response, error)

type WeiBoVideoRequest struct {
	Query  string
	host   string
	format string
}

func (W *WeiBoVideoRequest) formatUrl(path string) {
	W.host = defaultWBHost
	W.format = fmt.Sprintf("%s/%s", W.host, path)
}

func (W *WeiBoVideoRequest) parse(response *http.Response) ([]WeiBoVideoResult, error) {
	doc, _ := goquery.NewDocumentFromReader(response.Body)
	var results []WeiBoVideoResult

	doc.Find(".card").Each(func(i int, selection *goquery.Selection) {
		var result WeiBoVideoResult
		feed := selection.Find(".card-feed .content")
		info := feed.Find("div.info")
		result.Author = StringHandler{Value: info.Find("div").Eq(1).Text()}.Replacer("\n", "", "\t", "", " ", "")

		txt := feed.Find("p.txt")
		result.Text = StringHandler{Value: txt.Text()}.Replacer("\n", "", "\t", "", " ", "")

		result.Date = StringHandler{Value: feed.Find("p.from").Text()}.Replacer("\n", "", "\t", "", " ", "")
		link, _ := feed.Find("p.from a").Attr("href")
		result.Link = fmt.Sprintf("https:%s", link)
		media, _ := feed.Find("div.thumbnail a").Attr("action-data")
		mediaString := StringHandler{Value: media}
		regexpList := mediaString.Regexp(`full_url=(.*?)$`)
		if len(regexpList) > 0 {
			result.MediaLink = fmt.Sprintf("https:%s", mediaString.Url("video_src", regexpList[0][1]))
		}
		numbers := selection.Find(".card-act ul li")
		result.Transfer = numbers.Eq(1).Text()
		result.Discussion = numbers.Eq(2).Text()
		result.Like = numbers.Last().Text()
		results = append(results, result)
	})
	return results, nil
}

func (W *WeiBoVideoRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
	)
	method = http.MethodGet
	path.WriteString("video")
	W.formatUrl(path.String())

	searchUrl := SearchUrl{Q: W.format}
	searchUrl.Url()
	searchUrl.Add(defaultQueryKey, W.Query)
	searchUrl.Add(defaultReferKey, defaultReferVideo)
	searchUrl.Add("xsort", "hot")
	searchUrl.Add("hasvideo", "1")
	searchUrl.Add("tw", "video")

	req, _ := http.NewRequest(method, searchUrl.Format(), nil)
	response, e := transport.Perform(req)
	if e != nil {
		log.Println("transport perform fail")
		return nil, errors.New("transport perform fail")
	}
	results, e := W.parse(response)
	if e != nil {
		log.Println("weibo video parse fail")
		return nil, errors.New("weibo video parse fail")
	}
	newRes := newResponse(results, response)
	return newRes, nil
}

type WeiBoVideoResult struct {
	Author     string
	Text       string
	Link       string
	MediaLink  string
	Date       string
	Transfer   string
	Discussion string
	Like       string
}

type WeiBoVideoResults []WeiBoVideoResult
