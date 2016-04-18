package gorouter

import "net/http"

type HandlerFunc func(Request) Response

type WSHandler interface {
	ServeWS(Request) Response
}

type Handle interface {
	http.Handler
	WSHandler
}
