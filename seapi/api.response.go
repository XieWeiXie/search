package seapi

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type Response struct {
	StatusCode int
	Header     http.Header
	Body       io.ReadCloser
}

func newResponse(results interface{}, response *http.Response) *Response {
	b, _ := json.Marshal(results)
	var by = bytes.NewBuffer(b)
	var buf bytes.Buffer
	tee := io.TeeReader(by, &buf)
	newResponse := Response{
		Body:       ioutil.NopCloser(tee),
		StatusCode: response.StatusCode,
		Header:     response.Header,
	}
	return &newResponse
}
