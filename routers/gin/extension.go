package router

import (
	"github.com/gin-gonic/gin"
	. "github.com/suboat/go-router"
)

func (r *GinRouter) POST(relativePath string, handlers ...gin.HandlerFunc) HTTPRoute {
	r.Router.Handle("POST", relativePath, handlers...)
	return r
}

func (r *GinRouter) GET(relativePath string, handlers ...gin.HandlerFunc) HTTPRoute {
	r.Router.Handle("GET", relativePath, handlers...)
	return r
}

func (r *GinRouter) DELETE(relativePath string, handlers ...gin.HandlerFunc) HTTPRoute {
	r.Router.Handle("DELETE", relativePath, handlers...)
	return r
}

func (r *GinRouter) PATCH(relativePath string, handlers ...gin.HandlerFunc) HTTPRoute {
	r.Router.Handle("PATCH", relativePath, handlers...)
	return r
}

func (r *GinRouter) PUT(relativePath string, handlers ...gin.HandlerFunc) HTTPRoute {
	r.Router.Handle("PUT", relativePath, handlers...)
	return r
}

func (r *GinRouter) OPTIONS(relativePath string, handlers ...gin.HandlerFunc) HTTPRoute {
	r.Router.Handle("OPTIONS", relativePath, handlers...)
	return r
}

func (r *GinRouter) HEAD(relativePath string, handlers ...gin.HandlerFunc) HTTPRoute {
	r.Router.Handle("HEAD", relativePath, handlers...)
	return r
}

func (r *GinRouter) Any(relativePath string, handlers ...gin.HandlerFunc) HTTPRoute {
	r.Router.Handle("GET", relativePath, handlers...)
	r.Router.Handle("POST", relativePath, handlers...)
	r.Router.Handle("PUT", relativePath, handlers...)
	r.Router.Handle("PATCH", relativePath, handlers...)
	r.Router.Handle("HEAD", relativePath, handlers...)
	r.Router.Handle("OPTIONS", relativePath, handlers...)
	r.Router.Handle("DELETE", relativePath, handlers...)
	r.Router.Handle("CONNECT", relativePath, handlers...)
	r.Router.Handle("TRACE", relativePath, handlers...)
	return r
}
