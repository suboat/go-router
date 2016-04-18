package gorouter

import "net/http"

type Route interface {
	ListenAndServe(string) error
	Error() error
}

type HTTPRouter interface {
	Route
	Handle(string, http.Handler) HTTPRouter
	HandleFunc(string, func(http.ResponseWriter, *http.Request)) HTTPRouter

	Methods(...string) HTTPRouter
	PathPrefix(string) HTTPRouter
}

type WSRouter interface {
	Route
	Handle(string, http.Handler) WSRouter
	HandleFunc(string, func(http.ResponseWriter, *http.Request)) WSRouter

	Methods(...string) WSRouter
	PathPrefix(string) WSRouter
}

type Router struct {
	HTTPRouter
	WSRouter
}
