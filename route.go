package gorouter

import "net/http"

type Route interface {
	ListenAndServe(string) error
	Error() error
}

type HTTPRoute interface {
	Route
	Handle(string, http.Handler) HTTPRoute
	HandleFunc(string, func(http.ResponseWriter, *http.Request)) HTTPRoute

	Methods(...string) HTTPRoute
	PathPrefix(string) HTTPRoute
}

type WSRoute interface {
	Route
	Handle(string) WSRoute
	//HandleBind(string, f) WSRoute
}

type Router struct {
	HTTPRoute
	WSRoute
}
