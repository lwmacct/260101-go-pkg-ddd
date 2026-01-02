package gin

import (
	crmhandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/crm/transport/gin/handler"
	platformhttp "github.com/lwmacct/260101-go-pkg-ddd/pkg/platform/http"
)

// CRMRoutes 返回 CRM 域的所有路由
func CRMRoutes(
	companyHandler *crmhandler.CompanyHandler,
	contactHandler *crmhandler.ContactHandler,
	leadHandler *crmhandler.LeadHandler,
	opportunityHandler *crmhandler.OpportunityHandler,
) []platformhttp.Route {
	return Routes(
		companyHandler,
		contactHandler,
		leadHandler,
		opportunityHandler,
	)
}
