package users

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gin-demo/pkg/common/response"
)

func List(c *gin.Context) {
	c.JSON(http.StatusOK, response.Success("list"))
}

func Add(c *gin.Context) {
	c.JSON(http.StatusOK, response.Success("add"))
}

func Get(c *gin.Context) {
	userID := c.GetInt("id")

	c.JSON(http.StatusOK, response.Success(userID))
}

func Update(c *gin.Context) {
	userID := c.GetInt("id")

	c.JSON(http.StatusOK, response.Success(userID))
}

func Delete(c *gin.Context) {
	userID := c.GetInt("id")

	c.JSON(http.StatusOK, response.Success(userID))
}
