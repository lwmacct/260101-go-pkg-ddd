package gin

import (
	corehandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/core/transport/gin/handler"
	platformhttp "github.com/lwmacct/260101-go-pkg-ddd/pkg/platform/http"
)

// AdminRoutes Core 域管理员路由（组织、团队、任务、审计等）
func AdminRoutes(
	orgHandler *corehandler.OrgHandler,
	orgMemberHandler *corehandler.OrgMemberHandler,
	teamHandler *corehandler.TeamHandler,
	teamMemberHandler *corehandler.TeamMemberHandler,
	taskHandler *corehandler.TaskHandler,
	auditHandler *corehandler.AuditHandler,
	cacheHandler *corehandler.CacheHandler,
	overviewHandler *corehandler.OverviewHandler,
) []platformhttp.Route {
	baseMiddlewares := []platformhttp.MiddlewareConfig{
		{Name: platformhttp.MiddlewareRequestID},
		{Name: platformhttp.MiddlewareOperationID},
		{Name: platformhttp.MiddlewareAuth},
		{Name: platformhttp.MiddlewareRBAC},
	}

	auditMiddlewares := append(baseMiddlewares, platformhttp.MiddlewareConfig{
		Name: platformhttp.MiddlewareAudit,
	})

	var routes []platformhttp.Route

	// Organization routes
	routes = append(routes, []platformhttp.Route{
		{
			Method:      platformhttp.GET,
			Path:        "/api/admin/orgs",
			Handler:     orgHandler.List,
			Op:          "admin:orgs:list",
			Middlewares: baseMiddlewares,
			Tags:        "Admin - Organizations",
			Summary:     "组织列表",
			Description: "获取组织列表",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/admin/orgs",
			Handler:     orgHandler.Create,
			Op:          "admin:orgs:create",
			Middlewares: auditMiddlewares,
			Tags:        "Admin - Organizations",
			Summary:     "创建组织",
			Description: "创建新组织",
		},
		{
			Method:      platformhttp.GET,
			Path:        "/api/admin/orgs/:id",
			Handler:     orgHandler.Get,
			Op:          "admin:orgs:get",
			Middlewares: baseMiddlewares,
			Tags:        "Admin - Organizations",
			Summary:     "组织详情",
			Description: "获取组织详细信息",
		},
		{
			Method:      platformhttp.PUT,
			Path:        "/api/admin/orgs/:id",
			Handler:     orgHandler.Update,
			Op:          "admin:orgs:update",
			Middlewares: auditMiddlewares,
			Tags:        "Admin - Organizations",
			Summary:     "更新组织",
			Description: "更新组织信息",
		},
		{
			Method:      platformhttp.DELETE,
			Path:        "/api/admin/orgs/:id",
			Handler:     orgHandler.Delete,
			Op:          "admin:orgs:delete",
			Middlewares: auditMiddlewares,
			Tags:        "Admin - Organizations",
			Summary:     "删除组织",
			Description: "删除组织",
		},
	}...)

	// Task routes
	routes = append(routes, []platformhttp.Route{
		{
			Method:      platformhttp.GET,
			Path:        "/api/admin/tasks",
			Handler:     taskHandler.List,
			Op:          "admin:tasks:list",
			Middlewares: baseMiddlewares,
			Tags:        "Admin - Tasks",
			Summary:     "任务列表",
			Description: "获取任务列表",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/admin/tasks",
			Handler:     taskHandler.Create,
			Op:          "admin:tasks:create",
			Middlewares: auditMiddlewares,
			Tags:        "Admin - Tasks",
			Summary:     "创建任务",
			Description: "创建新任务",
		},
		{
			Method:      platformhttp.GET,
			Path:        "/api/admin/tasks/:id",
			Handler:     taskHandler.Get,
			Op:          "admin:tasks:get",
			Middlewares: baseMiddlewares,
			Tags:        "Admin - Tasks",
			Summary:     "任务详情",
			Description: "获取任务详细信息",
		},
		{
			Method:      platformhttp.PUT,
			Path:        "/api/admin/tasks/:id",
			Handler:     taskHandler.Update,
			Op:          "admin:tasks:update",
			Middlewares: auditMiddlewares,
			Tags:        "Admin - Tasks",
			Summary:     "更新任务",
			Description: "更新任务信息",
		},
		{
			Method:      platformhttp.DELETE,
			Path:        "/api/admin/tasks/:id",
			Handler:     taskHandler.Delete,
			Op:          "admin:tasks:delete",
			Middlewares: auditMiddlewares,
			Tags:        "Admin - Tasks",
			Summary:     "删除任务",
			Description: "删除任务",
		},
	}...)

	return routes
}
