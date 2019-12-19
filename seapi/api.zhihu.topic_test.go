package seapi

import (
	"context"
	"fmt"
	"github.com/wuxiaoxiaoshen/search/setransport"
	"io/ioutil"
	"testing"
)

func TestZhiHuTopic(t *testing.T) {
	req := ZhiHuTopicRequest{
		Query: "杨幂",
	}
	client := setransport.NewClient(req.Query)
	response, _ := req.Do(context.TODO(), client)
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
