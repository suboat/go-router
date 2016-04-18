package gorouter

import "net/http"

type WSHandler interface {
	ServeWS(Request) Response
}

type HandleBind interface {
	http.Handler
	WSHandler
}

type WSHandlerFunc func(Request) Response

func (f WSHandlerFunc) ServeWS(r Request) Response {
	return f(r)
}
