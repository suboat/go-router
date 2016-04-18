package gorouter

import "net/http"

type WSHandlerFunc func(Request) Response

type WSHandler interface {
	ServeWS(Request) Response
}

type HandleBind interface {
	http.Handler
	WSHandler
}
