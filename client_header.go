// Created at 1/20/2021

package go_http_client

func (c *Client) AddHeader(key, value string) *Client {
	c.header.Add(key, value)
	return c
}

func (c *Client) SetHeader(key, value string) *Client {
	c.header.Set(key, value)
	return c
}
