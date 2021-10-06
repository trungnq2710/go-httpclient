// Created at 1/20/2021
// Developer: trungnq2710 (trungnq2710@gmail.com)

package go_httpclient

func (c *Client) SetQueryParam(param, value string) *Client {
	c.queries.Set(param, value)
	return c
}

func (c *Client) SetQueryParams(params map[string]string) *Client {
	for p, v := range params {
		c.SetQueryParam(p, v)
	}
	return c
}
