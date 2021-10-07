// Created at 1/20/2021
// Developer: trungnq2710 (trungnq2710@gmail.com)

package httpclient

import "strings"

func (c *Client) UseGet() *Client {
	c.method = MethodGet
	return c
}

func (c *Client) UseHead() *Client {
	c.method = MethodHead
	return c
}

func (c *Client) UsePost() *Client {
	c.method = MethodPost
	return c
}

func (c *Client) UsePut() *Client {
	c.method = MethodPut
	return c
}

func (c *Client) UsePatch() *Client {
	c.method = MethodPatch
	return c
}

func (c *Client) UseDelete() *Client {
	c.method = MethodDelete
	return c
}

func (c *Client) UseConnect() *Client {
	c.method = MethodConnect
	return c
}

func (c *Client) UseOptions() *Client {
	c.method = MethodOptions
	return c
}

func (c *Client) UseTrace() *Client {
	c.method = MethodTrace
	return c
}

func (c *Client) SetMethod(m string) *Client {
	switch strings.ToUpper(m) {
	case MethodGet:
		return c.UseGet()
	case MethodHead:
		return c.UseHead()
	case MethodPost:
		return c.UsePost()
	case MethodPut:
		return c.UsePut()
	case MethodPatch:
		return c.UsePatch()
	case MethodDelete:
		return c.UseDelete()
	case MethodConnect:
		return c.UseConnect()
	case MethodOptions:
		return c.UseOptions()
	case MethodTrace:
		return c.UseTrace()
	default:
		// TODO trungnq set err
	}
	return c.UseGet()
}
