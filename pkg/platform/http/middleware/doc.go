// Package middleware 提供 HTTP 中间件实现。
//
// 本包提供了 Gin 框架的通用中间件：
//   - [Logger]: HTTP 请求日志记录
//   - [RequestID]: 请求 ID 生成
//   - [OperationID]: 操作 ID 设置
//   - [Auth]: JWT 认证
//   - [RBAC]: 基于角色的访问控制
//   - [Audit]: 审计日志
//   - [OrgContext]: 多租户上下文（组织/团队）
//
// 使用方式：
//
//	engine := gin.New()
//	engine.Use(middleware.Logger())
//	engine.Use(middleware.RequestID())
//
//	// 路由级别中间件
//	engine.GET("/api/users", middleware.Auth(jwtManager), handler.GetUsers)
//
// 并发安全性：
//   - 所有中间件都是并发安全的。
//
// 依赖注入：
//   - Auth 和 RBAC 中间件需要通过工厂函数注入依赖（见 internal/app/di/）
package middleware
