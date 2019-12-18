package seapi

import "net/url"

type SearchUrl struct {
	Q     string
	toUrl *url.URL
	query url.Values
}

func (S *SearchUrl) Url() {
	S.toUrl, _ = url.Parse(S.Q)
	S.query = make(url.Values)
	S.query = S.toUrl.Query()
}

func (S *SearchUrl) Add(key, value string) {
	S.query.Add(key, value)
}
func (S *SearchUrl) Format() string {
	S.toUrl.RawQuery = S.query.Encode()
	return S.toUrl.String()
}
