package routes

import (
	corehandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/app/transport/gin/handler"
	"github.com/lwmacct/260101-go-pkg-gin/pkg/routes"
)

// Admin Core 域管理员路由（组织、团队、任务、审计等）
func Admin(
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

	// Organization routes
	allRoutes = append(allRoutes, []routes.Route{
		{
			Method:      routes.GET,
			Path:        "/api/admin/orgs",
			Handler:     orgHandler.List,
			Operation:   "admin:orgs:list",
			Tags:        "Admin - Organizations",
			Summary:     "组织列表",
			Description: "获取组织列表",
		},
		{
			Method:      routes.POST,
			Path:        "/api/admin/orgs",
			Handler:     orgHandler.Create,
			Operation:   "admin:orgs:create",
			Tags:        "Admin - Organizations",
			Summary:     "创建组织",
			Description: "创建新组织",
		},
		{
			Method:      routes.GET,
			Path:        "/api/admin/orgs/:id",
			Handler:     orgHandler.Get,
			Operation:   "admin:orgs:get",
			Tags:        "Admin - Organizations",
			Summary:     "组织详情",
			Description: "获取组织详细信息",
		},
		{
			Method:      routes.PUT,
			Path:        "/api/admin/orgs/:id",
			Handler:     orgHandler.Update,
			Operation:   "admin:orgs:update",
			Tags:        "Admin - Organizations",
			Summary:     "更新组织",
			Description: "更新组织信息",
		},
		{
			Method:      routes.DELETE,
			Path:        "/api/admin/orgs/:id",
			Handler:     orgHandler.Delete,
			Operation:   "admin:orgs:delete",
			Tags:        "Admin - Organizations",
			Summary:     "删除组织",
			Description: "删除组织",
		},
	}...)

	// Task routes
	allRoutes = append(allRoutes, []routes.Route{
		{
			Method:      routes.GET,
			Path:        "/api/admin/tasks",
			Handler:     taskHandler.List,
			Operation:   "admin:tasks:list",
			Tags:        "Admin - Tasks",
			Summary:     "任务列表",
			Description: "获取任务列表",
		},
		{
			Method:      routes.POST,
			Path:        "/api/admin/tasks",
			Handler:     taskHandler.Create,
			Operation:   "admin:tasks:create",
			Tags:        "Admin - Tasks",
			Summary:     "创建任务",
			Description: "创建新任务",
		},
		{
			Method:      routes.GET,
			Path:        "/api/admin/tasks/:id",
			Handler:     taskHandler.Get,
			Operation:   "admin:tasks:get",
			Tags:        "Admin - Tasks",
			Summary:     "任务详情",
			Description: "获取任务详细信息",
		},
		{
			Method:      routes.PUT,
			Path:        "/api/admin/tasks/:id",
			Handler:     taskHandler.Update,
			Operation:   "admin:tasks:update",
			Tags:        "Admin - Tasks",
			Summary:     "更新任务",
			Description: "更新任务信息",
		},
		{
			Method:      routes.DELETE,
			Path:        "/api/admin/tasks/:id",
			Handler:     taskHandler.Delete,
			Operation:   "admin:tasks:delete",
			Tags:        "Admin - Tasks",
			Summary:     "删除任务",
			Description: "删除任务",
		},
	}...)

	// Audit log routes
	allRoutes = append(allRoutes, []routes.Route{
		{
			Method:      routes.GET,
			Path:        "/api/admin/audit",
			Handler:     auditHandler.ListLogs,
			Operation:   "admin:audit:list",
			Tags:        "Admin - Audit",
			Summary:     "审计日志列表",
			Description: "获取审计日志列表",
		},
		{
			Method:      routes.GET,
			Path:        "/api/admin/audit/:id",
			Handler:     auditHandler.GetLog,
			Operation:   "admin:audit:get",
			Tags:        "Admin - Audit",
			Summary:     "审计日志详情",
			Description: "获取审计日志详情",
		},
		{
			Method:      routes.GET,
			Path:        "/api/admin/audit/actions",
			Handler:     auditHandler.GetActions,
			Operation:   "admin:audit:actions",
			Tags:        "Admin - Audit",
			Summary:     "审计操作列表",
			Description: "获取可审计的操作列表",
		},
	}...)

	// Cache management routes
	allRoutes = append(allRoutes, []routes.Route{
		{
			Method:      routes.GET,
			Path:        "/api/admin/cache/info",
			Handler:     cacheHandler.Info,
			Operation:   "admin:cache:info",
			Tags:        "Admin - Cache",
			Summary:     "缓存信息",
			Description: "获取 Redis 缓存信息",
		},
		{
			Method:      routes.GET,
			Path:        "/api/admin/cache/keys",
			Handler:     cacheHandler.ScanKeys,
			Operation:   "admin:cache:keys",
			Tags:        "Admin - Cache",
			Summary:     "扫描缓存键",
			Description: "扫描 Redis 缓存键",
		},
		{
			Method:      routes.GET,
			Path:        "/api/admin/cache/key",
			Handler:     cacheHandler.GetKey,
			Operation:   "admin:cache:key:get",
			Tags:        "Admin - Cache",
			Summary:     "获取缓存键值",
			Description: "获取指定 Redis 键的值",
		},
		{
			Method:      routes.DELETE,
			Path:        "/api/admin/cache/key",
			Handler:     cacheHandler.DeleteKey,
			Operation:   "admin:cache:key:delete",
			Tags:        "Admin - Cache",
			Summary:     "删除缓存键",
			Description: "删除指定 Redis 键",
		},
		{
			Method:      routes.DELETE,
			Path:        "/api/admin/cache/keys",
			Handler:     cacheHandler.DeleteByPattern,
			Operation:   "admin:cache:keys:delete",
			Tags:        "Admin - Cache",
			Summary:     "批量删除缓存键",
			Description: "按模式批量删除 Redis 键",
		},
	}...)

	// Overview routes
	allRoutes = append(allRoutes, []routes.Route{
		{
			Method:      routes.GET,
			Path:        "/api/admin/overview/stats",
			Handler:     overviewHandler.GetStats,
			Operation:   "admin:overview:stats",
			Tags:        "Admin - Overview",
			Summary:     "系统统计",
			Description: "获取系统统计数据",
		},
	}...)

	return allRoutes
}
