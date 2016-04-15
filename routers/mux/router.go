package mux

import (
	"github.com/gorilla/mux"
	gr "github.com/suboat/go-router"
	"net/http"
)

type MuxRouter struct {
	gr.HTTPRouter
	Router *mux.Router
	route  *mux.Route
}

func newMuxRouter(r *mux.Router) *MuxRouter {
	return &MuxRouter{Router: r}
}

func NewMuxRouter() *MuxRouter {
	return newMuxRouter(mux.NewRouter())
}

func (r *MuxRouter) newRoute() *mux.Route {
	r.route = r.Router.NewRoute()
	return r.route
}

func (r *MuxRouter) getRoute() *mux.Route {
	if r.route == nil {
		return r.newRoute()
	}
	return r.route
}

func (r *MuxRouter) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, r.Router)
}

func (r *MuxRouter) Handle(path string, handler http.Handler) *MuxRouter {
	r.newRoute().Path(path).Handler(handler)
	return r
}

func (r *MuxRouter) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *MuxRouter {
	r.newRoute().Path(path).HandlerFunc(f)
	return r
}

func (r *MuxRouter) Methods(methods ...string) *MuxRouter {
	r.getRoute().Methods(methods...)
	return r
}

func (r *MuxRouter) PathPrefix(tpl string) *MuxRouter {
	return newMuxRouter(r.newRoute().PathPrefix(tpl).Subrouter())
}
