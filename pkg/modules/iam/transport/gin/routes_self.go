package gin

import (
	corehandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/core/transport/gin/handler"
	iamhandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/iam/transport/gin/handler"
	platformhttp "github.com/lwmacct/260101-go-pkg-ddd/pkg/platform/http"
)

// SelfRoutes 用户自服务路由（资料、令牌、配置）
func SelfRoutes(
	userProfileHandler *iamhandler.UserProfileHandler,
	patHandler *iamhandler.PATHandler,
	settingHandler *corehandler.SettingHandler,
	userSettingHandler *corehandler.UserSettingHandler,
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

	// User Profile routes
	routes = append(routes, []platformhttp.Route{
		{
			Method:      platformhttp.GET,
			Path:        "/api/user/profile",
			Handler:     userProfileHandler.GetProfile,
			Op:          "self:profile:get",
			Middlewares: baseMiddlewares,
			Tags:        "User - Profile",
			Summary:     "个人资料",
			Description: "获取当前用户资料",
		},
		{
			Method:      platformhttp.PUT,
			Path:        "/api/user/profile",
			Handler:     userProfileHandler.UpdateProfile,
			Op:          "self:profile:update",
			Middlewares: auditMiddlewares,
			Tags:        "User - Profile",
			Summary:     "更新资料",
			Description: "更新当前用户资料",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/user/profile/password",
			Handler:     userProfileHandler.ChangePassword,
			Op:          "self:password:change",
			Middlewares: auditMiddlewares,
			Tags:        "User - Profile",
			Summary:     "修改密码",
			Description: "修改当前用户密码",
		},
	}...)

	// PAT routes
	routes = append(routes, []platformhttp.Route{
		{
			Method:      platformhttp.GET,
			Path:        "/api/user/pats",
			Handler:     patHandler.ListTokens,
			Op:          "self:pats:list",
			Middlewares: baseMiddlewares,
			Tags:        "User - PAT",
			Summary:     "令牌列表",
			Description: "获取当前用户的 PAT 列表",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/user/pats",
			Handler:     patHandler.CreateToken,
			Op:          "self:pats:create",
			Middlewares: auditMiddlewares,
			Tags:        "User - PAT",
			Summary:     "创建令牌",
			Description: "创建个人访问令牌",
		},
		{
			Method:      platformhttp.DELETE,
			Path:        "/api/user/pats/:id",
			Handler:     patHandler.DeleteToken,
			Op:          "self:pats:delete",
			Middlewares: auditMiddlewares,
			Tags:        "User - PAT",
			Summary:     "删除令牌",
			Description: "删除个人访问令牌",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/user/pats/:id/enable",
			Handler:     patHandler.EnableToken,
			Op:          "self:pats:enable",
			Middlewares: auditMiddlewares,
			Tags:        "User - PAT",
			Summary:     "启用令牌",
			Description: "启用禁用的个人访问令牌",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/user/pats/:id/disable",
			Handler:     patHandler.DisableToken,
			Op:          "self:pats:disable",
			Middlewares: auditMiddlewares,
			Tags:        "User - PAT",
			Summary:     "禁用令牌",
			Description: "禁用个人访问令牌",
		},
	}...)

	// User Settings routes
	routes = append(routes, []platformhttp.Route{
		{
			Method:      platformhttp.GET,
			Path:        "/api/user/settings",
			Handler:     userSettingHandler.GetUserSettings,
			Op:          "self:settings:list",
			Middlewares: baseMiddlewares,
			Tags:        "User - Settings",
			Summary:     "配置列表",
			Description: "获取用户配置列表",
		},
		{
			Method:      platformhttp.GET,
			Path:        "/api/user/settings/categories",
			Handler:     userSettingHandler.ListUserSettingCategories,
			Op:          "self:settings:categories",
			Middlewares: baseMiddlewares,
			Tags:        "User - Settings",
			Summary:     "配置分类",
			Description: "获取配置分类",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/user/settings",
			Handler:     userSettingHandler.SetUserSetting,
			Op:          "self:settings:set",
			Middlewares: auditMiddlewares,
			Tags:        "User - Settings",
			Summary:     "设置配置",
			Description: "设置用户配置值",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/user/settings/batch",
			Handler:     userSettingHandler.BatchSetUserSettings,
			Op:          "self:settings:batch_set",
			Middlewares: auditMiddlewares,
			Tags:        "User - Settings",
			Summary:     "批量设置",
			Description: "批量设置用户配置",
		},
		{
			Method:      platformhttp.DELETE,
			Path:        "/api/user/settings",
			Handler:     userSettingHandler.ResetUserSetting,
			Op:          "self:settings:reset",
			Middlewares: auditMiddlewares,
			Tags:        "User - Settings",
			Summary:     "重置配置",
			Description: "重置用户配置为默认值",
		},
	}...)

	return routes
}
