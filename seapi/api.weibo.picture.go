package seapi

import (
	"context"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func newWeiBoPicture(t Transport) WeiBoPicture {
	return func(query string, page int, o ...func(request *WeiBoPictureRequest)) (response *Response, err error) {
		r := &WeiBoPictureRequest{
			Query: query,
			Page:  page,
			host:  defaultWBHost,
		}
		for _, f := range o {
			f(r)
		}
		return r.Do(context.TODO(), t)
	}
}

type WeiBoPicture func(query string, page int, o ...func(request *WeiBoPictureRequest)) (*Response, error)

type WeiBoPictureRequest struct {
	Query  string
	Page   int
	host   string
	format string
}

func (W *WeiBoPictureRequest) formatUrl(path string) {
	W.host = defaultWBHost
	W.format = fmt.Sprintf("%s/%s", W.host, path)
}

func (W *WeiBoPictureRequest) parse(response *http.Response) ([]WeiBoPictureResult, error) {
	content, _ := ioutil.ReadAll(response.Body)
	doc := gjson.ParseBytes(content)
	var results []WeiBoPictureResult
	array := doc.Get("data.pic_list").Array()
	stringsHandler := func(v string) string {
		replacer := strings.NewReplacer("\n", "", "\t", "", " ", "")
		return replacer.Replace(v)
	}
	for _, i := range array {
		var result WeiBoPictureResult
		result = WeiBoPictureResult{
			Picture: fmt.Sprintf("https:%s", i.Get("original_pic").String()),
			Text:    stringsHandler(strings.TrimSpace(i.Get("text").String())),
			SubText: stringsHandler(strings.TrimSpace(i.Get("sub_text").String())),
		}
		log.Println(result)
		results = append(results, result)
	}
	return results, nil

}

func (W *WeiBoPictureRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
	)
	method = http.MethodGet
	path.WriteString("pic")
	W.formatUrl(path.String())

	referUrl := SearchUrl{Q: W.format}
	referUrl.Url()
	referUrl.Add(defaultQueryKey, W.Query)
	referUrl.Add(defaultReferKey, defaultReferPicture)
	log.Println("refer url", referUrl.Format())

	searchUrl := SearchUrl{Q: fmt.Sprintf("https://s.weibo.com/ajax_pic/list")}
	searchUrl.Url()
	searchUrl.Add(defaultQueryKey, W.Query)
	searchUrl.Add(defaultPageKey, strconv.Itoa(W.Page))
	log.Println("search url", searchUrl.Format())

	req, _ := http.NewRequest(method, searchUrl.Format(), nil)
	req.Header.Add(defaultHeaderReferer, referUrl.Format())
	req.Header.Add(defaultHeaderXRequestWith, defaultHeaderXRequestWithValue)
	req.Header.Add("User-Agent", defaultWeiBoUserAgent)
	response, e := transport.Perform(req)
	if e != nil {
		log.Println("transport perform fail", e.Error())
		return nil, e
	}
	results, e := W.parse(response)
	log.Println(results)
	if e != nil {
		log.Println("parse picture result fail,", e.Error())
		return nil, e
	}
	newResponse := newResponse(results, response)
	return newResponse, nil

}

type WeiBoPictureResult struct {
	Picture string
	Text    string
	SubText string
}

type WeiBoPictureResults []WeiBoTopicResult
