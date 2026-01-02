package gin

import (
	iamhandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/iam/transport/gin/handler"
	platformhttp "github.com/lwmacct/260101-go-pkg-ddd/pkg/platform/http"
)

// PublicRoutes 公开路由（无需认证）
func PublicRoutes(authHandler *iamhandler.AuthHandler) []platformhttp.Route {
	baseMiddlewares := []platformhttp.MiddlewareConfig{
		{Name: platformhttp.MiddlewareRequestID},
		{Name: platformhttp.MiddlewareOperationID},
		{Name: platformhttp.MiddlewareAudit},
	}

	return []platformhttp.Route{
		{
			Method:      platformhttp.POST,
			Path:        "/api/auth/register",
			Handler:     authHandler.Register,
			Op:          "public:auth:register",
			Middlewares: baseMiddlewares,
			Tags:        "Authentication",
			Summary:     "注册",
			Description: "用户注册",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/auth/login",
			Handler:     authHandler.Login,
			Op:          "public:auth:login",
			Middlewares: baseMiddlewares,
			Tags:        "Authentication",
			Summary:     "登录",
			Description: "用户登录",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/auth/login/2fa",
			Handler:     authHandler.Login2FA,
			Op:          "public:auth:login2fa",
			Middlewares: baseMiddlewares,
			Tags:        "Authentication",
			Summary:     "2FA 登录",
			Description: "两步验证登录",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/auth/refresh",
			Handler:     authHandler.RefreshToken,
			Op:          "public:auth:refresh",
			Middlewares: baseMiddlewares,
			Tags:        "Authentication",
			Summary:     "刷新令牌",
			Description: "刷新访问令牌",
		},
	}
}
