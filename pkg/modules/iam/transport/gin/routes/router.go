// Package routes 定义 IAM 模块的所有 HTTP 路由。
//
// 本包遵循 DDD 架构原则，BC 层只负责定义路由结构和 Handler，
// 中间件由应用层（internal/container）注入。
//
// # 路由分组
//
//   - [Public]: 公开路由（登录、注册、验证码）
//   - [Auth]: 认证路由（双因素认证）
//   - [Self]: 用户自服务路由（个人资料、PAT、设置）
//   - [Admin]: 管理员路由（用户管理、角色管理）
//   - [Org]: 组织管理路由（成员、团队、任务）
//   - [UserOrg]: 用户组织视图路由
//
// # 使用方式
//
//	routes := iamroutes.All(authHandler, twoFAHandler, ...)
//	injector := di.NewMiddlewareInjector(deps)
//	di.RegisterRoutesV2(engine, routes, injector)
package routes

import (
	"github.com/lwmacct/260101-go-pkg-gin/pkg/routes"

	corehandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/app/transport/gin/handler"
	iamhandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/iam/transport/gin/handler"
)

// All 返回 IAM 域的所有路由
//
// 参数：各模块的 handler（按需传递，避免依赖 god object）
func All(
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

	// Org handlers (organization/team management)
	orgMemberHandler *corehandler.OrgMemberHandler,
	teamHandler *corehandler.TeamHandler,
	teamMemberHandler *corehandler.TeamMemberHandler,
	taskHandler *corehandler.TaskHandler,

	// Core handlers (IAM uses some core handlers)
	captchaHandler *corehandler.CaptchaHandler,
	settingHandler *corehandler.SettingHandler,
	userSettingHandler *corehandler.UserSettingHandler,
) []routes.Route {
	var allRoutes []routes.Route

	// Public routes (auth)
	allRoutes = append(allRoutes, Public(
		authHandler,
		captchaHandler,
	)...)

	// Auth routes (2FA)
	allRoutes = append(allRoutes, Auth(twoFAHandler)...)

	// Self routes (user profile, PAT, settings)
	allRoutes = append(allRoutes, Self(
		userProfileHandler,
		patHandler,
		settingHandler,
		userSettingHandler,
	)...)

	// Admin routes (user/role/setting management)
	allRoutes = append(allRoutes, Admin(
		adminUserHandler,
		roleHandler,
		settingHandler,
	)...)

	// Org routes (organization/team/task management)
	allRoutes = append(allRoutes, Org(
		orgMemberHandler,
		teamHandler,
		teamMemberHandler,
		taskHandler,
	)...)

	// Org routes (user's org view)
	allRoutes = append(allRoutes, UserOrg(userOrgHandler)...)

	return allRoutes
}
