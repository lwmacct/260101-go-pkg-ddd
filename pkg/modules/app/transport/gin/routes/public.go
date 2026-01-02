package routes

import (
	"github.com/lwmacct/260101-go-pkg-gin/pkg/routes"

	corehandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/app/transport/gin/handler"
)

// Public 公开路由（健康检查）
func Public(healthHandler *corehandler.HealthHandler) []routes.Route {
	return []routes.Route{
		{
			Method:      routes.GET,
			Path:        "/health",
			Handler:     healthHandler.Check,
			Operation:   "public:health:check",
			Tags:        "System",
			Summary:     "健康检查",
			Description: "系统健康状态检查",
		},
	}
}
