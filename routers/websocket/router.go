package router

import (
	. "github.com/suboat/go-router"
	"net/http"
)

type WebSocketRouter struct {
	Router HTTPRouter
	err    error
}

func newMuxRouter(r HTTPRouter) *WebSocketRouter {
	return &WebSocketRouter{Router: r}
}

func NewWSRouter(r HTTPRouter) *WebSocketRouter {
	ws := newMuxRouter(r)
	if r == nil {
		ws.err = ErrRouter
	}
	return ws
}

func (r *WebSocketRouter) ListenAndServe(addr string) error {
	r.err = r.Router.ListenAndServe(addr)
	return r.err
}

func (r *WebSocketRouter) Error() error {
	return r.err
}

func (r *WebSocketRouter) Handle(path string, handler http.Handler) WSRouter {
	r.Router.Handle(path, handler)
	return r
}

func (r *WebSocketRouter) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) WSRouter {
	r.Router.HandleFunc(path, f)
	return r
}

func (r *WebSocketRouter) Methods(methods ...string) WSRouter {
	r.Router.Methods(methods...)
	return r
}

func (r *WebSocketRouter) PathPrefix(tpl string) WSRouter {
	return newMuxRouter(r.Router.PathPrefix(tpl))
}
