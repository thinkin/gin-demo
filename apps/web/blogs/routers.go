package blogs

import "github.com/gin-gonic/gin"

// Routers 路由
func Routers(rg *gin.RouterGroup) {
	r := rg.Group("/blogs")
	{
		r.GET("", List)
		r.POST("", Add)
		r.GET("/:id", Get)
		r.DELETE("/:id", Delete)
		r.PATCH("/:id", Update)
	}
}
