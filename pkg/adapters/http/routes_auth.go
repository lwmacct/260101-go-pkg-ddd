package http

import (
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/adapters/http/permission"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/adapters/http/routes"
)

// authRoutes 返回 Auth 域路由（两步验证）
func (deps *RouterDependencies) authRoutes() []routes.Route {
	// Auth 域中间件模式
	baseAuthMiddlewares := []routes.MiddlewareConfig{
		{Name: routes.MiddlewareRequestID},
		{Name: routes.MiddlewareOperationID},
		{Name: routes.MiddlewareAuth},
		{Name: routes.MiddlewareRBAC},
	}

	auditAuthMiddlewares := append(cloneMiddlewares(baseAuthMiddlewares), routes.MiddlewareConfig{Name: routes.MiddlewareAudit})

	return []routes.Route{
		{
			Method:      routes.POST,
			Path:        "/api/auth/2fa/setup",
			Handler:     deps.TwoFAHandler.Setup,
			Op:          permission.Self2FASetup,
			Middlewares: auditAuthMiddlewares,
			Tags:        "Authentication - 2FA",
			Summary:     "设置 2FA",
			Description: "设置两步验证",
		},
		{
			Method:      routes.POST,
			Path:        "/api/auth/2fa/verify",
			Handler:     deps.TwoFAHandler.VerifyAndEnable,
			Op:          permission.Self2FAVerify,
			Middlewares: auditAuthMiddlewares,
			Tags:        "Authentication - 2FA",
			Summary:     "启用 2FA",
			Description: "验证并启用两步验证",
		},
		{
			Method:      routes.POST,
			Path:        "/api/auth/2fa/disable",
			Handler:     deps.TwoFAHandler.Disable,
			Op:          permission.Self2FADisable,
			Middlewares: auditAuthMiddlewares,
			Tags:        "Authentication - 2FA",
			Summary:     "禁用 2FA",
			Description: "禁用两步验证",
		},
		{
			Method:      routes.GET,
			Path:        "/api/auth/2fa/status",
			Handler:     deps.TwoFAHandler.GetStatus,
			Op:          permission.Self2FAStatus,
			Middlewares: baseAuthMiddlewares,
			Tags:        "Authentication - 2FA",
			Summary:     "2FA 状态",
			Description: "获取两步验证状态",
		},
	}
}
