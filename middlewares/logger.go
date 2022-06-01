package middlewares

import (
	"math"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"gin-demo/logger"
)

// logger record access log
// refer from https://www.cnblogs.com/xinliangcoder/p/11212573.html
func loggerM() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		info := map[string]interface{}{
			"clientip":   c.ClientIP(),
			"reqmethod":  c.Request.Method,
			"requri":     c.Request.RequestURI,
			"referer":    c.Request.Referer(),
			"useragent":  c.Request.UserAgent(),
			"statuscode": c.Writer.Status(),
			"datalength": dataLength,
			"cost":       int(math.Ceil(float64(time.Since(startTime).Nanoseconds()) / 1000000.0)),
		}

		// 日志格式
		logger.L.WithFields(
			logrus.Fields(info),
		).Info("-")
	}
}
