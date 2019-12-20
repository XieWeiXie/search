package search

import (
	"log"
	"testing"
)

func TestClientUser(t *testing.T) {
	cfg := NewClientConfig("杨幂")
	client := NewClient(*cfg)
	response, _ := client.WeiBo.User(cfg.Query)
	log.Println(response.String())

}
func TestClientTopic(t *testing.T) {
	cfg := NewClientConfig("杨幂")
	client := NewClient(*cfg)
	response, _ := client.WeiBo.Topic(cfg.Query)
	log.Println(response.String())
}

func TestClientPicture(t *testing.T) {
	cfg := NewClientConfig("杨幂")
	client := NewClient(*cfg)
	response, _ := client.WeiBo.Picture(cfg.Query, 1)
	log.Println(response.String())
}

func TestClientPassage(t *testing.T) {
	cfg := NewClientConfig("杨幂")
	client := NewClient(*cfg)
	response, _ := client.WeiBo.Passage(cfg.Query)
	log.Println(response.String())
}

func TestClientVideo(t *testing.T) {
	cfg := NewClientConfig("杨幂")
	client := NewClient(*cfg)
	response, _ := client.WeiBo.Video(cfg.Query)
	log.Println(response.String())
}

func TestClientZhiHuTopic(t *testing.T) {
	cfg := NewClientConfig("杨幂")
	client := NewClient(*cfg)
	response, _ := client.ZhiHu.Topic(cfg.Query)
	log.Println(response.String())
}
