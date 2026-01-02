package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORS 创建跨域资源共享中间件。
//
// 默认配置：
//   - Allow-Origin: *
//   - Allow-Methods: POST, GET, PUT, DELETE, PATCH, OPTIONS
//   - Allow-Headers: Content-Type, Authorization, X-Requested-With
//   - Allow-Credentials: true
//
// 对于 OPTIONS 预检请求，直接返回 204 No Content。
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, "+
				"Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods",
			"POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
