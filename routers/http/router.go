package router

import (
	. "github.com/suboat/go-router"
	"net/http"
	"path"
)

type HTTPRouter struct {
	Router *http.ServeMux
	Tpl    string
	err    error
}

func newHTTPRouter() *HTTPRouter {
	return &HTTPRouter{Router: http.NewServeMux()}
}

func NewHTTPRouter() *HTTPRouter {
	return newHTTPRouter()
}

func (r *HTTPRouter) ListenAndServe(addr string) error {
	r.err = http.ListenAndServe(addr, r.Router)
	return r.err
}

func (r *HTTPRouter) Error() error {
	return r.err
}

func (r *HTTPRouter) Handle(_path string, handler http.Handler) HPRouter {
	r.Router.Handle(path.Join(r.Tpl, _path), handler)
	return r
}

func (r *HTTPRouter) HandleFunc(_path string, handler func(http.ResponseWriter, *http.Request)) HPRouter {
	r.Router.HandleFunc(path.Join(r.Tpl, _path), handler)
	return r
}

func (r *HTTPRouter) Methods(methods ...string) HPRouter {
	return r
}

func (r *HTTPRouter) PathPrefix(tpl string) HPRouter {
	r.Tpl = tpl
	return r
}
