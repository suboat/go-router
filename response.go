package gorouter

import "net/http"

type Response interface {
	IsHTTP() bool
	Valid() bool
	Bytes() ([]byte, error)
	FromBytes([]byte) error
	GetResponseWriter() http.ResponseWriter
	// TODO: 待需求扩展
}
