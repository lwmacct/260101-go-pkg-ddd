package gin

import (
	corehandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/core/transport/gin/handler"
	iamhandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/iam/transport/gin/handler"
	platformhttp "github.com/lwmacct/260101-go-pkg-ddd/pkg/platform/http"
)

// AdminRoutes Admin 路由（用户、角色管理）
func AdminRoutes(
	adminUserHandler *corehandler.AdminUserHandler,
	roleHandler *iamhandler.RoleHandler,
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

	// User management routes
	routes = append(routes, []platformhttp.Route{
		{
			Method:      platformhttp.GET,
			Path:        "/api/admin/users",
			Handler:     adminUserHandler.ListUsers,
			Op:          "admin:users:list",
			Middlewares: baseMiddlewares,
			Tags:        "Admin - User Management",
			Summary:     "用户列表",
			Description: "分页获取用户列表",
		},
		{
			Method:      platformhttp.GET,
			Path:        "/api/admin/users/:id",
			Handler:     adminUserHandler.GetUser,
			Op:          "admin:users:get",
			Middlewares: baseMiddlewares,
			Tags:        "Admin - User Management",
			Summary:     "用户详情",
			Description: "获取用户详细信息",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/admin/users",
			Handler:     adminUserHandler.CreateUser,
			Op:          "admin:users:create",
			Middlewares: auditMiddlewares,
			Tags:        "Admin - User Management",
			Summary:     "创建用户",
			Description: "创建新用户",
		},
		{
			Method:      platformhttp.PUT,
			Path:        "/api/admin/users/:id",
			Handler:     adminUserHandler.UpdateUser,
			Op:          "admin:users:update",
			Middlewares: auditMiddlewares,
			Tags:        "Admin - User Management",
			Summary:     "更新用户",
			Description: "更新用户信息",
		},
		{
			Method:      platformhttp.DELETE,
			Path:        "/api/admin/users/:id",
			Handler:     adminUserHandler.DeleteUser,
			Op:          "admin:users:delete",
			Middlewares: auditMiddlewares,
			Tags:        "Admin - User Management",
			Summary:     "删除用户",
			Description: "删除用户",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/admin/users/batch",
			Handler:     adminUserHandler.BatchCreateUsers,
			Op:          "admin:users:batch_create",
			Middlewares: auditMiddlewares,
			Tags:        "Admin - User Management",
			Summary:     "批量创建",
			Description: "批量创建用户",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/admin/users/:id/roles",
			Handler:     adminUserHandler.AssignRoles,
			Op:          "admin:users:assign_roles",
			Middlewares: auditMiddlewares,
			Tags:        "Admin - User Management",
			Summary:     "分配角色",
			Description: "为用户分配角色",
		},
	}...)

	// Role management routes
	routes = append(routes, []platformhttp.Route{
		{
			Method:      platformhttp.GET,
			Path:        "/api/admin/roles",
			Handler:     roleHandler.ListRoles,
			Op:          "admin:roles:list",
			Middlewares: baseMiddlewares,
			Tags:        "Admin - Role Management",
			Summary:     "角色列表",
			Description: "分页获取角色列表",
		},
		{
			Method:      platformhttp.GET,
			Path:        "/api/admin/roles/:id",
			Handler:     roleHandler.GetRole,
			Op:          "admin:roles:get",
			Middlewares: baseMiddlewares,
			Tags:        "Admin - Role Management",
			Summary:     "角色详情",
			Description: "获取角色详细信息",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/admin/roles",
			Handler:     roleHandler.CreateRole,
			Op:          "admin:roles:create",
			Middlewares: auditMiddlewares,
			Tags:        "Admin - Role Management",
			Summary:     "创建角色",
			Description: "创建新角色",
		},
		{
			Method:      platformhttp.PUT,
			Path:        "/api/admin/roles/:id",
			Handler:     roleHandler.UpdateRole,
			Op:          "admin:roles:update",
			Middlewares: auditMiddlewares,
			Tags:        "Admin - Role Management",
			Summary:     "更新角色",
			Description: "更新角色信息",
		},
		{
			Method:      platformhttp.DELETE,
			Path:        "/api/admin/roles/:id",
			Handler:     roleHandler.DeleteRole,
			Op:          "admin:roles:delete",
			Middlewares: auditMiddlewares,
			Tags:        "Admin - Role Management",
			Summary:     "删除角色",
			Description: "删除角色",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/admin/roles/:id/permissions",
			Handler:     roleHandler.SetPermissions,
			Op:          "admin:roles:set_permissions",
			Middlewares: auditMiddlewares,
			Tags:        "Admin - Role Management",
			Summary:     "设置权限",
			Description: "为角色设置权限",
		},
	}...)

	return routes
}

// UserOrgRoutes 用户组织视图路由
func UserOrgRoutes(userOrgHandler *iamhandler.UserOrgHandler) []platformhttp.Route {
	baseMiddlewares := []platformhttp.MiddlewareConfig{
		{Name: platformhttp.MiddlewareRequestID},
		{Name: platformhttp.MiddlewareOperationID},
		{Name: platformhttp.MiddlewareAuth},
		{Name: platformhttp.MiddlewareRBAC},
	}

	return []platformhttp.Route{
		{
			Method:      platformhttp.GET,
			Path:        "/api/user/orgs",
			Handler:     userOrgHandler.ListMyOrganizations,
			Op:          "self:orgs:list",
			Middlewares: baseMiddlewares,
			Tags:        "User - Organizations",
			Summary:     "我的组织",
			Description: "获取用户所属组织列表",
		},
		{
			Method:      platformhttp.GET,
			Path:        "/api/user/teams",
			Handler:     userOrgHandler.ListMyTeams,
			Op:          "self:teams:list",
			Middlewares: baseMiddlewares,
			Tags:        "User - Organizations",
			Summary:     "我的团队",
			Description: "获取用户所属团队列表",
		},
	}
}
