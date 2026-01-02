// Package bootstrap 提供 Gin Engine 和 HTTP Server 启动逻辑。
package bootstrap

import (
	"github.com/gin-gonic/gin"
)

// NewEngine 创建 Gin Engine，注册全局中间件
func NewEngine() *gin.Engine {
	engine := gin.New()
	// TODO: 添加 telemetry 中间件
	engine.Use(gin.Recovery())
	// TODO: 添加其他全局中间件
	return engine
}
