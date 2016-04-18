package router

import (
	"github.com/gin-gonic/gin"
	. "github.com/suboat/go-router"
	"net/http"
)

type GinRouter struct {
	Router *gin.RouterGroup
	err    error
}

func newGinRouter(r *gin.RouterGroup) *GinRouter {
	return &GinRouter{Router: r}
}

func NewGinRouter(r *gin.Engine) *GinRouter {
	return newGinRouter(r.RouterGroup)
}

func (r *GinRouter) newRoute() *mux.Route {
	r.route = r.Router.NewRoute()
	return r.route
}

func (r *GinRouter) getRoute() *mux.Route {
	if r.route == nil {
		return r.newRoute()
	}
	return r.route
}

func (r *GinRouter) ListenAndServe(addr string) error {
	r.err = http.ListenAndServe(addr, r.Router)
	return r.err
}

func (r *GinRouter) Error() error {
	return r.err
}

func (r *GinRouter) Handle(path string, handler http.Handler) HTTPRoute {
	r.newRoute().Path(path).Handler(handler)
	return r
}

func (r *GinRouter) HandleFunc(path string, handler func(http.ResponseWriter, *http.Request)) HTTPRoute {
	r.newRoute().Path(path).HandlerFunc(handler)
	return r
}

func (r *GinRouter) Methods(methods ...string) HTTPRoute {
	r.getRoute().Methods(methods...)
	return r
}

func (r *GinRouter) PathPrefix(tpl string) HTTPRoute {
	return newGinRouter(r.newRoute().PathPrefix(tpl).Subrouter())
}
