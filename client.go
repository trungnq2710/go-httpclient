// Created at 1/19/2021

package go_http_client

import (
	"io"
	"net/http"
	"time"
)

type Client struct {
	setting    *Settings
	httpClient *http.Client
	endpoint   string
	path       string
	method     string
	header     http.Header
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

var defaultHttpClient = &http.Client{
	Transport: &http.Transport{
		MaxIdleConns:        defaultMaxIdleConn,
		MaxIdleConnsPerHost: defaultMaxIdleConnPerHost,
	},
}

func New() *Client {
	return &Client{
		setting: &Settings{
			Retries:    3,
			RetryDelay: 400,
		},
		httpClient:      defaultHttpClient,
		method:          MethodGet,
		header:          make(http.Header),
		bodyType:        bodyTypeDefault,
		responseDecoder: jsonDecoder{},
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
