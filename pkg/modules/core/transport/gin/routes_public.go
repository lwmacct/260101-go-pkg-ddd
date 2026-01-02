package gin

import (
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/core/transport/gin/routes"
)

// publicRoutes 返回公开域路由（无需认证）
func (deps *RouterDependencies) publicRoutes() []routes.Route {
	return []routes.Route{
		{
			Method:  routes.POST,
			Path:    "/api/auth/register",
			Handler: deps.AuthHandler.Register,
			Op:      "public:auth:register",
			Middlewares: []routes.MiddlewareConfig{
				{Name: routes.MiddlewareRequestID},
				{Name: routes.MiddlewareOperationID},
				{Name: routes.MiddlewareAudit},
			},
			Tags:        "Authentication",
			Summary:     "注册",
			Description: "用户注册",
		},
		{
			Method:  routes.POST,
			Path:    "/api/auth/login",
			Handler: deps.AuthHandler.Login,
			Op:      "public:auth:login",
			Middlewares: []routes.MiddlewareConfig{
				{Name: routes.MiddlewareRequestID},
				{Name: routes.MiddlewareOperationID},
				{Name: routes.MiddlewareAudit},
			},
			Tags:        "Authentication",
			Summary:     "登录",
			Description: "用户登录",
		},
		{
			Method:  routes.POST,
			Path:    "/api/auth/login/2fa",
			Handler: deps.AuthHandler.Login2FA,
			Op:      "public:auth:login2fa",
			Middlewares: []routes.MiddlewareConfig{
				{Name: routes.MiddlewareRequestID},
				{Name: routes.MiddlewareOperationID},
				{Name: routes.MiddlewareAudit},
			},
			Tags:        "Authentication",
			Summary:     "2FA 登录",
			Description: "两步验证登录",
		},
		{
			Method:  routes.POST,
			Path:    "/api/auth/refresh",
			Handler: deps.AuthHandler.RefreshToken,
			Op:      "public:auth:refresh",
			Middlewares: []routes.MiddlewareConfig{
				{Name: routes.MiddlewareRequestID},
				{Name: routes.MiddlewareOperationID},
				{Name: routes.MiddlewareAudit},
			},
			Tags:        "Authentication",
			Summary:     "刷新令牌",
			Description: "刷新访问令牌",
		},
		{
			Method:  routes.GET,
			Path:    "/api/auth/captcha",
			Handler: deps.CaptchaHandler.GetCaptcha,
			Op:      "public:auth:captcha",
			Middlewares: []routes.MiddlewareConfig{
				{Name: routes.MiddlewareRequestID},
				{Name: routes.MiddlewareOperationID},
			},
			Tags:        "Authentication",
			Summary:     "获取验证码",
			Description: "获取图形验证码",
		},
	}
}
