package http

import "github.com/gin-gonic/gin"

// Method HTTP 方法类型
type Method string

const (
	GET     Method = "GET"
	POST    Method = "POST"
	PUT     Method = "PUT"
	DELETE  Method = "DELETE"
	PATCH   Method = "PATCH"
	HEAD    Method = "HEAD"
	OPTIONS Method = "OPTIONS"
)

// MiddlewareConfig 中间件配置
type MiddlewareConfig struct {
	Name MiddlewareType
	// 可扩展：配置参数等
}

// Route 路由定义（声明式）
type Route struct {
	// 基本信息
	Method      Method
	Path        string
	Handler     gin.HandlerFunc
	Op          string // operation: domain:resource:action
	Middlewares []MiddlewareConfig

	// Swagger 文档
	Tags        string
	Summary     string
	Description string
}
