// Created at 1/20/2021
// Developer: trungnq2710 (trungnq2710@gmail.com)

package httpclient

import "time"

const (
	defaultMaxIdleConnPerHost               = 50
	defaultMaxIdleConn                      = 50
	defaultRetries                          = 3
	defaultRetryDelay         time.Duration = 400
	defaultMethod                           = MethodGet
)

const (
	contentTypeKey       = "Content-Type"
	contentTypeJSON      = "application/json"
	contentTypeXML       = "application/xml"
	contentTypeForm      = "application/x-www-form-urlencoded"
	contentTypePlainText = "text/plain; charset=utf-8"
)

const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH" // RFC 5789
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
)

const (
	bodyTypeDefault = "BODY_TYPE_DEFAULT"
	bodyTypeJSON    = "BODY_TYPE_JSON"
	bodyTypeXML     = "BODY_TYPE_XML"
	bodyTypeForm    = "BODY_TYPE_FORM"
)
