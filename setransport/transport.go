package setransport

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
)

type CacheTransport struct {
	mux  sync.Mutex
	Data map[string]io.ReadCloser
	Ori  http.RoundTripper
}

func newCacheTransport() *CacheTransport {
	return &CacheTransport{
		Data: make(map[string]io.ReadCloser),
		Ori:  http.DefaultTransport,
	}
}

func (C *CacheTransport) set(r *http.Request, res *http.Response) {
	// 获取网络请求数据
	// key: 代表 url
	// val: 代表 网络请求响应
	C.mux.Lock()
	defer C.mux.Unlock()
	body := ioutil.NopCloser(res.Body)
	C.Data[r.URL.String()] = body
}
func (C *CacheTransport) get(r *http.Request) (io.ReadCloser, error) {
	// 获取网络请求数据
	C.mux.Lock()
	defer C.mux.Unlock()
	if val, ok := C.Data[r.URL.String()]; ok {
		return val, nil
	}
	return nil, errors.New("key not found")
}

func (C *CacheTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	if val, err := C.get(r); err == nil {
		return &http.Response{
			StatusCode: http.StatusOK,
			Header:     r.Header,
			Body:       val,
			Request:    r,
		}, nil
	}

	resp, err := C.Ori.RoundTrip(r)
	if err != nil {
		return nil, err
	}
	C.set(r, resp)
	return resp, nil
}
