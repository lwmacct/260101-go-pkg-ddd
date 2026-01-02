package gin

import (
	iamhandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/iam/transport/gin/handler"
	platformhttp "github.com/lwmacct/260101-go-pkg-ddd/pkg/platform/http"
)

// AuthRoutes Auth 域路由（两步验证）
func AuthRoutes(twoFAHandler *iamhandler.TwoFAHandler) []platformhttp.Route {
	baseMiddlewares := []platformhttp.MiddlewareConfig{
		{Name: platformhttp.MiddlewareRequestID},
		{Name: platformhttp.MiddlewareOperationID},
		{Name: platformhttp.MiddlewareAuth},
		{Name: platformhttp.MiddlewareRBAC},
	}

	auditMiddlewares := append(baseMiddlewares, platformhttp.MiddlewareConfig{
		Name: platformhttp.MiddlewareAudit,
	})

	return []platformhttp.Route{
		{
			Method:      platformhttp.POST,
			Path:        "/api/auth/2fa/setup",
			Handler:     twoFAHandler.Setup,
			Op:          "self:2fa:setup",
			Middlewares: auditMiddlewares,
			Tags:        "Authentication - 2FA",
			Summary:     "设置 2FA",
			Description: "设置两步验证",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/auth/2fa/verify-enable",
			Handler:     twoFAHandler.VerifyAndEnable,
			Op:          "self:2fa:enable",
			Middlewares: auditMiddlewares,
			Tags:        "Authentication - 2FA",
			Summary:     "启用 2FA",
			Description: "验证并启用两步验证",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/auth/2fa/disable",
			Handler:     twoFAHandler.Disable,
			Op:          "self:2fa:disable",
			Middlewares: auditMiddlewares,
			Tags:        "Authentication - 2FA",
			Summary:     "禁用 2FA",
			Description: "禁用两步验证",
		},
		{
			Method:      platformhttp.GET,
			Path:        "/api/auth/2fa/status",
			Handler:     twoFAHandler.GetStatus,
			Op:          "self:2fa:status",
			Middlewares: baseMiddlewares,
			Tags:        "Authentication - 2FA",
			Summary:     "2FA 状态",
			Description: "获取两步验证状态",
		},
	}
}
