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



### Support API

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