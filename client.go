package cxsign

import (
	"net/http"
	"net/http/cookiejar"

	"golang.org/x/net/publicsuffix"
)

func NewClient() *http.Client {
	// Cookie
	cookieJar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	c := &http.Client{Jar: cookieJar}
	return c
}
