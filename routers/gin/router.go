package router

import (
	"github.com/gin-gonic/gin"
	. "github.com/suboat/go-router"
	"net/http"
)

type GinRouter struct {
	Engine *gin.Engine
	Router *gin.RouterGroup
	route  *ginRoute
	err    error
}

func newGinRouter(e *gin.Engine, r *gin.RouterGroup) *GinRouter {
	return &GinRouter{Engine: e, Router: r}
}

func NewGinRouterWithEngine(r *gin.Engine) *GinRouter {
	return newGinRouter(r, &r.RouterGroup)
}

func NewGinRouter() *GinRouter {
	return NewGinRouterWithEngine(gin.New())
}

func (s *GinRouter) GetEngine() *gin.Engine {
	return s.Engine
}

type ginRoute struct {
	Router  *gin.RouterGroup
	Path    string
	Handler gin.HandlerFunc
}

func (s *ginRoute) Handle(method string) {
	if len(s.Path) != 0 {
		s.Router.Handle(method, s.Path, s.Handler)
	}
}

func (r *GinRouter) newRoute(path string, handler gin.HandlerFunc) *ginRoute {
	r.route = &ginRoute{Router: r.Router, Path: path, Handler: handler}
	return r.route
}

func (r *GinRouter) getRoute() *ginRoute {
	if r.route == nil {
		return &ginRoute{}
	}
	return r.route
}

func (r *GinRouter) ListenAndServe(addr string) error {
	r.err = r.Engine.Run(addr)
	return r.err
}

func (r *GinRouter) Error() error {
	return r.err
}

func (r *GinRouter) Handle(path string, handler http.Handler) HTTPRoute {
	r.HandleGin(path, gin.WrapH(handler))
	return r
}

func (r *GinRouter) HandleFunc(path string, handler func(http.ResponseWriter, *http.Request)) HTTPRoute {
	r.HandleGin(path, gin.WrapF(handler))
	return r
}

func (r *GinRouter) HandleBind(path string, handler HandleBind) HTTPRoute {
	return r.Handle(path, handler)
}

func (r *GinRouter) Methods(methods ...string) HTTPRoute {
	for _, method := range methods {
		r.getRoute().Handle(method)
	}
	return r
}

func (r *GinRouter) PathPrefix(tpl string) HTTPRoute {
	return newGinRouter(r.Engine, r.Router.Group(tpl))
}
