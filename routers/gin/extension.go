package router

import (
	"github.com/gin-gonic/gin"
	. "github.com/suboat/go-router"
	"net/http"
)

func (r *GinRouter) Use(middleware ...gin.HandlerFunc) HTTPRoute {
	r.Engine.Use(middleware...)
	return r
}

func (r *GinRouter) Static(relativePath, root string) HTTPRoute {
	r.Engine.Static(relativePath, root)
	return r
}

func (r *GinRouter) StaticFile(relativePath, filepath string) HTTPRoute {
	r.Engine.StaticFile(relativePath, filepath)
	return r
}

func (r *GinRouter) StaticFS(relativePath string, fs http.FileSystem) HTTPRoute {
	r.Engine.StaticFS(relativePath, fs)
	return r
}
