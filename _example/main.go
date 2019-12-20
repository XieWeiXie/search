package main

import (
	"context"
	"github.com/wuxiaoxiaoshen/search"
	"github.com/wuxiaoxiaoshen/search/seapi"
	"github.com/wuxiaoxiaoshen/search/setransport"
	"log"
)

func main() {
	for _, i := range []string{"杨幂", "刘诗诗", "刘亦菲"} {
		Demo1ZhiHuTopic(i)
		Demo2ZhiHuTopic(i)
		Demo3WeiBoUser(i)
		Demo4WeiBoUser(i)
	}

}

func Demo1ZhiHuTopic(query string) {
	req := seapi.ZhiHuTopicRequest{
		Query: query,
	}
	client := setransport.NewClient(req.Query)
	response, _ := req.Do(context.TODO(), client)
	log.Println(response.String())
}

func Demo2ZhiHuTopic(query string) {
	cfg := search.NewClientConfig(query)
	client := search.NewClient(*cfg)
	response, _ := client.ZhiHu.Topic(cfg.Query)
	log.Println(response.String())
}

func Demo3WeiBoUser(query string) {
	req := seapi.WeiBoUserRequest{
		Query: query,
	}
	client := setransport.NewClient(req.Query)
	response, _ := req.Do(context.TODO(), client)
	log.Println(response.String())
}
func Demo4WeiBoUser(query string) {
	cfg := search.NewClientConfig(query)
	client := search.NewClient(*cfg)
	response, _ := client.WeiBo.User(cfg.Query)
	log.Println(response.String())
}
