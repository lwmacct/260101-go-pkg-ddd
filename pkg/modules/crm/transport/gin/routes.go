package gin

import (
	crmhandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/crm/transport/gin/handler"
	platformhttp "github.com/lwmacct/260101-go-pkg-ddd/pkg/platform/http"
)

// Routes 返回 CRM 域的所有路由
func Routes(
	companyHandler *crmhandler.CompanyHandler,
	contactHandler *crmhandler.ContactHandler,
	leadHandler *crmhandler.LeadHandler,
	opportunityHandler *crmhandler.OpportunityHandler,
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

	// Company routes
	routes = append(routes, []platformhttp.Route{
		{
			Method:      platformhttp.GET,
			Path:        "/api/crm/companies",
			Handler:     companyHandler.List,
			Op:          "crm:companies:list",
			Middlewares: baseMiddlewares,
			Tags:        "CRM - Companies",
			Summary:     "公司列表",
			Description: "获取公司列表",
		},
		{
			Method:      platformhttp.GET,
			Path:        "/api/crm/companies/:id",
			Handler:     companyHandler.Get,
			Op:          "crm:companies:get",
			Middlewares: baseMiddlewares,
			Tags:        "CRM - Companies",
			Summary:     "公司详情",
			Description: "获取公司详细信息",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/crm/companies",
			Handler:     companyHandler.Create,
			Op:          "crm:companies:create",
			Middlewares: auditMiddlewares,
			Tags:        "CRM - Companies",
			Summary:     "创建公司",
			Description: "创建新公司",
		},
		{
			Method:      platformhttp.PUT,
			Path:        "/api/crm/companies/:id",
			Handler:     companyHandler.Update,
			Op:          "crm:companies:update",
			Middlewares: auditMiddlewares,
			Tags:        "CRM - Companies",
			Summary:     "更新公司",
			Description: "更新公司信息",
		},
		{
			Method:      platformhttp.DELETE,
			Path:        "/api/crm/companies/:id",
			Handler:     companyHandler.Delete,
			Op:          "crm:companies:delete",
			Middlewares: auditMiddlewares,
			Tags:        "CRM - Companies",
			Summary:     "删除公司",
			Description: "删除公司",
		},
	}...)

	// Contact routes
	routes = append(routes, []platformhttp.Route{
		{
			Method:      platformhttp.GET,
			Path:        "/api/crm/contacts",
			Handler:     contactHandler.List,
			Op:          "crm:contacts:list",
			Middlewares: baseMiddlewares,
			Tags:        "CRM - Contacts",
			Summary:     "联系人列表",
			Description: "获取联系人列表",
		},
		{
			Method:      platformhttp.GET,
			Path:        "/api/crm/contacts/:id",
			Handler:     contactHandler.Get,
			Op:          "crm:contacts:get",
			Middlewares: baseMiddlewares,
			Tags:        "CRM - Contacts",
			Summary:     "联系人详情",
			Description: "获取联系人详细信息",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/crm/contacts",
			Handler:     contactHandler.Create,
			Op:          "crm:contacts:create",
			Middlewares: auditMiddlewares,
			Tags:        "CRM - Contacts",
			Summary:     "创建联系人",
			Description: "创建新联系人",
		},
		{
			Method:      platformhttp.PUT,
			Path:        "/api/crm/contacts/:id",
			Handler:     contactHandler.Update,
			Op:          "crm:contacts:update",
			Middlewares: auditMiddlewares,
			Tags:        "CRM - Contacts",
			Summary:     "更新联系人",
			Description: "更新联系人信息",
		},
		{
			Method:      platformhttp.DELETE,
			Path:        "/api/crm/contacts/:id",
			Handler:     contactHandler.Delete,
			Op:          "crm:contacts:delete",
			Middlewares: auditMiddlewares,
			Tags:        "CRM - Contacts",
			Summary:     "删除联系人",
			Description: "删除联系人",
		},
	}...)

	// Lead routes
	routes = append(routes, []platformhttp.Route{
		{
			Method:      platformhttp.GET,
			Path:        "/api/crm/leads",
			Handler:     leadHandler.List,
			Op:          "crm:leads:list",
			Middlewares: baseMiddlewares,
			Tags:        "CRM - Leads",
			Summary:     "线索列表",
			Description: "获取线索列表",
		},
		{
			Method:      platformhttp.GET,
			Path:        "/api/crm/leads/:id",
			Handler:     leadHandler.Get,
			Op:          "crm:leads:get",
			Middlewares: baseMiddlewares,
			Tags:        "CRM - Leads",
			Summary:     "线索详情",
			Description: "获取线索详细信息",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/crm/leads",
			Handler:     leadHandler.Create,
			Op:          "crm:leads:create",
			Middlewares: auditMiddlewares,
			Tags:        "CRM - Leads",
			Summary:     "创建线索",
			Description: "创建新线索",
		},
		{
			Method:      platformhttp.PUT,
			Path:        "/api/crm/leads/:id",
			Handler:     leadHandler.Update,
			Op:          "crm:leads:update",
			Middlewares: auditMiddlewares,
			Tags:        "CRM - Leads",
			Summary:     "更新线索",
			Description: "更新线索信息",
		},
		{
			Method:      platformhttp.DELETE,
			Path:        "/api/crm/leads/:id",
			Handler:     leadHandler.Delete,
			Op:          "crm:leads:delete",
			Middlewares: auditMiddlewares,
			Tags:        "CRM - Leads",
			Summary:     "删除线索",
			Description: "删除线索",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/crm/leads/:id/contact",
			Handler:     leadHandler.Contact,
			Op:          "crm:leads:contact",
			Middlewares: auditMiddlewares,
			Tags:        "CRM - Leads",
			Summary:     "联系线索",
			Description: "记录线索联系",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/crm/leads/:id/qualify",
			Handler:     leadHandler.Qualify,
			Op:          "crm:leads:qualify",
			Middlewares: auditMiddlewares,
			Tags:        "CRM - Leads",
			Summary:     "资质线索",
			Description: "将线索标记为已资质",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/crm/leads/:id/convert",
			Handler:     leadHandler.Convert,
			Op:          "crm:leads:convert",
			Middlewares: auditMiddlewares,
			Tags:        "CRM - Leads",
			Summary:     "转化线索",
			Description: "将线索转化为商机",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/crm/leads/:id/lose",
			Handler:     leadHandler.Lose,
			Op:          "crm:leads:lose",
			Middlewares: auditMiddlewares,
			Tags:        "CRM - Leads",
			Summary:     "失失线索",
			Description: "将线索标记为已失失",
		},
	}...)

	// Opportunity routes
	routes = append(routes, []platformhttp.Route{
		{
			Method:      platformhttp.GET,
			Path:        "/api/crm/opportunities",
			Handler:     opportunityHandler.List,
			Op:          "crm:opportunities:list",
			Middlewares: baseMiddlewares,
			Tags:        "CRM - Opportunities",
			Summary:     "商机列表",
			Description: "获取商机列表",
		},
		{
			Method:      platformhttp.GET,
			Path:        "/api/crm/opportunities/:id",
			Handler:     opportunityHandler.Get,
			Op:          "crm:opportunities:get",
			Middlewares: baseMiddlewares,
			Tags:        "CRM - Opportunities",
			Summary:     "商机详情",
			Description: "获取商机详细信息",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/crm/opportunities",
			Handler:     opportunityHandler.Create,
			Op:          "crm:opportunities:create",
			Middlewares: auditMiddlewares,
			Tags:        "CRM - Opportunities",
			Summary:     "创建商机",
			Description: "创建新商机",
		},
		{
			Method:      platformhttp.PUT,
			Path:        "/api/crm/opportunities/:id",
			Handler:     opportunityHandler.Update,
			Op:          "crm:opportunities:update",
			Middlewares: auditMiddlewares,
			Tags:        "CRM - Opportunities",
			Summary:     "更新商机",
			Description: "更新商机信息",
		},
		{
			Method:      platformhttp.DELETE,
			Path:        "/api/crm/opportunities/:id",
			Handler:     opportunityHandler.Delete,
			Op:          "crm:opportunities:delete",
			Middlewares: auditMiddlewares,
			Tags:        "CRM - Opportunities",
			Summary:     "删除商机",
			Description: "删除商机",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/crm/opportunities/:id/advance",
			Handler:     opportunityHandler.Advance,
			Op:          "crm:opportunities:advance",
			Middlewares: auditMiddlewares,
			Tags:        "CRM - Opportunities",
			Summary:     "推进商机",
			Description: "推进商机到下一阶段",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/crm/opportunities/:id/close-won",
			Handler:     opportunityHandler.CloseWon,
			Op:          "crm:opportunities:close_won",
			Middlewares: auditMiddlewares,
			Tags:        "CRM - Opportunities",
			Summary:     "赢单关闭",
			Description: "将商机标记为赢单",
		},
		{
			Method:      platformhttp.POST,
			Path:        "/api/crm/opportunities/:id/close-lost",
			Handler:     opportunityHandler.CloseLost,
			Op:          "crm:opportunities:close_lost",
			Middlewares: auditMiddlewares,
			Tags:        "CRM - Opportunities",
			Summary:     "输单关闭",
			Description: "将商机标记为输单",
		},
	}...)

	return routes
}
