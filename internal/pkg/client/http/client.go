package http

import (
	"time"

	"github.com/go-resty/resty/v2"

	"go-template/internal/pkg/config"
)

func NewClient() *resty.Client {
	cfg := config.Instance().Common
	var client = resty.New()
	client.
		SetRetryCount(cfg.HttpClientRetryCount).
		SetRetryWaitTime(time.Duration(cfg.HttpClientRetryWaitTimeSeconds) * time.Second).
		SetRetryMaxWaitTime(time.Duration(cfg.HttpClientRetryMaxWaitTimeSeconds) * time.Second).
		AddRetryCondition(
			func(r *resty.Response, err error) bool {
				if r == nil {
					return false
				}
				return r.StatusCode() >= 500
			},
		)
	return client
}

func R[R, E any](c *resty.Client) *Request[R, E] {
	restyReq := c.R()
	r := &Request[R, E]{
		LogReqRes: true,
		Request:   restyReq,
	}
	return r
}
