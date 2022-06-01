package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"gin-demo/env"
	"gin-demo/logger"
	"gin-demo/middlewares"
	"gin-demo/router"
)

func main() {
	env.LoadConfig()

	logger.InitLog()

	e := gin.New()

	middlewares.Register(e)

	router.Register(e)

	if err := e.Run(fmt.Sprintf("0.0.0.0:%d", env.APPConfig.Port)); err != nil {
		panic(fmt.Sprintf("gin run failed, %s", err.Error()))
	}
}
