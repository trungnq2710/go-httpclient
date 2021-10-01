// Created at 1/20/2021
// Developer: trungnq2710 (trungnq2710@gmail.com)

package go_httpclient

func (c *Client) AddHeader(key, value string) *Client {
	c.header.Add(key, value)
	return c
}

func (c *Client) SetHeader(key, value string) *Client {
	c.header.Set(key, value)
	return c
}
