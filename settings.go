// Created at 1/19/2021

package go_http_client

import "time"

type Settings struct {
	Gzip       bool
	Retries    int
	RetryDelay time.Duration
}
