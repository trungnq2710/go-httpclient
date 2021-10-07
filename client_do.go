// Created at 1/20/2021
// Developer: trungnq2710 (trungnq2710@gmail.com)

package httpclient

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func (c *Client) buildRequest() (err error) {
	reqURL, err := url.Parse(c.endpoint)
	if err != nil {
		return err
	}

	if c.body != nil {
		err = c.buildBody()
		if err != nil {
			return err
		}
	}

	c.req, err = http.NewRequest(c.method, reqURL.String(), c.reqBody)
	if err != nil {
		return err
	}

	for key, values := range c.header {
		for _, value := range values {
			c.req.Header.Add(key, value)
		}
	}

	if len(c.queries) > 0 {
		q := c.req.URL.Query()
		for k, v := range c.queries {
			for _, iv := range v {
				q.Add(k, iv)
			}
		}
		c.req.URL.RawQuery = q.Encode()
	}

	for _, cookie := range c.cookies {
		c.req.AddCookie(cookie)
	}

	return err
}

func (c *Client) Do() (resp *http.Response, err error) {
	err = c.buildRequest()
	if err != nil {
		return nil, err
	}

	// retries default value is 0, it will run once.
	// retries equal to -1, it will run forever until success
	for i := 0; c.setting.Retries == -1 || i <= c.setting.Retries; i++ {
		resp, err = c.httpClient.Do(c.req)
		if err == nil {
			break
		}
		time.Sleep(c.setting.RetryDelay)
	}

	if err != nil {
		return
	}

	//defer resp.Body.Close()
	//defer io.Copy(ioutil.Discard, resp.Body)

	// Don't try to decode on 204s or Content-Length is 0
	if resp.StatusCode == http.StatusNoContent || resp.ContentLength == 0 {
		return resp, nil
	}

	// Decode from json
	if c.successV != nil || c.failureV != nil {
		err = decodeResponse(resp, c.responseDecoder, c.successV, c.failureV)
	}
	return resp, err
}

func (c *Client) Exec() error {
	resp, err := c.Do()
	if err != nil {
		return err
	}

	// when err is nil, resp contains a non-nil resp.Body which must be closed
	defer resp.Body.Close()

	// The default HTTP client's Transport may not
	// reuse HTTP/1.x "keep-alive" TCP connections if the Body is
	// not read to completion and closed.
	// See: https://golang.org/pkg/net/http/#Response
	defer io.Copy(ioutil.Discard, resp.Body)

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// decoder <-

func decodeResponse(resp *http.Response, decoder ResponseDecoder, successV, failureV interface{}) error {
	if code := resp.StatusCode; 200 <= code && code <= 299 {
		if successV != nil {
			return decoder.Decode(resp, successV)
		}
	} else {
		if failureV != nil {
			return decoder.Decode(resp, failureV)
		}
	}
	return nil
}

type ResponseDecoder interface {
	Decode(resp *http.Response, v interface{}) error
}

type jsonDecoder struct {
}

func (d jsonDecoder) Decode(resp *http.Response, v interface{}) error {
	return json.NewDecoder(resp.Body).Decode(v)
}
