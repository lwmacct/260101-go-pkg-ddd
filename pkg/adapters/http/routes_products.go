package http

import (
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/adapters/http/routes"
)

// productRoutes 返回 Admin 域 - 产品管理路由 + Org 域 - 团队任务管理路由
func (deps *RouterDependencies) productRoutes() []routes.Route {
	// Admin 域中间件模式
	baseMiddlewares := []routes.MiddlewareConfig{
		{Name: routes.MiddlewareRequestID},
		{Name: routes.MiddlewareOperationID},
		{Name: routes.MiddlewareAuth},
		{Name: routes.MiddlewareRBAC},
	}

	auditMiddlewares := append(cloneMiddlewares(baseMiddlewares), routes.MiddlewareConfig{Name: routes.MiddlewareAudit})

	// Org 域中间件模式（用于团队任务管理）
	// 注意：OrgContext 必须在 RBAC 之前，以便注入 org_id 供 RBAC 检查权限
	orgBaseMiddlewares := []routes.MiddlewareConfig{
		{Name: routes.MiddlewareRequestID},
		{Name: routes.MiddlewareOperationID},
		{Name: routes.MiddlewareAuth},
	}

	orgContextMiddlewares := append(cloneMiddlewares(orgBaseMiddlewares), routes.MiddlewareConfig{Name: routes.MiddlewareOrgContext})
	orgRBACMiddlewares := append(cloneMiddlewares(orgContextMiddlewares), routes.MiddlewareConfig{Name: routes.MiddlewareRBAC})
	orgTeamContextMiddlewares := append(cloneMiddlewares(orgRBACMiddlewares),
		routes.MiddlewareConfig{Name: routes.MiddlewareTeamContext, Options: map[string]any{"optional": true}})
	orgTeamAuditMiddlewares := append(cloneMiddlewares(orgTeamContextMiddlewares), routes.MiddlewareConfig{Name: routes.MiddlewareAudit})

	return []routes.Route{
		// ==================== Admin 域 - 产品管理 ====================
		{
			Method:      routes.POST,
			Path:        "/api/admin/products",
			Handler:     deps.ProductHandler.Create,
			Op:          "admin:products:create",
			Middlewares: auditMiddlewares,
			Tags:        "Admin - Products",
			Summary:     "创建产品",
			Description: "创建产品",
		},
		{
			Method:      routes.GET,
			Path:        "/api/admin/products",
			Handler:     deps.ProductHandler.List,
			Op:          "admin:products:list",
			Middlewares: baseMiddlewares,
			Tags:        "Admin - Products",
			Summary:     "产品列表",
			Description: "获取产品列表",
		},
		{
			Method:      routes.GET,
			Path:        "/api/admin/products/:id",
			Handler:     deps.ProductHandler.Get,
			Op:          "admin:products:get",
			Middlewares: baseMiddlewares,
			Tags:        "Admin - Products",
			Summary:     "产品详情",
			Description: "获取产品详情",
		},
		{
			Method:      routes.PUT,
			Path:        "/api/admin/products/:id",
			Handler:     deps.ProductHandler.Update,
			Op:          "admin:products:update",
			Middlewares: auditMiddlewares,
			Tags:        "Admin - Products",
			Summary:     "更新产品",
			Description: "更新产品信息",
		},
		{
			Method:      routes.DELETE,
			Path:        "/api/admin/products/:id",
			Handler:     deps.ProductHandler.Delete,
			Op:          "admin:products:delete",
			Middlewares: auditMiddlewares,
			Tags:        "Admin - Products",
			Summary:     "删除产品",
			Description: "删除产品",
		},

		// ==================== Org 域 - 团队任务管理 ====================
		{
			Method:      routes.POST,
			Path:        "/api/org/:org_id/teams/:team_id/tasks",
			Handler:     deps.TaskHandler.Create,
			Op:          "org:tasks:create",
			Middlewares: orgTeamAuditMiddlewares,
			Tags:        "Org - Tasks",
			Summary:     "创建任务",
			Description: "创建团队任务",
		},
		{
			Method:      routes.GET,
			Path:        "/api/org/:org_id/teams/:team_id/tasks",
			Handler:     deps.TaskHandler.List,
			Op:          "org:tasks:list",
			Middlewares: orgTeamContextMiddlewares,
			Tags:        "Org - Tasks",
			Summary:     "任务列表",
			Description: "获取团队任务列表",
		},
		{
			Method:      routes.GET,
			Path:        "/api/org/:org_id/teams/:team_id/tasks/:id",
			Handler:     deps.TaskHandler.Get,
			Op:          "org:tasks:get",
			Middlewares: orgTeamContextMiddlewares,
			Tags:        "Org - Tasks",
			Summary:     "任务详情",
			Description: "获取任务详情",
		},
		{
			Method:      routes.PUT,
			Path:        "/api/org/:org_id/teams/:team_id/tasks/:id",
			Handler:     deps.TaskHandler.Update,
			Op:          "org:tasks:update",
			Middlewares: orgTeamAuditMiddlewares,
			Tags:        "Org - Tasks",
			Summary:     "更新任务",
			Description: "更新任务信息",
		},
		{
			Method:      routes.DELETE,
			Path:        "/api/org/:org_id/teams/:team_id/tasks/:id",
			Handler:     deps.TaskHandler.Delete,
			Op:          "org:tasks:delete",
			Middlewares: orgTeamAuditMiddlewares,
			Tags:        "Org - Tasks",
			Summary:     "删除任务",
			Description: "删除任务",
		},
	}
}
