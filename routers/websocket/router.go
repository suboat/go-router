package websocket

import (
	gr "github.com/suboat/go-router"
	"net/http"
)

type WSRouter struct {
	gr.WSRouter
	Router gr.HTTPRouter
}

func newMuxRouter(r gr.HTTPRouter) *WSRouter {
	return &WSRouter{Router: r}
}

func NewWSRouter(r gr.HTTPRouter) (*WSRouter, error) {
	if r == nil {
		return nil, gr.ErrRouter
	}
	return newMuxRouter(r), nil
}

func (r *WSRouter) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, r.Router)
}

func (r *WSRouter) Handle(path string, handler http.Handler) *WSRouter {
	r.Router.Handle(path, handler)
	return r
}

func (r *WSRouter) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *WSRouter {
	r.Router.HandleFunc(path, f)
	return r
}

func (r *WSRouter) Methods(methods ...string) *WSRouter {
	r.Router.Methods(methods...)
	return r
}

func (r *WSRouter) PathPrefix(tpl string) *WSRouter {
	return newMuxRouter(r.Router.PathPrefix(tpl))
}
