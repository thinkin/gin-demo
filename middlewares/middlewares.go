package middlewares

import "github.com/gin-gonic/gin"

// Register middlewares
func Register(e *gin.Engine) {
	e.Use(loggerM())
}
