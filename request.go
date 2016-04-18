package gorouter

import "net/http"

type Request interface {
	IsHTTP() bool
	Valid() bool
	Bytes() ([]byte, error)
	FromBytes([]byte) error
	GetRequest() *http.Request
}
