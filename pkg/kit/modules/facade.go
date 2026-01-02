package modules

// Facade 聚合所有业务模块的访问接口。
//
// 这是 pkg/kit/modules 包的顶层入口，将 IAM、App、CRM 三个
// Bounded Context 的 UseCase Handler 统一暴露给外部项目。
//
// # 使用示例
//
//	// 通过 pkg/kit/app 构建器获取 Facade 实例
//	facade := app.MustBuild(cfg)
//
//	// 使用 IAM 模块进行用户认证
//	ctx := context.Background()
//	loginResp, err := facade.IAM.Auth.Login.Handle(ctx, loginCmd)
//
//	// 使用 App 模块管理组织
//	org, err := facade.App.Organization.Get.Handle(ctx, getOrgQuery)
//
//	// 使用 CRM 模块管理商机
//	opps, err := facade.CRM.Opportunity.List.Handle(ctx, listOppsQuery)
type Facade struct {
	IAM *IAMFacade
	App *AppFacade
	CRM *CRMFacade
}

// NewFacade 创建 Facade 实例。
//
// 此构造函数由 pkg/kit/app 包的构建器调用，
// 外部项目不应直接使用此函数。
func NewFacade(iam *IAMFacade, app *AppFacade, crm *CRMFacade) *Facade {
	return &Facade{
		IAM: iam,
		App: app,
		CRM: crm,
	}
}
