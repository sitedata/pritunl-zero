package search

import (
	"github.com/dropbox/godropbox/container/set"
	"net/http"
	"net/url"
	"time"
)

var (
	RequestTypes = set.NewSet(
		"application/xml",
		"application/json",
		"application/x-www-form-urlencoded",
	)
)

type Request struct {
	User      string      `json:"user"`
	Session   string      `json:"session"`
	Address   string      `json:"address"`
	Timestamp time.Time   `json:"timestamp"`
	Scheme    string      `json:"scheme"`
	Host      string      `json:"host"`
	Path      string      `json:"path"`
	Query     url.Values  `json:"query"`
	Header    http.Header `json:"header"`
	Body      string      `json:"body"`
}

func (r *Request) Index() (err error) {
	Index("zero-requests", "request", r)
	return
}
