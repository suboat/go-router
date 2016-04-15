package mux

import (
	"github.com/gorilla/mux"
	gr "github.com/suboat/go-router"
	"net/http"
)

type MuxRouter struct {
	gr.HTTPRouter
	Router *mux.Router
}

func NewMuxRouter() *MuxRouter {
	return &MuxRouter{Router: mux.NewRouter()}
}

func (r *MuxRouter) newRoute() *mux.Route {
	return r.Router.NewRoute()
}

func (r *MuxRouter) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, r.Router)
}

func (r *MuxRouter) Handle(path string, handler http.Handler) *mux.Route {
	return r.newRoute().Path(path).Handler(handler)
}

func (r *MuxRouter) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	return r.newRoute().Path(path).HandlerFunc(f)
}

func (r *MuxRouter) Methods(methods ...string) *mux.Route {
	return r.newRoute().Methods(methods...)
}

func (r *MuxRouter) Path(tpl string) *mux.Route {
	return r.newRoute().Path(tpl)
}

func (r *MuxRouter) PathPrefix(tpl string) *MuxRouter {
	r.Router = r.newRoute().PathPrefix(tpl).Subrouter()
	return r
}
