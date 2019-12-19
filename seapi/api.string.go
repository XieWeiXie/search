package seapi

import (
	"net/url"
	"regexp"
	"strings"
)

type StringHandler struct {
	Value string
}

func (S StringHandler) Replacer(old ...string) string {
	replacer := strings.NewReplacer(old...)
	return replacer.Replace(S.Value)
}
func (S StringHandler) Regexp(pattern string) [][]string {
	reg := regexp.MustCompile(pattern)
	return reg.FindAllStringSubmatch(S.Value, -1)
}

func (S StringHandler) Url(key string, value string) string {
	query, _ := url.ParseQuery(value)
	return query.Get(key)

}
