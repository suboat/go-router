package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *GinRouter) HandleGin(path string, handler gin.HandlerFunc) *GinRouter {
	r.newRoute(path, handler)
	return r
}

func (r *GinRouter) Group(relativePath string) *GinRouter {
	return newGinRouter(r.Engine, r.Router.Group(relativePath))
}

func (r *GinRouter) Use(middleware ...gin.HandlerFunc) *GinRouter {
	r.Engine.Use(middleware...)
	return r
}

func (r *GinRouter) Static(relativePath, root string) *GinRouter {
	r.Engine.Static(relativePath, root)
	return r
}

func (r *GinRouter) StaticFile(relativePath, filepath string) *GinRouter {
	r.Engine.StaticFile(relativePath, filepath)
	return r
}

func (r *GinRouter) StaticFS(relativePath string, fs http.FileSystem) *GinRouter {
	r.Engine.StaticFS(relativePath, fs)
	return r
}
