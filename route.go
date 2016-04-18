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
	HandleBind(string, HandleBind) HTTPRoute

	Methods(...string) HTTPRoute
	PathPrefix(string) HTTPRoute
}

type WSRoute interface {
	Route
	Bind(string) WSRoute
	Handle(string, WSHandler) WSRoute
	HandleBind(string, HandleBind) WSRoute
}

type Router struct {
	HTTPRoute
	WSRoute
}
