package router

import (
	"github.com/gorilla/mux"
	. "github.com/suboat/go-router"
	"net/http"
)

type MuxRouter struct {
	route  *mux.Route
	Router *mux.Router
	err    error
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
	r.err = http.ListenAndServe(addr, r.Router)
	return r.err
}

func (r *MuxRouter) Error() error {
	return r.err
}

func (r *MuxRouter) Handle(path string, handler http.Handler) HPRouter {
	r.newRoute().Path(path).Handler(handler)
	return r
}

func (r *MuxRouter) HandleFunc(path string, handler func(http.ResponseWriter, *http.Request)) HPRouter {
	r.newRoute().Path(path).HandlerFunc(handler)
	return r
}

func (r *MuxRouter) Methods(methods ...string) HPRouter {
	r.getRoute().Methods(methods...)
	return r
}

func (r *MuxRouter) PathPrefix(tpl string) HPRouter {
	return newMuxRouter(r.newRoute().PathPrefix(tpl).Subrouter())
}
