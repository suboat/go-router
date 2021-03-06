package router

import (
	. "github.com/suboat/go-router"
	"net/http"
)

type WSRouter struct {
	Handler *WSHandle
	err     error
}

func NewWSRouter(h *WSHandle) *WSRouter {
	ws := &WSRouter{Handler: h}
	if h == nil {
		ws.err = ErrHandle
	}
	return ws
}

func (r *WSRouter) ListenAndServe(addr string) error {
	r.err = http.ListenAndServe(addr, nil)
	return r.err
}

func (r *WSRouter) Error() error {
	if r.err != nil {
		return r.err
	}
	return r.Handler.Error()
}

func (r *WSRouter) Bind(path string) WSRoute {
	http.Handle(path, r.Handler)
	return r
}

// TODO:未完成，实现ServerMux
func (r *WSRouter) Handle(path string, handler WSHandler) WSRoute {
	return r
}

func (r *WSRouter) HandleBind(path string, handler HandleBind) WSRoute {
	return r.Handle(path, handler)
}
