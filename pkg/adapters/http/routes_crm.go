package http

import (
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/adapters/http/routes"
)

// crmRoutes 返回 CRM 域路由配置。
func (deps *RouterDependencies) crmRoutes() []routes.Route {
	// 基础中间件：认证 + RBAC
	baseMiddlewares := []routes.MiddlewareConfig{
		{Name: routes.MiddlewareRequestID},
		{Name: routes.MiddlewareOperationID},
		{Name: routes.MiddlewareAuth},
		{Name: routes.MiddlewareRBAC},
	}

	// 审计中间件：需要记录操作日志
	auditMiddlewares := append(cloneMiddlewares(baseMiddlewares),
		routes.MiddlewareConfig{Name: routes.MiddlewareAudit})

	return []routes.Route{
		// ==================== 联系人 CRUD ====================
		{
			Method:      routes.POST,
			Path:        "/api/crm/contacts",
			Handler:     deps.ContactHandler.Create,
			Op:          "crm:contacts:create",
			Middlewares: auditMiddlewares,
			Tags:        "CRM - Contacts",
			Summary:     "创建联系人",
			Description: "创建新联系人",
		},
		{
			Method:      routes.GET,
			Path:        "/api/crm/contacts",
			Handler:     deps.ContactHandler.List,
			Op:          "crm:contacts:list",
			Middlewares: baseMiddlewares,
			Tags:        "CRM - Contacts",
			Summary:     "联系人列表",
			Description: "分页获取联系人列表",
		},
		{
			Method:      routes.GET,
			Path:        "/api/crm/contacts/:id",
			Handler:     deps.ContactHandler.Get,
			Op:          "crm:contacts:get",
			Middlewares: baseMiddlewares,
			Tags:        "CRM - Contacts",
			Summary:     "联系人详情",
			Description: "获取联系人详细信息",
		},
		{
			Method:      routes.PUT,
			Path:        "/api/crm/contacts/:id",
			Handler:     deps.ContactHandler.Update,
			Op:          "crm:contacts:update",
			Middlewares: auditMiddlewares,
			Tags:        "CRM - Contacts",
			Summary:     "更新联系人",
			Description: "更新联系人信息",
		},
		{
			Method:      routes.DELETE,
			Path:        "/api/crm/contacts/:id",
			Handler:     deps.ContactHandler.Delete,
			Op:          "crm:contacts:delete",
			Middlewares: auditMiddlewares,
			Tags:        "CRM - Contacts",
			Summary:     "删除联系人",
			Description: "删除联系人",
		},
	}
}
