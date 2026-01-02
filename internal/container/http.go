package container

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"

	"github.com/lwmacct/260101-go-pkg-ddd/ddd/config"
	ginhttp "github.com/lwmacct/260101-go-pkg-ddd/ddd/core/adapters/http"
	corehandler "github.com/lwmacct/260101-go-pkg-ddd/ddd/core/adapters/http/handler"
	"github.com/lwmacct/260101-go-pkg-ddd/ddd/core/application/cache"
	"github.com/lwmacct/260101-go-pkg-ddd/ddd/core/infrastructure/health"
	"github.com/lwmacct/260101-go-pkg-ddd/ddd/core/infrastructure/persistence"
	crmhandler "github.com/lwmacct/260101-go-pkg-ddd/ddd/crm/adapters/http/handler"
	iamhandler "github.com/lwmacct/260101-go-pkg-ddd/ddd/iam/adapters/http/handler"
	"github.com/lwmacct/260101-go-pkg-ddd/ddd/iam/infrastructure/auth"
)

// HandlersResult 使用 fx.Out 批量返回所有 HTTP 处理器。
type HandlersResult struct {
	fx.Out

	Health           *corehandler.HealthHandler
	Auth             *iamhandler.AuthHandler
	Captcha          *corehandler.CaptchaHandler
	AdminUser        *corehandler.AdminUserHandler
	UserProfile      *iamhandler.UserProfileHandler
	Role             *iamhandler.RoleHandler
	Setting          *corehandler.SettingHandler
	UserSetting      *corehandler.UserSettingHandler
	PAT              *iamhandler.PATHandler
	Audit            *corehandler.AuditHandler
	Overview         *corehandler.OverviewHandler
	TwoFA            *iamhandler.TwoFAHandler
	Cache            *corehandler.CacheHandler
	Operation        *corehandler.OperationHandler
	Organization     *corehandler.OrgHandler
	OrgMember        *corehandler.OrgMemberHandler
	Team             *corehandler.TeamHandler
	TeamMember       *corehandler.TeamMemberHandler
	UserOrganization *iamhandler.UserOrgHandler
	Task             *corehandler.TaskHandler
	Contact          *crmhandler.ContactHandler
	Company          *crmhandler.CompanyHandler
	Lead             *crmhandler.LeadHandler
	Opportunity      *crmhandler.OpportunityHandler
}

// HTTPModule 提供 HTTP 处理器、路由和服务器。
var HTTPModule = fx.Module("http",
	fx.Provide(
		health.NewSystemChecker,
		newAllHandlers,
		newRouter,
		newHTTPServer,
	),
	fx.Invoke(startHTTPServer),
)

// newHTTPServer 创建 HTTP 服务器实例。
func newHTTPServer(router *gin.Engine, cfg *config.Config) *ginhttp.Server {
	return ginhttp.NewServer(router, cfg.Server.Addr)
}

// startHTTPServer 注册 HTTP 服务器启动和关闭钩子。
func startHTTPServer(lc fx.Lifecycle, server *ginhttp.Server, cfg *config.Config) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			slog.Info("Starting HTTP server", "addr", cfg.Server.Addr, "env", cfg.Server.Env)

			// 在 goroutine 中启动服务器，避免阻塞 OnStart
			go func() {
				if err := server.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					slog.Error("HTTP server error", "error", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			slog.Info("Shutting down HTTP server")
			return server.Shutdown(ctx)
		},
	})
}

// handlersParams 聚合创建 Handler 所需的依赖。
type handlersParams struct {
	fx.In

	Config        *config.Config
	AdminCacheSvc cache.AdminCacheService
	HealthChecker *health.SystemChecker
	Auth          *AuthUseCases
	User          *UserUseCases
	Role          *RoleUseCases
	Setting       *SettingUseCases
	UserSetting   *UserSettingUseCases
	PAT           *PATUseCases
	Audit         *AuditUseCases
	Stats         *StatsUseCases
	Captcha       *CaptchaUseCases
	TwoFA         *TwoFAUseCases
	Organization  *OrganizationUseCases
	Task          *TaskUseCases
	Contact       *ContactUseCases
	Company       *CompanyUseCases
	Lead          *LeadUseCases
	Opportunity   *OpportunityUseCases
}

func newAllHandlers(p handlersParams) HandlersResult {
	return HandlersResult{
		Health: corehandler.NewHealthHandler(p.HealthChecker),
		Auth: iamhandler.NewAuthHandler(
			p.Auth.Login,
			p.Auth.Login2FA,
			p.Auth.Register,
			p.Auth.RefreshToken,
		),
		Captcha: corehandler.NewCaptchaHandler(p.Captcha.Generate, p.Config.Auth.DevSecret),
		AdminUser: corehandler.NewAdminUserHandler(
			p.User.Create,
			p.User.Update,
			p.User.Delete,
			p.User.AssignRoles,
			p.User.BatchCreate,
			p.User.Get,
			p.User.List,
		),
		UserProfile: iamhandler.NewUserProfileHandler(
			p.User.Get,
			p.User.Update,
			p.User.ChangePassword,
			p.User.Delete,
		),
		Role: iamhandler.NewRoleHandler(
			p.Role.Create,
			p.Role.Update,
			p.Role.Delete,
			p.Role.SetPermissions,
			p.Role.Get,
			p.Role.List,
		),
		Setting: corehandler.NewSettingHandler(
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
		UserSetting: corehandler.NewUserSettingHandler(
			p.UserSetting.Set,
			p.UserSetting.BatchSet,
			p.UserSetting.Reset,
			p.UserSetting.ResetAll,
			p.UserSetting.Get,
			p.UserSetting.List,
			p.UserSetting.ListSettings,
			p.UserSetting.ListCategories,
		),
		PAT: iamhandler.NewPATHandler(
			p.PAT.Create,
			p.PAT.Delete,
			p.PAT.Disable,
			p.PAT.Enable,
			p.PAT.Get,
			p.PAT.List,
		),
		Audit: corehandler.NewAuditHandler(
			p.Audit.List,
			p.Audit.Get,
		),
		Overview: corehandler.NewOverviewHandler(p.Stats.GetStats),
		TwoFA: iamhandler.NewTwoFAHandler(
			p.TwoFA.Setup,
			p.TwoFA.VerifyEnable,
			p.TwoFA.Disable,
			p.TwoFA.GetStatus,
		),
		Cache: corehandler.NewCacheHandler(
			cache.NewInfoHandler(p.AdminCacheSvc),
			cache.NewScanKeysHandler(p.AdminCacheSvc),
			cache.NewGetKeyHandler(p.AdminCacheSvc),
			cache.NewDeleteHandler(p.AdminCacheSvc),
		),
		Operation: corehandler.NewOperationHandler(),
		Organization: corehandler.NewOrgHandler(
			p.Organization.Create,
			p.Organization.Update,
			p.Organization.Delete,
			p.Organization.Get,
			p.Organization.List,
		),
		OrgMember: corehandler.NewOrgMemberHandler(
			p.Organization.MemberAdd,
			p.Organization.MemberRemove,
			p.Organization.MemberUpdateRole,
			p.Organization.MemberList,
		),
		Team: corehandler.NewTeamHandler(
			p.Organization.TeamCreate,
			p.Organization.TeamUpdate,
			p.Organization.TeamDelete,
			p.Organization.TeamGet,
			p.Organization.TeamList,
		),
		TeamMember: corehandler.NewTeamMemberHandler(
			p.Organization.TeamMemberAdd,
			p.Organization.TeamMemberRemove,
			p.Organization.TeamMemberList,
		),
		UserOrganization: iamhandler.NewUserOrgHandler(
			p.Organization.UserOrgs,
			p.Organization.UserTeams,
		),
		Task: corehandler.NewTaskHandler(
			p.Task.Create,
			p.Task.Update,
			p.Task.Delete,
			p.Task.Get,
			p.Task.List,
		),
		Contact: crmhandler.NewContactHandler(
			p.Contact.Create,
			p.Contact.Update,
			p.Contact.Delete,
			p.Contact.Get,
			p.Contact.List,
		),
		Company: crmhandler.NewCompanyHandler(
			p.Company.Create,
			p.Company.Update,
			p.Company.Delete,
			p.Company.Get,
			p.Company.List,
		),
		Lead: crmhandler.NewLeadHandler(
			p.Lead.Create,
			p.Lead.Update,
			p.Lead.Delete,
			p.Lead.Contact,
			p.Lead.Qualify,
			p.Lead.Convert,
			p.Lead.Lose,
			p.Lead.Get,
			p.Lead.List,
		),
		Opportunity: crmhandler.NewOpportunityHandler(
			p.Opportunity.Create,
			p.Opportunity.Update,
			p.Opportunity.Delete,
			p.Opportunity.Advance,
			p.Opportunity.CloseWon,
			p.Opportunity.CloseLost,
			p.Opportunity.Get,
			p.Opportunity.List,
		),
	}
}

// routerParams 聚合创建路由所需的依赖。
type routerParams struct {
	fx.In

	Config      *config.Config
	RedisClient *redis.Client

	// Services
	JWTManager      *auth.JWTManager
	PATService      *auth.PATService
	PermissionCache *auth.PermissionCacheService

	// UseCases
	Audit *AuditUseCases

	// Repositories (for middleware)
	MemberRepos     persistence.OrgMemberRepositories
	TeamRepos       persistence.TeamRepositories
	TeamMemberRepos persistence.TeamMemberRepositories

	// Handlers
	Health      *corehandler.HealthHandler
	Auth        *iamhandler.AuthHandler
	Captcha     *corehandler.CaptchaHandler
	AdminUser   *corehandler.AdminUserHandler
	UserProfile *iamhandler.UserProfileHandler
	Role        *iamhandler.RoleHandler
	Setting     *corehandler.SettingHandler
	UserSetting *corehandler.UserSettingHandler
	PAT         *iamhandler.PATHandler
	AuditH      *corehandler.AuditHandler
	Overview    *corehandler.OverviewHandler
	TwoFA       *iamhandler.TwoFAHandler
	Cache       *corehandler.CacheHandler
	Operation   *corehandler.OperationHandler
	Org         *corehandler.OrgHandler
	OrgMember   *corehandler.OrgMemberHandler
	Team        *corehandler.TeamHandler
	TeamMember  *corehandler.TeamMemberHandler
	UserOrg     *iamhandler.UserOrgHandler
	TaskHandler *corehandler.TaskHandler
	Contact     *crmhandler.ContactHandler
	Company     *crmhandler.CompanyHandler
	Lead        *crmhandler.LeadHandler
	Opportunity *crmhandler.OpportunityHandler
}

func newRouter(p routerParams) *gin.Engine {
	deps := &ginhttp.RouterDependencies{
		Config:                 p.Config,
		RedisClient:            p.RedisClient,
		CreateLogHandler:       p.Audit.CreateLog,
		JWTManager:             p.JWTManager,
		PATService:             p.PATService,
		PermissionCacheService: p.PermissionCache,
		OrgMemberQuery:         p.MemberRepos.Query,
		TeamQuery:              p.TeamRepos.Query,
		TeamMemberQuery:        p.TeamMemberRepos.Query,
		HealthHandler:          p.Health,
		AuthHandler:            p.Auth,
		CaptchaHandler:         p.Captcha,
		RoleHandler:            p.Role,
		SettingHandler:         p.Setting,
		UserSettingHandler:     p.UserSetting,
		PATHandler:             p.PAT,
		AuditHandler:           p.AuditH,
		AdminUserHandler:       p.AdminUser,
		UserProfileHandler:     p.UserProfile,
		OverviewHandler:        p.Overview,
		TwoFAHandler:           p.TwoFA,
		CacheHandler:           p.Cache,
		OperationHandler:       p.Operation,
		OrgHandler:             p.Org,
		OrgMemberHandler:       p.OrgMember,
		TeamHandler:            p.Team,
		TeamMemberHandler:      p.TeamMember,
		UserOrgHandler:         p.UserOrg,
		TaskHandler:            p.TaskHandler,
		ContactHandler:         p.Contact,
		CompanyHandler:         p.Company,
		LeadHandler:            p.Lead,
		OpportunityHandler:     p.Opportunity,
	}

	return ginhttp.SetupRouterWithDeps(deps)
}
