package seapi

import (
	"context"
	"fmt"
	"github.com/wuxiaoxiaoshen/search/setransport"
	"io/ioutil"
	"log"
	"net/url"
	"regexp"
	"testing"
)

func TestRegexp(t *testing.T) {
	content := `type=feedvideo&objectid=1034:4449749244248087&keys=4449749422834445&cover_img=https%3A%2F%2Fwx4.sinaimg.cn%2Fcrop.0.6.1921.1068%2F6ae8240aly1g9xbrdbxj3j21hd0u07wi.jpg&card_height=2136&card_width=3842&play_count=101854306&short_url=http%3A%2F%2Ft.cn%2FAiDKTfDl&encode_mode=&bitrate=277&biz_id=230444&current_mid=4449749995968936&key=tblog_card&title=%E9%9B%85%E8%AF%97%E5%85%B0%E9%BB%9B%E7%9A%84%E5%BE%AE%E5%8D%9A%E8%A7%86%E9%A2%91&full_url=https%3A%2F%2Fvideo.weibo.com%2Fshow%3Ffid%3D1034%3A4449749244248087&object_id=1034:4449749244248087&video_src=%2F%2Ff.video.weibocdn.com%2F004j8SgRlx07znj1zmZO010412008zT30E010.mp4%3Flabel%3Dmp4_ld%26template%3D640x360.25.0%26trans_finger%3D40a32e8439c5409a63ccf853562a60ef%26Expires%3D1576729427%26ssig%3DsotYdsogqZ%26KID%3Dunistore%2Cvideo"`
	reg := func(v string) string {
		pattern := `full_url=(.*?)"`
		regx := regexp.MustCompile(pattern)
		list := regx.FindAllStringSubmatch(v, -1)
		log.Println(list)
		return list[0][1]
	}
	log.Println(reg(content))
	values, _ := url.ParseQuery(reg(content))
	target := values.Get("video_src")
	log.Println(fmt.Sprintf("https:%s", target))
	//https://f.video.weibocdn.com/004j8SgRlx07znj1zmZO010412008zT30E010.mp4?label=mp4_ld&template=640x360.25.0&trans_finger=40a32e8439c5409a63ccf853562a60ef&Expires=1576729427&ssig=sotYdsogqZ&KID=unistore,video
	//https://f.video.weibocdn.com/004j8SgRlx07znj1zmZO010412008zT30E010.mp4?label=mp4_ld&template=640x360.25.0&trans_finger=40a32e8439c5409a63ccf853562a60ef&Expires=1576729427&ssig=sotYdsogqZ&KID=unistore,video
}

func TestWeiBoVideo(t *testing.T) {
	req := WeiBoVideoRequest{
		Query: "杨幂",
		Host:  defaultWBHost,
	}
	client := setransport.NewClient(req.Host, req.Query)
	response, e := req.Do(context.TODO(), client)
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content), e)
}
