package gorouter

import "net/http"

type Route interface {
	ListenAndServe(string) error
	Error() error
}

type HPRouter interface {
	Route
	Handle(string, http.Handler) HPRouter
	HandleFunc(string, func(http.ResponseWriter, *http.Request)) HPRouter

	Methods(...string) HPRouter
	PathPrefix(string) HPRouter
}

type WSRouter interface {
	Route
	Handle(string, http.Handler) WSRouter
	HandleFunc(string, func(http.ResponseWriter, *http.Request)) WSRouter
}

type Router struct {
	HPRouter
	WSRouter
}
