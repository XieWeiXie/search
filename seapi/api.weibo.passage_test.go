package seapi

import (
	"context"
	"fmt"
	"github.com/wuxiaoxiaoshen/search/setransport"
	"io/ioutil"
	"log"
	"testing"
)

func TestWeiBoPassage(t *testing.T) {
	req := &WeiBoPassageRequest{
		Query: "杨幂",
	}
	client := setransport.NewClient(req.Query)
	response, e := req.Do(context.TODO(), client)
	log.Println(response.StatusCode)
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content), e)
}
