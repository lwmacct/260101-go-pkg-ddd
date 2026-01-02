package gin

import (
	corehandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/core/transport/gin/handler"
	platformhttp "github.com/lwmacct/260101-go-pkg-ddd/pkg/platform/http"
)

// PublicRoutes 公开路由（健康检查）
func PublicRoutes(healthHandler *corehandler.HealthHandler) []platformhttp.Route {
	baseMiddlewares := []platformhttp.MiddlewareConfig{
		{Name: platformhttp.MiddlewareRequestID},
	}

	return []platformhttp.Route{
		{
			Method:      platformhttp.GET,
			Path:        "/health",
			Handler:     healthHandler.Check,
			Op:          "public:health:check",
			Middlewares: baseMiddlewares,
			Tags:        "System",
			Summary:     "健康检查",
			Description: "系统健康状态检查",
		},
	}
}
