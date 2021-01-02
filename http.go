package myshows

import (
	"crypto/tls"
	"net/http"
	"time"
)

var (
	tr = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}
	httpClient = &http.Client{Transport: tr, Timeout: 7 * time.Second}
)
