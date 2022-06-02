package routers

import (
	"github.com/gin-gonic/gin"

	"gin-demo/apps/web/blogs"
	"gin-demo/apps/web/users"
)

var routerFuncs = []func(*gin.RouterGroup){
	users.Routers,
	blogs.Routers,
}

// Register 注册router
func Register(e *gin.Engine) {
	registerV1(e)
}

func registerV1(e *gin.Engine) {
	rg := e.Group("/apps/v1")

	for _, routerFunc := range routerFuncs {
		routerFunc(rg)
	}
}
