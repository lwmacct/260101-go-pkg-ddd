package http

import (
	"github.com/lwmacct/260101-go-pkg-ddd/ddd/core/adapters/http/routes"
)

// orgRoutes 返回 Org 域路由
// 包含：Org 域成员管理、Org 域团队管理
func (deps *RouterDependencies) orgRoutes() []routes.Route {
	// Org 域中间件模式
	// 注意：OrgContext 必须在 RBAC 之前，以便注入 org_id 供 RBAC 检查权限
	//
	// IMPORTANT: 必须使用 slices.Clone 或显式复制避免 Go slice 别名 bug。
	// 多次 append 到同一 slice 会导致后续 append 覆盖前面 slice 的元素。
	baseMiddlewares := []routes.MiddlewareConfig{
		{Name: routes.MiddlewareRequestID},
		{Name: routes.MiddlewareOperationID},
		{Name: routes.MiddlewareAuth},
	}

	// Org 上下文中间件链
	orgContextMiddlewares := append(cloneMiddlewares(baseMiddlewares), routes.MiddlewareConfig{Name: routes.MiddlewareOrgContext})
	orgRBACMiddlewares := append(cloneMiddlewares(orgContextMiddlewares), routes.MiddlewareConfig{Name: routes.MiddlewareRBAC})
	orgAuditMiddlewares := append(cloneMiddlewares(orgRBACMiddlewares), routes.MiddlewareConfig{Name: routes.MiddlewareAudit})

	// Team 上下文中间件链（必须从 orgRBACMiddlewares 的副本构建，避免覆盖 orgAuditMiddlewares）
	orgTeamGetMiddlewares := append(cloneMiddlewares(orgRBACMiddlewares), routes.MiddlewareConfig{
		Name:    routes.MiddlewareTeamContext,
		Options: map[string]any{"optional": true},
	})
	orgTeamContextMiddlewares := append(cloneMiddlewares(orgRBACMiddlewares), routes.MiddlewareConfig{Name: routes.MiddlewareTeamContext})
	orgTeamAuditMiddlewares := append(cloneMiddlewares(orgTeamContextMiddlewares), routes.MiddlewareConfig{Name: routes.MiddlewareAudit})

	// 只读路由使用 orgRBACMiddlewares（有 RBAC 但无 Audit）
	orgReadMiddlewares := orgRBACMiddlewares

	return []routes.Route{
		// ==================== Org 域 - 成员管理 ====================
		{
			Method:      routes.GET,
			Path:        "/api/org/:org_id/members",
			Handler:     deps.OrgMemberHandler.List,
			Op:          "org:members:list",
			Middlewares: orgReadMiddlewares,
			Tags:        "Org - Members",
			Summary:     "成员列表",
			Description: "获取组织成员列表",
		},
		{
			Method:      routes.POST,
			Path:        "/api/org/:org_id/members",
			Handler:     deps.OrgMemberHandler.Add,
			Op:          "org:members:add",
			Middlewares: orgAuditMiddlewares,
			Tags:        "Org - Members",
			Summary:     "添加成员",
			Description: "添加组织成员",
		},
		{
			Method:      routes.DELETE,
			Path:        "/api/org/:org_id/members/:user_id",
			Handler:     deps.OrgMemberHandler.Remove,
			Op:          "org:members:remove",
			Middlewares: orgAuditMiddlewares,
			Tags:        "Org - Members",
			Summary:     "移除成员",
			Description: "移除组织成员",
		},
		{
			Method:      routes.PUT,
			Path:        "/api/org/:org_id/members/:user_id/role",
			Handler:     deps.OrgMemberHandler.UpdateRole,
			Op:          "org:members:update:role",
			Middlewares: orgAuditMiddlewares,
			Tags:        "Org - Members",
			Summary:     "更新成员角色",
			Description: "更新组织成员角色",
		},

		// ==================== Org 域 - 团队管理 ====================
		{
			Method:      routes.POST,
			Path:        "/api/org/:org_id/teams",
			Handler:     deps.TeamHandler.Create,
			Op:          "org:teams:create",
			Middlewares: orgAuditMiddlewares,
			Tags:        "Org - Teams",
			Summary:     "创建团队",
			Description: "创建团队",
		},
		{
			Method:      routes.GET,
			Path:        "/api/org/:org_id/teams",
			Handler:     deps.TeamHandler.List,
			Op:          "org:teams:list",
			Middlewares: orgReadMiddlewares,
			Tags:        "Org - Teams",
			Summary:     "团队列表",
			Description: "获取团队列表",
		},
		{
			Method:      routes.GET,
			Path:        "/api/org/:org_id/teams/:team_id",
			Handler:     deps.TeamHandler.Get,
			Op:          "org:teams:get",
			Middlewares: orgTeamGetMiddlewares,
			Tags:        "Org - Teams",
			Summary:     "团队详情",
			Description: "获取团队详情",
		},
		{
			Method:      routes.PUT,
			Path:        "/api/org/:org_id/teams/:team_id",
			Handler:     deps.TeamHandler.Update,
			Op:          "org:teams:update",
			Middlewares: orgTeamAuditMiddlewares,
			Tags:        "Org - Teams",
			Summary:     "更新团队",
			Description: "更新团队信息",
		},
		{
			Method:      routes.DELETE,
			Path:        "/api/org/:org_id/teams/:team_id",
			Handler:     deps.TeamHandler.Delete,
			Op:          "org:teams:delete",
			Middlewares: orgTeamAuditMiddlewares,
			Tags:        "Org - Teams",
			Summary:     "删除团队",
			Description: "删除团队",
		},

		// ==================== Org 域 - 团队成员管理 ====================
		{
			Method:      routes.GET,
			Path:        "/api/org/:org_id/teams/:team_id/members",
			Handler:     deps.TeamMemberHandler.List,
			Op:          "org:team:members:list",
			Middlewares: orgTeamGetMiddlewares,
			Tags:        "Org - Teams",
			Summary:     "团队成员列表",
			Description: "获取团队成员列表",
		},
		{
			Method:      routes.POST,
			Path:        "/api/org/:org_id/teams/:team_id/members",
			Handler:     deps.TeamMemberHandler.Add,
			Op:          "org:team:members:add",
			Middlewares: orgTeamAuditMiddlewares,
			Tags:        "Org - Teams",
			Summary:     "添加团队成员",
			Description: "添加团队成员",
		},
		{
			Method:      routes.DELETE,
			Path:        "/api/org/:org_id/teams/:team_id/members/:user_id",
			Handler:     deps.TeamMemberHandler.Remove,
			Op:          "org:team:members:remove",
			Middlewares: orgTeamAuditMiddlewares,
			Tags:        "Org - Teams",
			Summary:     "移除团队成员",
			Description: "移除团队成员",
		},

		// ==================== Org 域 - 任务管理 ====================
		{
			Method:      routes.POST,
			Path:        "/api/org/:org_id/teams/:team_id/tasks",
			Handler:     deps.TaskHandler.Create,
			Op:          "org:tasks:create",
			Middlewares: orgTeamAuditMiddlewares,
			Tags:        "Org - Tasks",
			Summary:     "创建任务",
			Description: "创建任务",
		},
		{
			Method:      routes.GET,
			Path:        "/api/org/:org_id/teams/:team_id/tasks",
			Handler:     deps.TaskHandler.List,
			Op:          "org:tasks:list",
			Middlewares: orgTeamGetMiddlewares,
			Tags:        "Org - Tasks",
			Summary:     "任务列表",
			Description: "获取任务列表",
		},
		{
			Method:      routes.GET,
			Path:        "/api/org/:org_id/teams/:team_id/tasks/:id",
			Handler:     deps.TaskHandler.Get,
			Op:          "org:tasks:get",
			Middlewares: orgTeamGetMiddlewares,
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
			Description: "更新任务",
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

// cloneMiddlewares 创建中间件配置 slice 的副本。
// 避免 Go slice 别名 bug：多次 append 到同一 slice 会导致后续 append 覆盖前面 slice 的元素。
func cloneMiddlewares(src []routes.MiddlewareConfig) []routes.MiddlewareConfig {
	dst := make([]routes.MiddlewareConfig, len(src))
	copy(dst, src)
	return dst
}
