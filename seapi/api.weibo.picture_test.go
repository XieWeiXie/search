package seapi

import (
	"context"
	"fmt"
	"github.com/wuxiaoxiaoshen/search/setransport"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func TestWeiBoPicture(t *testing.T) {
	req := &WeiBoPictureRequest{
		Query: "杨幂",
		Page:  1,
	}
	client := setransport.NewClient(req.Query)
	response, e := req.Do(context.TODO(), client)
	log.Println(response.StatusCode)
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content), e)
}

func TestWeiBoPicture2(t *testing.T) {
	url := "https://s.weibo.com/ajax_pic/list?q=%E6%9D%A8%E5%B9%82&page=20"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Referer", "https://s.weibo.com/pic?q=%E6%9D%A8%E5%B9%82&Refer=Spic_box")

	res, _ := http.DefaultClient.Do(req)
	fmt.Println(res.StatusCode)
}
