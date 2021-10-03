// Created at 10/3/2021 8:43 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package go_httpclient

import "net/http"

func (c *Client) SetCookieJar(jar http.CookieJar) *Client {
	c.httpClient.Jar = jar
	return c
}

func (c *Client) SetCookie(hc *http.Cookie) *Client {
	c.cookies = append(c.cookies, hc)
	return c
}

func (c *Client) SetCookies(cs []*http.Cookie) *Client {
	c.cookies = append(c.cookies, cs...)
	return c
}
