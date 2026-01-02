package di

import (
	"github.com/gin-gonic/gin"
	coregin "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/core/transport/gin"
	corehandlerpkg "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/core/transport/gin/handler"
	crmgin "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/crm/transport/gin"
	crmhandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/crm/transport/gin/handler"
	iamgin "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/iam/transport/gin"
	iamhandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/iam/transport/gin/handler"
	platformhttp "github.com/lwmacct/260101-go-pkg-ddd/pkg/platform/http"
)

// AllRoutes 返回所有模块的路由定义
func AllRoutes(
	// IAM Handlers
	authHandler *iamhandler.AuthHandler,
	twoFAHandler *iamhandler.TwoFAHandler,
	userProfileHandler *iamhandler.UserProfileHandler,
	userOrgHandler *iamhandler.UserOrgHandler,
	patHandler *iamhandler.PATHandler,
	// Core Handlers (used in IAM)
	adminUserHandler *corehandlerpkg.AdminUserHandler,
	roleHandler *iamhandler.RoleHandler,
	settingHandler *corehandlerpkg.SettingHandler,
	userSettingHandler *corehandlerpkg.UserSettingHandler,
	// Core Handlers
	orgHandler *corehandlerpkg.OrgHandler,
	orgMemberHandler *corehandlerpkg.OrgMemberHandler,
	teamHandler *corehandlerpkg.TeamHandler,
	teamMemberHandler *corehandlerpkg.TeamMemberHandler,
	taskHandler *corehandlerpkg.TaskHandler,
	auditHandler *corehandlerpkg.AuditHandler,
	healthHandler *corehandlerpkg.HealthHandler,
	cacheHandler *corehandlerpkg.CacheHandler,
	overviewHandler *corehandlerpkg.OverviewHandler,
	// CRM Handlers
	companyHandler *crmhandler.CompanyHandler,
	contactHandler *crmhandler.ContactHandler,
	leadHandler *crmhandler.LeadHandler,
	opportunityHandler *crmhandler.OpportunityHandler,
) []platformhttp.Route {
	var routes []platformhttp.Route

	// IAM 域路由
	routes = append(routes, iamgin.IAMRoutes(
		authHandler,
		twoFAHandler,
		userProfileHandler,
		userOrgHandler,
		adminUserHandler,
		roleHandler,
		patHandler,
		settingHandler,
		userSettingHandler,
	)...)

	// Core 域路由
	routes = append(routes, coregin.CoreDomainRoutes(
		orgHandler,
		orgMemberHandler,
		teamHandler,
		teamMemberHandler,
		taskHandler,
		auditHandler,
		healthHandler,
		cacheHandler,
		overviewHandler,
	)...)

	// CRM 域路由
	routes = append(routes, crmgin.CRMRoutes(
		companyHandler,
		contactHandler,
		leadHandler,
		opportunityHandler,
	)...)

	return routes
}

// RegisterRoutes 注册所有路由到 Gin Engine
func RegisterRoutes(engine *gin.Engine, routes []platformhttp.Route) {
	for _, route := range routes {
		middlewares := []gin.HandlerFunc{route.Handler}

		// TODO: Build middleware chain from MiddlewareConfig
		// For now, just register the handler directly

		switch route.Method {
		case platformhttp.GET:
			engine.GET(route.Path, middlewares...)
		case platformhttp.POST:
			engine.POST(route.Path, middlewares...)
		case platformhttp.PUT:
			engine.PUT(route.Path, middlewares...)
		case platformhttp.DELETE:
			engine.DELETE(route.Path, middlewares...)
		case platformhttp.PATCH:
			engine.PATCH(route.Path, middlewares...)
		case platformhttp.HEAD:
			engine.HEAD(route.Path, middlewares...)
		case platformhttp.OPTIONS:
			engine.OPTIONS(route.Path, middlewares...)
		}
	}
}
