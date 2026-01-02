// Package routes 定义 Core 模块的 HTTP 路由。
//
// Core 模块包含：健康检查、系统管理（组织、任务、审计、缓存、概览）
package routes

import (
	"github.com/lwmacct/260101-go-pkg-gin/pkg/routes"

	corehandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/app/transport/gin/handler"
)

// All 返回 Core 域的所有路由
func All(
	healthHandler *corehandler.HealthHandler,
	orgHandler *corehandler.OrgHandler,
	orgMemberHandler *corehandler.OrgMemberHandler,
	teamHandler *corehandler.TeamHandler,
	teamMemberHandler *corehandler.TeamMemberHandler,
	taskHandler *corehandler.TaskHandler,
	auditHandler *corehandler.AuditHandler,
	cacheHandler *corehandler.CacheHandler,
	overviewHandler *corehandler.OverviewHandler,
) []routes.Route {
	var allRoutes []routes.Route

	// Public 路由（健康检查）
	allRoutes = append(allRoutes, Public(healthHandler)...)

	// Admin 路由（系统管理）
	allRoutes = append(allRoutes, Admin(
		orgHandler,
		orgMemberHandler,
		teamHandler,
		teamMemberHandler,
		taskHandler,
		auditHandler,
		cacheHandler,
		overviewHandler,
	)...)

	return allRoutes
}
