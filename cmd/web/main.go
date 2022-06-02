package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"gin-demo/internal/web/configs"
	"gin-demo/internal/web/logger"
	"gin-demo/internal/web/middlewares"
	"gin-demo/internal/web/routers"
)

func main() {
	configs.LoadConfig()

	logger.InitLog()

	e := gin.New()

	middlewares.Register(e)

	routers.Register(e)

	if err := e.Run(fmt.Sprintf("0.0.0.0:%d", configs.APPConfig.Port)); err != nil {
		panic(fmt.Sprintf("gin run failed, %s", err.Error()))
	}
}
