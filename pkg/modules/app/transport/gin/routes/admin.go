package routes

import (
	corehandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/app/transport/gin/handler"
	"github.com/lwmacct/260101-go-pkg-gin/pkg/routes"
)

// Admin App 域管理员路由（任务、缓存管理、系统概览）
func Admin(
	taskHandler *corehandler.TaskHandler,
	cacheHandler *corehandler.CacheHandler,
	overviewHandler *corehandler.OverviewHandler,
) []routes.Route {
	var allRoutes []routes.Route

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
