// Created at 1/20/2021
// Developer: trungnq2710 (trungnq2710@gmail.com)

package go_http_client

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
	c.method = m
	return c
}
