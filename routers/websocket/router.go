package router

import (
	. "github.com/suboat/go-router"
	"net/http"
)

type WebSocketRouter struct {
	Router HTTPRoute
	err    error
}

func newWebSocketRouter(r HTTPRoute) *WebSocketRouter {
	return &WebSocketRouter{Router: r}
}

func NewWebSocketRouter(r HTTPRoute) *WebSocketRouter {
	ws := newWebSocketRouter(r)
	if r == nil {
		ws.err = ErrRouter
	}
	return ws
}

func (r *WebSocketRouter) ListenAndServe(addr string) error {
	r.err = http.ListenAndServe(addr, nil)
	return r.err
}

func (r *WebSocketRouter) Error() error {
	return r.err
}

func (r *WebSocketRouter) Handle(path string, handler http.Handler) WSRoute {
	http.Handle(path, handler)
	return r
}

func (r *WebSocketRouter) HandleFunc(path string, handler func(http.ResponseWriter, *http.Request)) WSRoute {
	http.HandleFunc(path, handler)
	return r
}
