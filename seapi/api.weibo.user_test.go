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
		User: "杨幂",
		Host: defaultWBHost,
	}
	client := setransport.NewClient(req.Host, req.User)
	response, e := req.Do(context.TODO(), client)
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content), e)
}
