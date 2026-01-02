// Package bootstrap 提供 Gin Engine 和 HTTP Server 启动逻辑。
package bootstrap

import (
	"io"

	"github.com/gin-gonic/gin"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/platform/http/middleware"
)

func init() {
	// 设置 Gin 为 Release 模式，禁用 debug 日志（必须在包初始化时设置）
	gin.SetMode(gin.ReleaseMode)
	// 禁用 Gin 的默认日志输出
	gin.DefaultWriter = io.Discard
	gin.DefaultErrorWriter = io.Discard
}

// NewEngine 创建 Gin Engine，注册全局中间件
func NewEngine() *gin.Engine {
	engine := gin.New()

	// 全局中间件
	engine.Use(gin.Recovery())         // Panic 恢复
	engine.Use(middleware.Logger())    // 请求日志（基于 slog）
	engine.Use(middleware.RequestID()) // 请求 ID
	engine.Use(middleware.CORS())      // CORS 支持

	return engine
}
