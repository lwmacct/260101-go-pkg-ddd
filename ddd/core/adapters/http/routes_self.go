package http

import (
	"github.com/lwmacct/260101-go-pkg-ddd/ddd/core/adapters/http/routes"
)

// selfRoutes 返回 Self 域路由（个人资料、令牌、配置）
func (deps *RouterDependencies) selfRoutes() []routes.Route {
	// Self 域中间件模式
	baseMiddlewares := []routes.MiddlewareConfig{
		{Name: routes.MiddlewareRequestID},
		{Name: routes.MiddlewareOperationID},
		{Name: routes.MiddlewareAuth},
		{Name: routes.MiddlewareRBAC},
	}

	auditMiddlewares := append(cloneMiddlewares(baseMiddlewares), routes.MiddlewareConfig{Name: routes.MiddlewareAudit})

	return []routes.Route{
		// ==================== 个人资料 ====================
		{
			Method:      routes.GET,
			Path:        "/api/user/profile",
			Handler:     deps.UserProfileHandler.GetProfile,
			Op:          "self:profile:get",
			Middlewares: baseMiddlewares,
			Tags:        "User - Profile",
			Summary:     "获取个人资料",
			Description: "获取用户个人资料",
		},
		{
			Method:      routes.PUT,
			Path:        "/api/user/profile",
			Handler:     deps.UserProfileHandler.UpdateProfile,
			Op:          "self:profile:update",
			Middlewares: auditMiddlewares,
			Tags:        "User - Profile",
			Summary:     "更新个人资料",
			Description: "更新用户个人资料",
		},
		{
			Method:      routes.PUT,
			Path:        "/api/user/password",
			Handler:     deps.UserProfileHandler.ChangePassword,
			Op:          "self:password:update",
			Middlewares: auditMiddlewares,
			Tags:        "User - Profile",
			Summary:     "修改密码",
			Description: "修改用户密码",
		},
		{
			Method:      routes.DELETE,
			Path:        "/api/user/account",
			Handler:     deps.UserProfileHandler.DeleteAccount,
			Op:          "self:account:delete",
			Middlewares: auditMiddlewares,
			Tags:        "User - Profile",
			Summary:     "删除账户",
			Description: "删除用户账户",
		},

		// ==================== 访问令牌 ====================
		// 注意：scopes 路由必须在 :id 路由之前
		{
			Method:      routes.POST,
			Path:        "/api/user/tokens",
			Handler:     deps.PATHandler.CreateToken,
			Op:          "self:tokens:create",
			Middlewares: auditMiddlewares,
			Tags:        "User - Tokens",
			Summary:     "创建令牌",
			Description: "创建个人访问令牌",
		},
		{
			Method:      routes.GET,
			Path:        "/api/user/tokens",
			Handler:     deps.PATHandler.ListTokens,
			Op:          "self:tokens:list",
			Middlewares: baseMiddlewares,
			Tags:        "User - Tokens",
			Summary:     "令牌列表",
			Description: "获取个人访问令牌列表",
		},
		{
			Method:      routes.GET,
			Path:        "/api/user/tokens/scopes",
			Handler:     deps.PATHandler.ListScopes,
			Op:          "self:tokens:scopes",
			Middlewares: baseMiddlewares,
			Tags:        "User - Tokens",
			Summary:     "令牌作用域",
			Description: "获取令牌作用域列表",
		},
		{
			Method:      routes.GET,
			Path:        "/api/user/tokens/:id",
			Handler:     deps.PATHandler.GetToken,
			Op:          "self:tokens:get",
			Middlewares: baseMiddlewares,
			Tags:        "User - Tokens",
			Summary:     "令牌详情",
			Description: "获取令牌详情",
		},
		{
			Method:      routes.DELETE,
			Path:        "/api/user/tokens/:id",
			Handler:     deps.PATHandler.DeleteToken,
			Op:          "self:tokens:delete",
			Middlewares: auditMiddlewares,
			Tags:        "User - Tokens",
			Summary:     "删除令牌",
			Description: "删除个人访问令牌",
		},
		{
			Method:      routes.PATCH,
			Path:        "/api/user/tokens/:id/disable",
			Handler:     deps.PATHandler.DisableToken,
			Op:          "self:tokens:disable",
			Middlewares: auditMiddlewares,
			Tags:        "User - Tokens",
			Summary:     "禁用令牌",
			Description: "禁用个人访问令牌",
		},
		{
			Method:      routes.PATCH,
			Path:        "/api/user/tokens/:id/enable",
			Handler:     deps.PATHandler.EnableToken,
			Op:          "self:tokens:enable",
			Middlewares: auditMiddlewares,
			Tags:        "User - Tokens",
			Summary:     "启用令牌",
			Description: "启用个人访问令牌",
		},

		// ==================== 用户配置 ====================
		// 注意：categories 和 batch 路由必须在 :key 路由之前
		{
			Method:      routes.GET,
			Path:        "/api/user/settings/categories",
			Handler:     deps.UserSettingHandler.ListUserSettingCategories,
			Op:          "self:settings:categories:list",
			Middlewares: baseMiddlewares,
			Tags:        "User - Settings",
			Summary:     "配置分类列表",
			Description: "获取用户配置分类列表",
		},
		{
			Method:      routes.POST,
			Path:        "/api/user/settings/batch",
			Handler:     deps.UserSettingHandler.BatchSetUserSettings,
			Op:          "self:settings:batch-set",
			Middlewares: auditMiddlewares,
			Tags:        "User - Settings",
			Summary:     "批量设置配置",
			Description: "批量设置用户配置",
		},
		{
			Method:      routes.GET,
			Path:        "/api/user/settings",
			Handler:     deps.UserSettingHandler.GetUserSettings,
			Op:          "self:settings:list",
			Middlewares: baseMiddlewares,
			Tags:        "User - Settings",
			Summary:     "配置列表",
			Description: "获取用户配置列表",
		},
		{
			Method:      routes.GET,
			Path:        "/api/user/settings/:key",
			Handler:     deps.UserSettingHandler.GetUserSetting,
			Op:          "self:settings:get",
			Middlewares: baseMiddlewares,
			Tags:        "User - Settings",
			Summary:     "配置详情",
			Description: "获取用户配置详情",
		},
		{
			Method:      routes.PUT,
			Path:        "/api/user/settings/:key",
			Handler:     deps.UserSettingHandler.SetUserSetting,
			Op:          "self:settings:set",
			Middlewares: auditMiddlewares,
			Tags:        "User - Settings",
			Summary:     "设置配置",
			Description: "设置用户配置",
		},
		{
			Method:      routes.DELETE,
			Path:        "/api/user/settings/:key",
			Handler:     deps.UserSettingHandler.ResetUserSetting,
			Op:          "self:settings:reset",
			Middlewares: auditMiddlewares,
			Tags:        "User - Settings",
			Summary:     "重置配置",
			Description: "重置用户配置",
		},

		// ==================== 用户组织/团队 ====================
		{
			Method:      routes.GET,
			Path:        "/api/user/orgs",
			Handler:     deps.UserOrgHandler.ListMyOrganizations,
			Op:          "self:orgs:list",
			Middlewares: baseMiddlewares,
			Tags:        "User - Organizations",
			Summary:     "我的组织",
			Description: "获取用户所属组织列表",
		},
		{
			Method:      routes.GET,
			Path:        "/api/user/teams",
			Handler:     deps.UserOrgHandler.ListMyTeams,
			Op:          "self:teams:list",
			Middlewares: baseMiddlewares,
			Tags:        "User - Organizations",
			Summary:     "我的团队",
			Description: "获取用户所属团队列表",
		},
	}
}
