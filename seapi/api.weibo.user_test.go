package seapi

import (
	"context"
	"fmt"
	"github.com/wuxiaoxiaoshen/search/setransport"
	"io/ioutil"
	"testing"
)

func TestWeiBoUser(t *testing.T) {
	req := WeiBoUserRequest{
		Query: "杨幂",
	}
	client := setransport.NewClient(req.Query)
	response, e := req.Do(context.TODO(), client)
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content), e)
}
