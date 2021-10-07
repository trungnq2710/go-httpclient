// Created at 9/28/2021 8:08 AM
// Developer: trungnq2710 (trungnq2710@gmail.com)

package httpclient

import "errors"

var (
	errorMethodUnknown = errors.New("found unknown method")
	errorXMLConvert = errors.New("obj could not be converted to XML data")
)
