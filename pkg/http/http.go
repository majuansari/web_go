package http

import (
	"crypto/tls"
	"net/http"
	"time"
)

func InitHttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 100,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 10 * time.Second,
	}
}
