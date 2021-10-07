// Created at 1/20/2021
// Developer: trungnq2710 (trungnq2710@gmail.com)

package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	goquery "github.com/google/go-querystring/query"
	"io/ioutil"
)

func (c *Client) SetBody(body interface{}) *Client {
	c.bodyType = bodyTypeDefault
	c.body = body
	return c
}

func (c *Client) SetJsonBody(body interface{}) *Client {
	c.SetHeader(contentTypeKey, contentTypeJson)
	c.bodyType = bodyTypeJson
	c.body = body
	return c
}

func (c *Client) SetFormBody(body interface{}) *Client {
	c.SetHeader(contentTypeKey, contentTypeForm)
	c.bodyType = bodyTypeForm
	c.body = body
	return c
}

func (c *Client) buildBody() (err error) {
	switch c.bodyType {
	case bodyTypeDefault:
	case bodyTypeJson:
		bf := &bytes.Buffer{}
		err = json.NewEncoder(bf).Encode(c.body)
		if err != nil {
			return
		}
		c.reqBody = ioutil.NopCloser(bf)
		return
	case bodyTypeForm:
		values, err1 := goquery.Values(c.body)
		if err1 != nil {
			return err1
		}
		bf := bytes.NewBufferString(values.Encode())
		c.reqBody = ioutil.NopCloser(bf)
		return
	default:
		err = fmt.Errorf("unsupported body type: %s", c.bodyType)
		return
	}

	switch t := c.body.(type) {
	case string:
		bf := bytes.NewBufferString(t)
		c.reqBody = ioutil.NopCloser(bf)
	case []byte:
		bf := bytes.NewBuffer(t)
		c.reqBody = ioutil.NopCloser(bf)
	default:
		err = fmt.Errorf("unsupported body data type: %s", t)
		return
	}
	return
}
