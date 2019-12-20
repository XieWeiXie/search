<h1 align="center">Search</h1>
<p align="center">
  <em>Why Search?  Inspired by <a href="https://github.com/elastic/go-elasticsearch">go-elasticsearch</a>.</em>
</p>
<p align="center">
    <a href="https://github.com/wuxiaoxiaoshen">
        <img src="https://img.shields.io/badge/Author-wuxiaoxiaoshen-green" alt="Author">
    </a>
    <a href="https://github.com/wuxiaoxiaoshen">
        <img src="https://img.shields.io/badge/progressing-95%25-green" alt="Author">
    </a>
</p>

> Just For Search ...
---
> Program to an interface, not an implementation

### Install

```text
go get -u -v github.com/wuxiaoxiaoshen/search
```

### Example

**1. Demo one: Use Api**
```go
package main

import (
	"context"
	"fmt"
	"github.com/wuxiaoxiaoshen/search"
	"github.com/wuxiaoxiaoshen/search/seapi"
	"github.com/wuxiaoxiaoshen/search/setransport"
	"io/ioutil"
	"log"
)
func Demo1ZhiHuTopic(query string) {
	req := seapi.ZhiHuTopicRequest{
		Query: query,
	}
	client := setransport.NewClient(req.Query)
	response, _ := req.Do(context.TODO(), client)
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
func main(){
	for _, i := range []string{"杨幂", "刘诗诗", "刘亦菲"} {
		Demo1ZhiHuTopic(i)
	}
}
```
```
>>
2019/12/20 10:28:40 ZhiHu Topic Url https://www.zhihu.com/search?q=%E6%9D%A8%E5%B9%82&type=topic
[{"Title":"杨幂","Text":"杨幂（1986 年 9 月 12 日－）是一名中国著名女演员23,136 关注 · 4,286 问题 · 999 精华内u.com/topic/19580325","Concern":{"Number":"23,136 关注","SubLink":"https://www.zhihu.com/topic/19580325/top-aners"},"Question":{"Number":"4,286 问题","SubLink":""},"Hot":{"Number":"999 精华内容","SubLink":""}},{"Title":","Text":"杨幂工作室是欢瑞世纪影视传媒股份有限公司旗下的独立工作室，于2013年3月31日在北京启动的“大制片计划”中正嘉行天下杨幂工作室。近日，完美世界对杨幂公司投资增至5亿元 持股10%。35 关注 · 1 问题 · 3 精华内容","Link":"httpern":{"Number":"35 关注","SubLink":"https://www.zhihu.com/topic/20594781/top-answers"},"Question":{"Number":"1题","SubLink":""},"Hot":{"Number":"3 精华内容","SubLink":""}},{"Title":"谈判官（杨幂、黄子韬主演同名电视剧）（和胆大心细的谈判风格，在商务谈判桌上无1 关注 · 19 问题 · 5 精华内容","Link":"https://www.zhihu.com/topic/20867 关注","SubLink":"https://www.zhihu.com/topic/20867668/top-answers"},"Question":{"Number":"19 问题","SubLink":"Hot":{"Number":"5 精华内容","SubLink":""}},{"Title":"明日之子 第二季（综艺）","Text":"由李宇春、吴青峰、华晨宇星推官」。5,660 关注 · 1,003 问题 · 713 精华内容","Link":"https://www.zhihu.com/topic/20206639","Concern":{"Nu关注","SubLink":"https://www.zhihu.com/topic/20206639/top-answers"},"Question":{"Number":"1,003 问题","SubLink},"Hot":{"Number":"713 精华内容","SubLink":""}},{"Title":"倪妮（演员）","Text":"2013年1月，与Angelababy、杨幂及315 关注 · 379 问题 · 364 精华内容","Link":"https://www.zhihu.com/topic/20087897","Concern":{"Number":"7,315 关ink":"https://www.zhihu.com/topic/20087897/top-answers"},"Question":{"Number":"379 问题","SubLink":""},"Hot":{umber":"364 精华内容","SubLink":""}}]
2019/12/20 10:28:41 ZhiHu Topic Url https://www.zhihu.com/search?q=%E5%88%98%E8%AF%97%E8%AF%97&type=topic
[{"Title":"刘诗诗（演员）","Text":"刘诗诗，原名刘诗施，1987年3月10日出生于北京市，中国内地影视女演员、影视出品tps://www.zhihu.com/topic/19899085","Concern":{"Number":"8,807 关注","SubLink":"https://www.zhihu.com/topic/199085/top-answers"},"Question":{"Number":"491 问题","SubLink":""},"Hot":{"Number":"505 精华内容","SubLink":""}}le":"刘诗诗粉丝","Text":"58 关注 · 22 问题 · 13 精华内容","Link":"https://www.zhihu.com/topic/20086882","Conce:"58 关注","SubLink":"https://www.zhihu.com/topic/20086882/top-answers"},"Question":{"Number":"22 问题","SubLi""},"Hot":{"Number":"13 精华内容","SubLink":""}},{"Title":"倪妮（演员）","Text":"与Angelababy、杨幂及刘诗诗被评电影《杀戒》、爱情电影《我想和你好好的》及治愈系电影《等风来》7,315 关注 · 379 问题 · 364 精华内容","Link":"htncern":{"Number":"7,315 关注","SubLink":"https://www.zhihu.com/topic/20087897/top-answers"},"Question":{"Numbe:"379 问题","SubLink":""},"Hot":{"Number":"364 精华内容","SubLink":""}},{"Title":"绣春刀（电影）","Text":"录电影视文化有限公司、北京合力映画影视文化传媒有限公司联合出品的浪漫武侠电影，由路阳执导，张震、刘诗诗、王千源、李s://www.zhihu.com/topic/20007663","Concern":{"Number":"5,384 关注","SubLink":"https://www.zhihu.com/topic/200063/top-answers"},"Question":{"Number":"225 问题","SubLink":""},"Hot":{"Number":"190 精华内容","SubLink":""}},{":"步步惊心","Text":"由上海唐人电影制作有限公司和湖南卫视联合出品，该剧由李国立执导，由刘诗诗、吴奇隆、郑嘉颖、https://www.zhihu.com/topic/19647030","Concern":{"Number":"3,459 关注","SubLink":"https://www.zhihu.com/topic/647030/top-answers"},"Question":{"Number":"250 问题","SubLink":""},"Hot":{"Number":"185 精华内容","SubLink":""}}]
2019/12/20 10:28:41 ZhiHu Topic Url https://www.zhihu.com/search?q=%E5%88%98%E4%BA%A6%E8%8F%B2&type=topic
[{"Title":"刘亦菲（演员）","Text":"刘亦菲，1987年8月25日出生于湖北省武汉市，华语影视女演员、歌手32,994 关注 · /www.zhihu.com/topic/19589214","Concern":{"Number":"32,994 关注","SubLink":"https://www.zhihu.com/topic/195892/top-answers"},"Question":{"Number":"1,748 问题","SubLink":""},"Hot":{"Number":"999 精华内容","SubLink":""}},{":"刘亦菲（刘亦菲同名专辑）","Text":"是刘亦菲正式出道后于2006年8月31日发行的首张同名国语专辑，共收录10首歌曲，":"https://www.zhihu.com/topic/20426177","Concern":{"Number":"47 关注","SubLink":"https://www.zhihu.com/topic/426177/top-answers"},"Question":{"Number":"0 问题","SubLink":""},"Hot":{"Number":"0 精华内容","SubLink":""}},{":"仙剑奇侠传（电视剧）","Text":"李国立制作并导演，由胡歌、刘亦菲、安以轩、刘品言、彭于晏等主演4,850 关注 · 1,u.com/topic/20012651","Concern":{"Number":"4,850 关注","SubLink":"https://www.zhihu.com/topic/20012651/top-ansrs"},"Question":{"Number":"1,376 问题","SubLink":""},"Hot":{"Number":"998 精华内容","SubLink":""}},{"Title":"三花（电影）","Text":"LaMolinara）联合执导，刘亦菲、杨洋领衔主演的古装玄幻爱情仙侠电影2,718 关注 · 647 问题 · 42m/topic/20076732","Concern":{"Number":"2,718 关注","SubLink":"https://www.zhihu.com/topic/20076732/top-answers,"Question":{"Number":"647 问题","SubLink":""},"Hot":{"Number":"420 精华内容","SubLink":""}},{"Title":"方大同"叱吒乐坛流行榜颁奖典礼上同时夺得“叱吒乐坛男歌手金奖”、“叱吒乐坛唱作人金奖”及“叱吒乐坛作曲人大奖”。2010年，与王精华内容","Link":"https://www.zhihu.com/topic/19748614","Concern":{"Number":"2,454 关注","SubLink":"https://wwu.com/topic/19748614/top-answers"},"Question":{"Number":"195 问题","SubLink":""},"Hot":{"Number":"121 精华内容Link":""}}]
```

**2. Demo Two: Use Client**
```go
package main

import (
	"context"
	"fmt"
	"github.com/wuxiaoxiaoshen/search"
	"github.com/wuxiaoxiaoshen/search/seapi"
	"github.com/wuxiaoxiaoshen/search/setransport"
	"io/ioutil"
	"log"
)

func Demo2ZhiHuTopic(query string) {
	cfg := search.NewClientConfig(query)
	client := search.NewClient(*cfg)
	response, _ := client.ZhiHu.Topic(cfg.Query)
	content, _ := ioutil.ReadAll(response.Body)
	log.Println(string(content))
}
func main(){
	for _, i := range []string{"杨幂", "刘诗诗", "刘亦菲"} {
		Demo2ZhiHuTopic(i)
	}
}
```

```text
>>
2019/12/20 10:28:40 ZhiHu Topic Url https://www.zhihu.com/search?q=%E6%9D%A8%E5%B9%82&type=topic
[{"Title":"杨幂","Text":"杨幂（1986 年 9 月 12 日－）是一名中国著名女演员23,136 关注 · 4,286 问题 · 999 精华内u.com/topic/19580325","Concern":{"Number":"23,136 关注","SubLink":"https://www.zhihu.com/topic/19580325/top-aners"},"Question":{"Number":"4,286 问题","SubLink":""},"Hot":{"Number":"999 精华内容","SubLink":""}},{"Title":","Text":"杨幂工作室是欢瑞世纪影视传媒股份有限公司旗下的独立工作室，于2013年3月31日在北京启动的“大制片计划”中正嘉行天下杨幂工作室。近日，完美世界对杨幂公司投资增至5亿元 持股10%。35 关注 · 1 问题 · 3 精华内容","Link":"httpern":{"Number":"35 关注","SubLink":"https://www.zhihu.com/topic/20594781/top-answers"},"Question":{"Number":"1题","SubLink":""},"Hot":{"Number":"3 精华内容","SubLink":""}},{"Title":"谈判官（杨幂、黄子韬主演同名电视剧）（和胆大心细的谈判风格，在商务谈判桌上无1 关注 · 19 问题 · 5 精华内容","Link":"https://www.zhihu.com/topic/20867 关注","SubLink":"https://www.zhihu.com/topic/20867668/top-answers"},"Question":{"Number":"19 问题","SubLink":"Hot":{"Number":"5 精华内容","SubLink":""}},{"Title":"明日之子 第二季（综艺）","Text":"由李宇春、吴青峰、华晨宇星推官」。5,660 关注 · 1,003 问题 · 713 精华内容","Link":"https://www.zhihu.com/topic/20206639","Concern":{"Nu关注","SubLink":"https://www.zhihu.com/topic/20206639/top-answers"},"Question":{"Number":"1,003 问题","SubLink},"Hot":{"Number":"713 精华内容","SubLink":""}},{"Title":"倪妮（演员）","Text":"2013年1月，与Angelababy、杨幂及315 关注 · 379 问题 · 364 精华内容","Link":"https://www.zhihu.com/topic/20087897","Concern":{"Number":"7,315 关ink":"https://www.zhihu.com/topic/20087897/top-answers"},"Question":{"Number":"379 问题","SubLink":""},"Hot":{umber":"364 精华内容","SubLink":""}}]
2019/12/20 10:28:41 ZhiHu Topic Url https://www.zhihu.com/search?q=%E5%88%98%E8%AF%97%E8%AF%97&type=topic
[{"Title":"刘诗诗（演员）","Text":"刘诗诗，原名刘诗施，1987年3月10日出生于北京市，中国内地影视女演员、影视出品tps://www.zhihu.com/topic/19899085","Concern":{"Number":"8,807 关注","SubLink":"https://www.zhihu.com/topic/199085/top-answers"},"Question":{"Number":"491 问题","SubLink":""},"Hot":{"Number":"505 精华内容","SubLink":""}}le":"刘诗诗粉丝","Text":"58 关注 · 22 问题 · 13 精华内容","Link":"https://www.zhihu.com/topic/20086882","Conce:"58 关注","SubLink":"https://www.zhihu.com/topic/20086882/top-answers"},"Question":{"Number":"22 问题","SubLi""},"Hot":{"Number":"13 精华内容","SubLink":""}},{"Title":"倪妮（演员）","Text":"与Angelababy、杨幂及刘诗诗被评电影《杀戒》、爱情电影《我想和你好好的》及治愈系电影《等风来》7,315 关注 · 379 问题 · 364 精华内容","Link":"htncern":{"Number":"7,315 关注","SubLink":"https://www.zhihu.com/topic/20087897/top-answers"},"Question":{"Numbe:"379 问题","SubLink":""},"Hot":{"Number":"364 精华内容","SubLink":""}},{"Title":"绣春刀（电影）","Text":"录电影视文化有限公司、北京合力映画影视文化传媒有限公司联合出品的浪漫武侠电影，由路阳执导，张震、刘诗诗、王千源、李s://www.zhihu.com/topic/20007663","Concern":{"Number":"5,384 关注","SubLink":"https://www.zhihu.com/topic/200063/top-answers"},"Question":{"Number":"225 问题","SubLink":""},"Hot":{"Number":"190 精华内容","SubLink":""}},{":"步步惊心","Text":"由上海唐人电影制作有限公司和湖南卫视联合出品，该剧由李国立执导，由刘诗诗、吴奇隆、郑嘉颖、https://www.zhihu.com/topic/19647030","Concern":{"Number":"3,459 关注","SubLink":"https://www.zhihu.com/topic/647030/top-answers"},"Question":{"Number":"250 问题","SubLink":""},"Hot":{"Number":"185 精华内容","SubLink":""}}]
2019/12/20 10:28:41 ZhiHu Topic Url https://www.zhihu.com/search?q=%E5%88%98%E4%BA%A6%E8%8F%B2&type=topic
[{"Title":"刘亦菲（演员）","Text":"刘亦菲，1987年8月25日出生于湖北省武汉市，华语影视女演员、歌手32,994 关注 · /www.zhihu.com/topic/19589214","Concern":{"Number":"32,994 关注","SubLink":"https://www.zhihu.com/topic/195892/top-answers"},"Question":{"Number":"1,748 问题","SubLink":""},"Hot":{"Number":"999 精华内容","SubLink":""}},{":"刘亦菲（刘亦菲同名专辑）","Text":"是刘亦菲正式出道后于2006年8月31日发行的首张同名国语专辑，共收录10首歌曲，":"https://www.zhihu.com/topic/20426177","Concern":{"Number":"47 关注","SubLink":"https://www.zhihu.com/topic/426177/top-answers"},"Question":{"Number":"0 问题","SubLink":""},"Hot":{"Number":"0 精华内容","SubLink":""}},{":"仙剑奇侠传（电视剧）","Text":"李国立制作并导演，由胡歌、刘亦菲、安以轩、刘品言、彭于晏等主演4,850 关注 · 1,u.com/topic/20012651","Concern":{"Number":"4,850 关注","SubLink":"https://www.zhihu.com/topic/20012651/top-ansrs"},"Question":{"Number":"1,376 问题","SubLink":""},"Hot":{"Number":"998 精华内容","SubLink":""}},{"Title":"三花（电影）","Text":"LaMolinara）联合执导，刘亦菲、杨洋领衔主演的古装玄幻爱情仙侠电影2,718 关注 · 647 问题 · 42m/topic/20076732","Concern":{"Number":"2,718 关注","SubLink":"https://www.zhihu.com/topic/20076732/top-answers,"Question":{"Number":"647 问题","SubLink":""},"Hot":{"Number":"420 精华内容","SubLink":""}},{"Title":"方大同"叱吒乐坛流行榜颁奖典礼上同时夺得“叱吒乐坛男歌手金奖”、“叱吒乐坛唱作人金奖”及“叱吒乐坛作曲人大奖”。2010年，与王精华内容","Link":"https://www.zhihu.com/topic/19748614","Concern":{"Number":"2,454 关注","SubLink":"https://wwu.com/topic/19748614/top-answers"},"Question":{"Number":"195 问题","SubLink":""},"Hot":{"Number":"121 精华内容Link":""}}]

```


### Supported API

zhihu:

```text
- seapi.ZhiHuTopicRequest
```

weibo:

```text
- seapi.WeiBoUserRequest
- seapi.WeiBoPassageRequest
- seapi.WeiBoVideoRequest
- seapi.WeiBoPictureRequest
- seapi.WeiBoTopicRequest
```


### License
MIT [©wuxiaoxiaoshen](https://github.com/wuxiaoxiaoshen)