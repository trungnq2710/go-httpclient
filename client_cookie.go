// Created at 10/3/2021 8:43 PM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package httpclient

import (
	"golang.org/x/net/publicsuffix"
	"net/http"
	"net/http/cookiejar"
)

func (c *Client) UseCookieJar() *Client {
	cookieJar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	c.httpClient.Jar = cookieJar
	return c
}

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
