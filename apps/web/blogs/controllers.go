package blogs

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-demo/pkg/response"
)

func List(c *gin.Context) {
	c.JSON(http.StatusOK, response.Success("list"))
}

func Add(c *gin.Context) {
	c.JSON(http.StatusOK, response.Success("add"))
}

func Get(c *gin.Context) {
	blogID := c.GetInt("id")

	c.JSON(http.StatusOK, response.Success(blogID))
}

func Update(c *gin.Context) {
	blogID := c.GetInt("id")

	c.JSON(http.StatusOK, response.Success(blogID))
}

func Delete(c *gin.Context) {
	blogID := c.GetInt("id")

	c.JSON(http.StatusOK, response.Success(blogID))
}
