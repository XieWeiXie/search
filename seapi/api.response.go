package seapi

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func (r *Response) String() string {
	var (
		out = new(bytes.Buffer)
		b1  = bytes.NewBuffer([]byte{})
		b2  = bytes.NewBuffer([]byte{})
		tr  io.Reader
	)

	if r != nil && r.Body != nil {
		tr = io.TeeReader(r.Body, b1)
		defer r.Body.Close()

		if _, err := io.Copy(b2, tr); err != nil {
			out.WriteString(fmt.Sprintf("<error reading response body: %v>", err))
			return out.String()
		}
		defer func() { r.Body = ioutil.NopCloser(b1) }()
	}

	if r != nil {
		out.WriteString(fmt.Sprintf("[%d %s]", r.StatusCode, http.StatusText(r.StatusCode)))
		if r.StatusCode > 0 {
			out.WriteRune(' ')
		}
	} else {
		out.WriteString("[0 <nil>]")
	}

	if r != nil && r.Body != nil {
		out.ReadFrom(b2) // errcheck exclude (*bytes.Buffer).ReadFrom
	}

	return out.String()
}
