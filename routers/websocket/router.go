package router

import (
	. "github.com/suboat/go-router"
	"net/http"
)

type WSRouter struct {
	Router HTTPRoute
	err    error
}

func newWSRouter(r HTTPRoute) *WSRouter {
	return &WSRouter{Router: r}
}

func NewWSRouter(r HTTPRoute) *WSRouter {
	ws := newWSRouter(r)
	if r == nil {
		ws.err = ErrRouter
	}
	return ws
}

func (r *WSRouter) ListenAndServe(addr string) error {
	r.err = http.ListenAndServe(addr, nil)
	return r.err
}

func (r *WSRouter) Error() error {
	return r.err
}

func (r *WSRouter) Handle(path string, handler http.Handler) WSRoute {
	http.Handle(path, handler)
	return r
}

func (r *WSRouter) HandleFunc(path string, handler func(http.ResponseWriter, *http.Request)) WSRoute {
	http.HandleFunc(path, handler)
	return r
}
