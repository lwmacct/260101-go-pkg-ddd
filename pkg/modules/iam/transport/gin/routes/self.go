package routes

import (
	"github.com/lwmacct/260101-go-pkg-gin/pkg/routes"

	corehandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/app/transport/gin/handler"
	iamhandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/iam/transport/gin/handler"
)

// Self 用户自服务路由（资料、令牌、配置）
func Self(
	userProfileHandler *iamhandler.UserProfileHandler,
	patHandler *iamhandler.PATHandler,
	settingHandler *corehandler.SettingHandler,
	userSettingHandler *corehandler.UserSettingHandler,
) []routes.Route {
	var allRoutes []routes.Route

	// User Profile routes
	allRoutes = append(allRoutes, []routes.Route{
		{
			Method:      routes.GET,
			Path:        "/api/user/profile",
			Handler:     userProfileHandler.GetProfile,
			Operation:   "self:profile:get",
			Tags:        "User - Profile",
			Summary:     "个人资料",
			Description: "获取当前用户资料",
		},
		{
			Method:      routes.PUT,
			Path:        "/api/user/profile",
			Handler:     userProfileHandler.UpdateProfile,
			Operation:   "self:profile:update",
			Tags:        "User - Profile",
			Summary:     "更新资料",
			Description: "更新当前用户资料",
		},
		{
			Method:      routes.PUT,
			Path:        "/api/user/password",
			Handler:     userProfileHandler.ChangePassword,
			Operation:   "self:password:change",
			Tags:        "User - Profile",
			Summary:     "修改密码",
			Description: "修改当前用户密码",
		},
		{
			Method:      routes.DELETE,
			Path:        "/api/user/account",
			Handler:     userProfileHandler.DeleteAccount,
			Operation:   "self:account:delete",
			Tags:        "User - Profile",
			Summary:     "删除账户",
			Description: "删除当前用户账户",
		},
	}...)

	// PAT routes
	allRoutes = append(allRoutes, []routes.Route{
		{
			Method:      routes.GET,
			Path:        "/api/user/tokens",
			Handler:     patHandler.ListTokens,
			Operation:   "self:tokens:list",
			Tags:        "User - PAT",
			Summary:     "令牌列表",
			Description: "获取当前用户的 PAT 列表",
		},
		{
			Method:      routes.POST,
			Path:        "/api/user/tokens",
			Handler:     patHandler.CreateToken,
			Operation:   "self:tokens:create",
			Tags:        "User - PAT",
			Summary:     "创建令牌",
			Description: "创建个人访问令牌",
		},
		{
			Method:      routes.GET,
			Path:        "/api/user/tokens/:id",
			Handler:     patHandler.GetToken,
			Operation:   "self:tokens:get",
			Tags:        "User - PAT",
			Summary:     "令牌详情",
			Description: "获取个人访问令牌详情",
		},
		{
			Method:      routes.GET,
			Path:        "/api/user/tokens/scopes",
			Handler:     patHandler.ListScopes,
			Operation:   "self:tokens:scopes",
			Tags:        "User - PAT",
			Summary:     "令牌作用域",
			Description: "获取令牌作用域列表",
		},
		{
			Method:      routes.DELETE,
			Path:        "/api/user/tokens/:id",
			Handler:     patHandler.DeleteToken,
			Operation:   "self:tokens:delete",
			Tags:        "User - PAT",
			Summary:     "删除令牌",
			Description: "删除个人访问令牌",
		},
		{
			Method:      routes.PATCH,
			Path:        "/api/user/tokens/:id/disable",
			Handler:     patHandler.DisableToken,
			Operation:   "self:tokens:disable",
			Tags:        "User - PAT",
			Summary:     "禁用令牌",
			Description: "禁用个人访问令牌",
		},
		{
			Method:      routes.PATCH,
			Path:        "/api/user/tokens/:id/enable",
			Handler:     patHandler.EnableToken,
			Operation:   "self:tokens:enable",
			Tags:        "User - PAT",
			Summary:     "启用令牌",
			Description: "启用禁用的个人访问令牌",
		},
	}...)

	// User Settings routes
	allRoutes = append(allRoutes, []routes.Route{
		{
			Method:      routes.GET,
			Path:        "/api/user/settings",
			Handler:     userSettingHandler.GetUserSettings,
			Operation:   "self:settings:list",
			Tags:        "User - Settings",
			Summary:     "配置列表",
			Description: "获取用户配置列表",
		},
		{
			Method:      routes.GET,
			Path:        "/api/user/settings/:key",
			Handler:     userSettingHandler.GetUserSetting,
			Operation:   "self:settings:get",
			Tags:        "User - Settings",
			Summary:     "配置详情",
			Description: "获取单个用户配置详情",
		},
		{
			Method:      routes.GET,
			Path:        "/api/user/settings/categories",
			Handler:     userSettingHandler.ListUserSettingCategories,
			Operation:   "self:settings:categories",
			Tags:        "User - Settings",
			Summary:     "配置分类",
			Description: "获取配置分类",
		},
		{
			Method:      routes.PUT,
			Path:        "/api/user/settings/:key",
			Handler:     userSettingHandler.SetUserSetting,
			Operation:   "self:settings:set",
			Tags:        "User - Settings",
			Summary:     "设置配置",
			Description: "设置用户配置值",
		},
		{
			Method:      routes.POST,
			Path:        "/api/user/settings/batch",
			Handler:     userSettingHandler.BatchSetUserSettings,
			Operation:   "self:settings:batch_set",
			Tags:        "User - Settings",
			Summary:     "批量设置",
			Description: "批量设置用户配置",
		},
		{
			Method:      routes.DELETE,
			Path:        "/api/user/settings/:key",
			Handler:     userSettingHandler.ResetUserSetting,
			Operation:   "self:settings:reset",
			Tags:        "User - Settings",
			Summary:     "重置配置",
			Description: "重置用户配置为默认值",
		},
	}...)

	return allRoutes
}
