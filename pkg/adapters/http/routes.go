package http

import (
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/adapters/http/routes"
)

// AllRoutes 返回所有路由定义
// 路由顺序决定注册顺序，对于同路径前缀的路由需注意顺序（具体路径在参数路径之前）
func (deps *RouterDependencies) AllRoutes() []routes.Route {
	var allRoutes []routes.Route

	// Public 域 - 公开路由（无需认证）
	allRoutes = append(allRoutes, deps.publicRoutes()...)

	// Auth 域 - 认证路由（2FA）
	allRoutes = append(allRoutes, deps.authRoutes()...)

	// Admin 域 - 系统管理路由（用户、角色、组织、系统设置、缓存等）
	allRoutes = append(allRoutes, deps.adminRoutes()...)

	// Self 域 - 用户自服务路由（资料、令牌、配置）
	allRoutes = append(allRoutes, deps.selfRoutes()...)

	// Org 域 - 组织和团队路由（组织、成员、团队、任务）
	allRoutes = append(allRoutes, deps.orgRoutes()...)

	// CRM 域 - 客户关系管理路由（联系人、公司、线索、商机）
	allRoutes = append(allRoutes, deps.crmRoutes()...)

	return allRoutes
}
