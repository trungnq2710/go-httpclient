// Created at 1/20/2021
// Developer: trungnq2710 (trungnq2710@gmail.com)

package go_http_client

const (
	defaultMaxIdleConnPerHost = 50
	defaultMaxIdleConn        = 50
)

const (
	contentTypeKey  = "Content-Type"
	contentTypeJson = "application/json"
	contentTypeForm = "application/x-www-form-urlencoded"
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
	bodyTypeJson    = "BODY_TYPE_JSON"
	bodyTypeForm    = "BODY_TYPE_FORM"
)
