package search

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestClientUser(t *testing.T) {
	cfg := NewClientConfig("https://s.weibo.com", "杨幂")
	client := NewClient(*cfg)
	response, _ := client.WeiBo.User(cfg.Query)
	content, _ := ioutil.ReadAll(response.Body)
	log.Println(string(content))

}
func TestClientTopic(t *testing.T) {
	cfg := NewClientConfig("https://s.weibo.com", "杨幂")
	client := NewClient(*cfg)
	response, _ := client.WeiBo.Topic(cfg.Query)
	content, _ := ioutil.ReadAll(response.Body)
	log.Println(string(content))
}

func TestClientPicture(t *testing.T) {
	cfg := NewClientConfig("https://s.weibo.com", "杨幂")
	client := NewClient(*cfg)
	response, _ := client.WeiBo.Picture(cfg.Query, 1)
	content, _ := ioutil.ReadAll(response.Body)
	log.Println(string(content))
}
