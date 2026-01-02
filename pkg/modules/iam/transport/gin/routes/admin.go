package routes

import (
	"github.com/lwmacct/260101-go-pkg-gin/pkg/routes"

	corehandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/app/transport/gin/handler"
	iamhandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/iam/transport/gin/handler"
)

// Admin
func Admin(
	adminUserHandler *corehandler.AdminUserHandler,
	roleHandler *iamhandler.RoleHandler,
	settingHandler *corehandler.SettingHandler,
) []routes.Route {
	var allRoutes []routes.Route

	// User management routes
	allRoutes = append(allRoutes, []routes.Route{
		{
			Method:      routes.GET,
			Path:        "/api/admin/users",
			Handler:     adminUserHandler.ListUsers,
			Operation:   "admin:users:list",
			Tags:        "Admin - User Management",
			Summary:     "用户列表",
			Description: "分页获取用户列表",
		},
		{
			Method:      routes.GET,
			Path:        "/api/admin/users/:id",
			Handler:     adminUserHandler.GetUser,
			Operation:   "admin:users:get",
			Tags:        "Admin - User Management",
			Summary:     "用户详情",
			Description: "获取用户详细信息",
		},
		{
			Method:      routes.POST,
			Path:        "/api/admin/users",
			Handler:     adminUserHandler.CreateUser,
			Operation:   "admin:users:create",
			Tags:        "Admin - User Management",
			Summary:     "创建用户",
			Description: "创建新用户",
		},
		{
			Method:      routes.PUT,
			Path:        "/api/admin/users/:id",
			Handler:     adminUserHandler.UpdateUser,
			Operation:   "admin:users:update",
			Tags:        "Admin - User Management",
			Summary:     "更新用户",
			Description: "更新用户信息",
		},
		{
			Method:      routes.DELETE,
			Path:        "/api/admin/users/:id",
			Handler:     adminUserHandler.DeleteUser,
			Operation:   "admin:users:delete",
			Tags:        "Admin - User Management",
			Summary:     "删除用户",
			Description: "删除用户",
		},
		{
			Method:      routes.POST,
			Path:        "/api/admin/users/batch",
			Handler:     adminUserHandler.BatchCreateUsers,
			Operation:   "admin:users:batch_create",
			Tags:        "Admin - User Management",
			Summary:     "批量创建",
			Description: "批量创建用户",
		},
		{
			Method:      routes.PUT,
			Path:        "/api/admin/users/:id/roles",
			Handler:     adminUserHandler.AssignRoles,
			Operation:   "admin:users:assign_roles",
			Tags:        "Admin - User Management",
			Summary:     "分配角色",
			Description: "为用户分配角色",
		},
	}...)

	// Role management routes
	allRoutes = append(allRoutes, []routes.Route{
		{
			Method:      routes.GET,
			Path:        "/api/admin/roles",
			Handler:     roleHandler.ListRoles,
			Operation:   "admin:roles:list",
			Tags:        "Admin - Role Management",
			Summary:     "角色列表",
			Description: "分页获取角色列表",
		},
		{
			Method:      routes.GET,
			Path:        "/api/admin/roles/:id",
			Handler:     roleHandler.GetRole,
			Operation:   "admin:roles:get",
			Tags:        "Admin - Role Management",
			Summary:     "角色详情",
			Description: "获取角色详细信息",
		},
		{
			Method:      routes.POST,
			Path:        "/api/admin/roles",
			Handler:     roleHandler.CreateRole,
			Operation:   "admin:roles:create",
			Tags:        "Admin - Role Management",
			Summary:     "创建角色",
			Description: "创建新角色",
		},
		{
			Method:      routes.PUT,
			Path:        "/api/admin/roles/:id",
			Handler:     roleHandler.UpdateRole,
			Operation:   "admin:roles:update",
			Tags:        "Admin - Role Management",
			Summary:     "更新角色",
			Description: "更新角色信息",
		},
		{
			Method:      routes.DELETE,
			Path:        "/api/admin/roles/:id",
			Handler:     roleHandler.DeleteRole,
			Operation:   "admin:roles:delete",
			Tags:        "Admin - Role Management",
			Summary:     "删除角色",
			Description: "删除角色",
		},
		{
			Method:      routes.PUT,
			Path:        "/api/admin/roles/:id/permissions",
			Handler:     roleHandler.SetPermissions,
			Operation:   "admin:roles:set_permissions",
			Tags:        "Admin - Role Management",
			Summary:     "设置权限",
			Description: "为角色设置权限",
		},
	}...)

	// Setting management routes
	allRoutes = append(allRoutes, []routes.Route{
		{
			Method:      routes.GET,
			Path:        "/api/admin/settings",
			Handler:     settingHandler.GetSettings,
			Operation:   "admin:settings:list",
			Tags:        "Admin - Settings",
			Summary:     "设置列表",
			Description: "获取系统设置列表（Schema + 值）",
		},
		{
			Method:      routes.GET,
			Path:        "/api/admin/settings/:key",
			Handler:     settingHandler.GetSetting,
			Operation:   "admin:settings:get",
			Tags:        "Admin - Settings",
			Summary:     "设置详情",
			Description: "获取单个设置详情",
		},
		{
			Method:      routes.POST,
			Path:        "/api/admin/settings",
			Handler:     settingHandler.CreateSetting,
			Operation:   "admin:settings:create",
			Tags:        "Admin - Settings",
			Summary:     "创建设置",
			Description: "创建新设置",
		},
		{
			Method:      routes.PUT,
			Path:        "/api/admin/settings/:key",
			Handler:     settingHandler.UpdateSetting,
			Operation:   "admin:settings:update",
			Tags:        "Admin - Settings",
			Summary:     "更新设置",
			Description: "更新设置值",
		},
		{
			Method:      routes.DELETE,
			Path:        "/api/admin/settings/:key",
			Handler:     settingHandler.DeleteSetting,
			Operation:   "admin:settings:delete",
			Tags:        "Admin - Settings",
			Summary:     "删除设置",
			Description: "删除设置",
		},
		{
			Method:      routes.POST,
			Path:        "/api/admin/settings/batch",
			Handler:     settingHandler.BatchUpdateSettings,
			Operation:   "admin:settings:batch_update",
			Tags:        "Admin - Settings",
			Summary:     "批量更新设置",
			Description: "批量更新多个设置值",
		},
		{
			Method:      routes.GET,
			Path:        "/api/admin/settings/categories",
			Handler:     settingHandler.GetCategories,
			Operation:   "admin:settings:categories:list",
			Tags:        "Admin - Settings",
			Summary:     "设置分类列表",
			Description: "获取设置分类列表",
		},
		{
			Method:      routes.GET,
			Path:        "/api/admin/settings/categories/:id",
			Handler:     settingHandler.GetCategory,
			Operation:   "admin:settings:categories:get",
			Tags:        "Admin - Settings",
			Summary:     "设置分类详情",
			Description: "获取设置分类详情",
		},
		{
			Method:      routes.POST,
			Path:        "/api/admin/settings/categories",
			Handler:     settingHandler.CreateCategory,
			Operation:   "admin:settings:categories:create",
			Tags:        "Admin - Settings",
			Summary:     "创建设置分类",
			Description: "创建设置分类",
		},
		{
			Method:      routes.PUT,
			Path:        "/api/admin/settings/categories/:id",
			Handler:     settingHandler.UpdateCategory,
			Operation:   "admin:settings:categories:update",
			Tags:        "Admin - Settings",
			Summary:     "更新设置分类",
			Description: "更新设置分类",
		},
		{
			Method:      routes.DELETE,
			Path:        "/api/admin/settings/categories/:id",
			Handler:     settingHandler.DeleteCategory,
			Operation:   "admin:settings:categories:delete",
			Tags:        "Admin - Settings",
			Summary:     "删除设置分类",
			Description: "删除设置分类",
		},
	}...)

	return allRoutes
}

// UserOrg 用户组织视图路由
func UserOrg(userOrgHandler *iamhandler.UserOrgHandler) []routes.Route {
	return []routes.Route{
		{
			Method:      routes.GET,
			Path:        "/api/user/orgs",
			Handler:     userOrgHandler.ListMyOrganizations,
			Operation:   "self:orgs:list",
			Tags:        "User - Organizations",
			Summary:     "我的组织",
			Description: "获取用户所属组织列表",
		},
		{
			Method:      routes.GET,
			Path:        "/api/user/teams",
			Handler:     userOrgHandler.ListMyTeams,
			Operation:   "self:teams:list",
			Tags:        "User - Organizations",
			Summary:     "我的团队",
			Description: "获取用户所属团队列表",
		},
	}
}
