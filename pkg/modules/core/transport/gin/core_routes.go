package gin

import (
	corehandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/core/transport/gin/handler"
	platformhttp "github.com/lwmacct/260101-go-pkg-ddd/pkg/platform/http"
)

// CoreDomainRoutes 返回 Core 域的所有路由（组织、团队、任务、审计等）
func CoreDomainRoutes(
	orgHandler *corehandler.OrgHandler,
	orgMemberHandler *corehandler.OrgMemberHandler,
	teamHandler *corehandler.TeamHandler,
	teamMemberHandler *corehandler.TeamMemberHandler,
	taskHandler *corehandler.TaskHandler,
	auditHandler *corehandler.AuditHandler,
	healthHandler *corehandler.HealthHandler,
	cacheHandler *corehandler.CacheHandler,
	overviewHandler *corehandler.OverviewHandler,
) []platformhttp.Route {
	var routes []platformhttp.Route
	routes = append(routes, PublicRoutes(healthHandler)...)
	routes = append(routes, AdminRoutes(
		orgHandler,
		orgMemberHandler,
		teamHandler,
		teamMemberHandler,
		taskHandler,
		auditHandler,
		cacheHandler,
		overviewHandler,
	)...)
	return routes
}
