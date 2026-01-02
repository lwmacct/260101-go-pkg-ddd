package gin

import (
	corehandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/core/transport/gin/handler"
	iamhandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/iam/transport/gin/handler"
	platformhttp "github.com/lwmacct/260101-go-pkg-ddd/pkg/platform/http"
)

// IAMRoutes 返回 IAM 域的所有路由
//
// 参数：各模块的 handler（按需传递，避免依赖 god object）
func IAMRoutes(
	// Auth handlers
	authHandler *iamhandler.AuthHandler,
	twoFAHandler *iamhandler.TwoFAHandler,

	// User handlers
	userProfileHandler *iamhandler.UserProfileHandler,
	userOrgHandler *iamhandler.UserOrgHandler,

	// Admin handlers
	adminUserHandler *corehandler.AdminUserHandler,
	roleHandler *iamhandler.RoleHandler,
	patHandler *iamhandler.PATHandler,

	// Core handlers (IAM uses some core handlers)
	settingHandler *corehandler.SettingHandler,
	userSettingHandler *corehandler.UserSettingHandler,
) []platformhttp.Route {
	var routes []platformhttp.Route

	// Public routes (auth)
	routes = append(routes, PublicRoutes(authHandler)...)

	// Auth routes (2FA)
	routes = append(routes, AuthRoutes(twoFAHandler)...)

	// Self routes (user profile, PAT, settings)
	routes = append(routes, SelfRoutes(
		userProfileHandler,
		patHandler,
		settingHandler,
		userSettingHandler,
	)...)

	// Admin routes (user/role management)
	routes = append(routes, AdminRoutes(
		adminUserHandler,
		roleHandler,
	)...)

	// Org routes (user's org view)
	routes = append(routes, UserOrgRoutes(userOrgHandler)...)

	return routes
}
