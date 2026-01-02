package handler

import (
	"go.uber.org/fx"

	"github.com/lwmacct/260101-go-pkg-ddd/pkg/config"
	appapplication "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/app/application"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/app/application/cache"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/platform/health"
)

// HandlersResult 使用 fx.Out 批量返回 App 模块的所有 HTTP 处理器。
type HandlersResult struct {
	fx.Out

	Health       *HealthHandler
	Captcha      *CaptchaHandler
	Setting      *SettingHandler
	UserSetting  *UserSettingHandler
	Audit        *AuditHandler
	Overview     *OverviewHandler
	Cache        *CacheHandler
	Operation    *OperationHandler
	Organization *OrgHandler
	OrgMember    *OrgMemberHandler
	Team         *TeamHandler
	TeamMember   *TeamMemberHandler
	Task         *TaskHandler
}

// HandlerModule 提供 App 模块的所有 HTTP 处理器。
//
// 注意：AdminUserHandler 因跨模块依赖（需要 IAM 的 UserUseCases）而在顶层容器注册。
var HandlerModule = fx.Module("app.handler",
	fx.Provide(newAllHandlers),
)

// handlersParams 聚合创建 Handler 所需的依赖。
type handlersParams struct {
	fx.In

	Config        *config.Config
	AdminCacheSvc cache.AdminCacheService
	HealthChecker *health.SystemChecker

	// App 模块用例
	Audit        *appapplication.AuditUseCases
	Setting      *appapplication.SettingUseCases
	UserSetting  *appapplication.UserSettingUseCases
	Stats        *appapplication.StatsUseCases
	Captcha      *appapplication.CaptchaUseCases
	Organization *appapplication.OrganizationUseCases
	Task         *appapplication.TaskUseCases
}

func newAllHandlers(p handlersParams) HandlersResult {
	return HandlersResult{
		Health:  NewHealthHandler(p.HealthChecker),
		Captcha: NewCaptchaHandler(p.Captcha.Generate, p.Config.Auth.DevSecret),
		Setting: NewSettingHandler(
			p.Setting.Create,
			p.Setting.Update,
			p.Setting.Delete,
			p.Setting.BatchUpdate,
			p.Setting.Get,
			p.Setting.List,
			p.Setting.ListSettings,
			p.Setting.CreateCategory,
			p.Setting.UpdateCategory,
			p.Setting.DeleteCategory,
			p.Setting.GetCategory,
			p.Setting.ListCategories,
		),
		UserSetting: NewUserSettingHandler(
			p.UserSetting.Set,
			p.UserSetting.BatchSet,
			p.UserSetting.Reset,
			p.UserSetting.ResetAll,
			p.UserSetting.Get,
			p.UserSetting.List,
			p.UserSetting.ListSettings,
			p.UserSetting.ListCategories,
		),
		Audit: NewAuditHandler(
			p.Audit.List,
			p.Audit.Get,
		),
		Overview: NewOverviewHandler(p.Stats.GetStats),
		Cache: NewCacheHandler(
			cache.NewInfoHandler(p.AdminCacheSvc),
			cache.NewScanKeysHandler(p.AdminCacheSvc),
			cache.NewGetKeyHandler(p.AdminCacheSvc),
			cache.NewDeleteHandler(p.AdminCacheSvc),
		),
		Operation: NewOperationHandler(),
		Organization: NewOrgHandler(
			p.Organization.Create,
			p.Organization.Update,
			p.Organization.Delete,
			p.Organization.Get,
			p.Organization.List,
		),
		OrgMember: NewOrgMemberHandler(
			p.Organization.MemberAdd,
			p.Organization.MemberRemove,
			p.Organization.MemberUpdateRole,
			p.Organization.MemberList,
		),
		Team: NewTeamHandler(
			p.Organization.TeamCreate,
			p.Organization.TeamUpdate,
			p.Organization.TeamDelete,
			p.Organization.TeamGet,
			p.Organization.TeamList,
		),
		TeamMember: NewTeamMemberHandler(
			p.Organization.TeamMemberAdd,
			p.Organization.TeamMemberRemove,
			p.Organization.TeamMemberList,
		),
		Task: NewTaskHandler(
			p.Task.Create,
			p.Task.Update,
			p.Task.Delete,
			p.Task.Get,
			p.Task.List,
		),
	}
}
