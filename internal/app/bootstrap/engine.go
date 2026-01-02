// Package bootstrap 提供 Gin Engine 和 HTTP Server 启动逻辑。
package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/platform/http/middleware"
)

// NewEngine 创建 Gin Engine，注册全局中间件
func NewEngine() *gin.Engine {
	engine := gin.New()

	// 全局中间件
	engine.Use(gin.Recovery())     // Panic 恢复
	engine.Use(middleware.Logger())   // 请求日志
	engine.Use(middleware.RequestID()) // 请求 ID

	// TODO: 添加 telemetry 中间件
	// TODO: 添加 CORS 中间件

	return engine
}
