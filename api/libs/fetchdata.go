package libs

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
)

type transport struct {
	headers map[string]string
	base    http.RoundTripper
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range t.headers {
		req.Header.Add(k, v)
	}
	base := t.base
	if base == nil {
		base = http.DefaultTransport
	}
	return base.RoundTrip(req)
}

type FetchDataConfig struct {
	URL         string
	TokenHeader string
	TimeOut     int64
}

type Option func(sr *FetchDataConfig)

func NewFetchData(opts ...Option) *FetchDataConfig {

	fetchcfg := FetchDataConfig{}
	// get all options need
	for _, option := range opts {
		option(&fetchcfg)
	}

	return &fetchcfg
}

func (f *FetchDataConfig) GetJSON(query string) (gjson.Result, error) {

	timeout := time.Duration(time.Duration(f.TimeOut) * time.Second)
	cli := &http.Client{
		Timeout: timeout,
		Transport: &transport{
			headers: map[string]string{
				"x-api-key": f.TokenHeader,
			},
		},
	}

	resp, err := cli.Get(fmt.Sprintf("%s?q=%s", f.URL, query))
	if err != nil {
		return gjson.Parse(""), err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200, 201:
		// Status OK
	default:
		return gjson.Parse(""), errors.New("Something wrong with the connection")
	}

	buf, _ := ioutil.ReadAll(resp.Body)
	return gjson.ParseBytes(buf), nil
}
