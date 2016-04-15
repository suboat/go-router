package gorouter

import "net/http"

type Route interface {
	ListenAndServe(string, http.Handler) error

	Handle(string, http.Handler) *Route
	HandleFunc(string, func(http.ResponseWriter, *http.Request)) *Route

	Methods(...string) *Route

	Path(string) *Route
	PathPrefix(string) *Route
}

type HTTPRouter interface {
	Route
}

type WSRouter interface {
	Route
}

type Router struct {
	HTTPRouter
	WSRouter
}
