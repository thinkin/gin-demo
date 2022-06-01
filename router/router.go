package router

import (
	"github.com/gin-gonic/gin"

	"gin-demo/app/users"
)

var routerFuncs = []func(*gin.RouterGroup){
	users.Routers,
}

// Register 注册router
func Register(e *gin.Engine) {
	registerV1(e)
}

func registerV1(e *gin.Engine) {
	rg := e.Group("/api/v1")

	for _, routerFunc := range routerFuncs {
		routerFunc(rg)
	}
}
