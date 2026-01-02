package modules

import (
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/crm/application/company"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/crm/application/contact"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/crm/application/lead"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/crm/application/opportunity"
)

// CRMFacade 提供 CRM 模块的稳定访问接口。
//
// CRM (Customer Relationship Management) 模块负责客户关系管理，
// 包含公司、联系人、销售线索、商机的完整销售漏斗管理。
type CRMFacade struct {
	Company     *CompanyUseCases
	Contact     *ContactUseCases
	Lead        *LeadUseCases
	Opportunity *OpportunityUseCases
}

// CompanyUseCases 公司管理用例处理器。
//
// 提供公司信息的 CRUD 操作。
type CompanyUseCases struct {
	Create *company.CreateHandler
	Update *company.UpdateHandler
	Delete *company.DeleteHandler
	Get    *company.GetHandler
	List   *company.ListHandler
}

// ContactUseCases 联系人管理用例处理器。
//
// 提供联系人信息的 CRUD 操作。
type ContactUseCases struct {
	Create *contact.CreateHandler
	Update *contact.UpdateHandler
	Delete *contact.DeleteHandler
	Get    *contact.GetHandler
	List   *contact.ListHandler
}

// LeadUseCases 销售线索管理用例处理器。
//
// 提供线索的 CRUD 操作和状态转换功能：
// 联系（Contact）→ 合格（Qualify）→ 转换为商机（Convert）或丢失（Lose）。
type LeadUseCases struct {
	// CRUD Operations
	Create *lead.CreateHandler
	Update *lead.UpdateHandler
	Delete *lead.DeleteHandler
	Get    *lead.GetHandler
	List   *lead.ListHandler

	// State Transitions
	Contact *lead.ContactHandler // 联系线索
	Qualify *lead.QualifyHandler // 标记为合格线索
	Convert *lead.ConvertHandler // 转换为商机
	Lose    *lead.LoseHandler    // 标记为丢失
}

// OpportunityUseCases 商机管理用例处理器。
//
// 提供商机的 CRUD 操作和状态转换功能：
// 推进阶段（Advance）→ 赢单（CloseWon）或丢单（CloseLost）。
type OpportunityUseCases struct {
	// CRUD Operations
	Create *opportunity.CreateHandler
	Update *opportunity.UpdateHandler
	Delete *opportunity.DeleteHandler
	Get    *opportunity.GetHandler
	List   *opportunity.ListHandler

	// State Transitions
	Advance   *opportunity.AdvanceHandler   // 推进到下一阶段
	CloseWon  *opportunity.CloseWonHandler  // 赢单
	CloseLost *opportunity.CloseLostHandler // 丢单
}
