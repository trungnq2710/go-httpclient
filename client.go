// Created at 1/19/2021
// Developer: trungnq2710 (trungnq2710@gmail.com)

package go_httpclient

import (
	"golang.org/x/net/publicsuffix"
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type Client struct {
	setting    *Settings
	httpClient *http.Client
	endpoint   string
	path       string
	method     string
	header     http.Header
	cookies    []*http.Cookie
	queries    map[string]string
	bodyType   string
	body       interface{}
	reqBody    io.Reader
	req        *http.Request
	resp       *http.Response

	responseDecoder ResponseDecoder
	successV        interface{}
	failureV        interface{}
}

func New() *Client {
	return &Client{
		setting: &Settings{
			Retries:    defaultRetries,
			RetryDelay: defaultRetryDelay,
		},
		httpClient:      getDefaultHttpClient(),
		method:          defaultMethod,
		header:          make(http.Header),
		cookies:         make([]*http.Cookie, 0),
		bodyType:        bodyTypeDefault,
		responseDecoder: jsonDecoder{},
	}
}

func getDefaultHttpClient() *http.Client {
	cookieJar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:        defaultMaxIdleConn,
			MaxIdleConnsPerHost: defaultMaxIdleConnPerHost,
		},
		Jar: cookieJar,
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// other <-

func (c *Client) Base(endpoint string) *Client {
	c.endpoint = endpoint
	return c
}

func (c *Client) SetRetries(r int) *Client {
	c.setting.Retries = r
	return c
}

func (c *Client) SetRetryDelay(d time.Duration) *Client {
	c.setting.RetryDelay = d
	return c
}

func (c *Client) SetSuccessV(v interface{}) *Client {
	c.successV = v
	return c
}

func (c *Client) SetFailureV(v interface{}) *Client {
	c.failureV = v
	return c
}

func (c *Client) SetTransport(transport http.RoundTripper) *Client {
	if transport != nil {
		c.httpClient.Transport = transport
	}
	return c
}
