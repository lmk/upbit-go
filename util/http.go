package util

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

type RequestOptions struct {
	Url     string
	Method  string
	Body    io.Reader
	Query   map[string]string
	Headers map[string]string
}

// Encode encodes the values into ``URL encoded'' form
// ("bar=baz&foo=quux") sorted by key.
func Encode(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		keyEscaped := k
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(keyEscaped)
			buf.WriteByte('=')
			buf.WriteString(v)
		}
	}
	return buf.String()
}

// Request request
func Request(options *RequestOptions, result interface{}) (
	err error,
) {
	client := &http.Client{}

	req, err := http.NewRequest(options.Method, options.Url, options.Body)
	if err != nil {
		return
	}

	if options.Query != nil {
		q := req.URL.Query()
		for index, value := range options.Query {
			q.Add(index, value)
		}

		req.URL.RawQuery = Encode(q)
	}

	if options.Headers != nil {
		for prop, value := range options.Headers {
			req.Header.Add(prop, value)
		}
	}

	res, err := client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	Body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(Body, result)
	if err != nil {
		err = errors.New(err.Error() + " " + string(Body))
		return
	}
	return
}
