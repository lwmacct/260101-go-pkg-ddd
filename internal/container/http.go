package container

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"

	ginhttp "github.com/lwmacct/260101-go-pkg-ddd/pkg/adapters/http"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/adapters/http/handler"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/application/cache"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/config"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/infrastructure/auth"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/infrastructure/health"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/infrastructure/persistence"
)

// HandlersResult 使用 fx.Out 批量返回所有 HTTP 处理器。
type HandlersResult struct {
	fx.Out

	Health           *handler.HealthHandler
	Auth             *handler.AuthHandler
	Captcha          *handler.CaptchaHandler
	AdminUser        *handler.AdminUserHandler
	UserProfile      *handler.UserProfileHandler
	Role             *handler.RoleHandler
	Setting          *handler.SettingHandler
	UserSetting      *handler.UserSettingHandler
	PAT              *handler.PATHandler
	Audit            *handler.AuditHandler
	Overview         *handler.OverviewHandler
	TwoFA            *handler.TwoFAHandler
	Cache            *handler.CacheHandler
	Operation        *handler.OperationHandler
	Organization     *handler.OrgHandler
	OrgMember        *handler.OrgMemberHandler
	Team             *handler.TeamHandler
	TeamMember       *handler.TeamMemberHandler
	UserOrganization *handler.UserOrgHandler
	Task             *handler.TaskHandler
	Contact          *handler.ContactHandler
	Company          *handler.CompanyHandler
	Lead             *handler.LeadHandler
	Opportunity      *handler.OpportunityHandler
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
		Health: handler.NewHealthHandler(p.HealthChecker),
		Auth: handler.NewAuthHandler(
			p.Auth.Login,
			p.Auth.Login2FA,
			p.Auth.Register,
			p.Auth.RefreshToken,
		),
		Captcha: handler.NewCaptchaHandler(p.Captcha.Generate, p.Config.Auth.DevSecret),
		AdminUser: handler.NewAdminUserHandler(
			p.User.Create,
			p.User.Update,
			p.User.Delete,
			p.User.AssignRoles,
			p.User.BatchCreate,
			p.User.Get,
			p.User.List,
		),
		UserProfile: handler.NewUserProfileHandler(
			p.User.Get,
			p.User.Update,
			p.User.ChangePassword,
			p.User.Delete,
		),
		Role: handler.NewRoleHandler(
			p.Role.Create,
			p.Role.Update,
			p.Role.Delete,
			p.Role.SetPermissions,
			p.Role.Get,
			p.Role.List,
		),
		Setting: handler.NewSettingHandler(
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
		UserSetting: handler.NewUserSettingHandler(
			p.UserSetting.Set,
			p.UserSetting.BatchSet,
			p.UserSetting.Reset,
			p.UserSetting.ResetAll,
			p.UserSetting.Get,
			p.UserSetting.List,
			p.UserSetting.ListSettings,
			p.UserSetting.ListCategories,
		),
		PAT: handler.NewPATHandler(
			p.PAT.Create,
			p.PAT.Delete,
			p.PAT.Disable,
			p.PAT.Enable,
			p.PAT.Get,
			p.PAT.List,
		),
		Audit: handler.NewAuditHandler(
			p.Audit.List,
			p.Audit.Get,
		),
		Overview: handler.NewOverviewHandler(p.Stats.GetStats),
		TwoFA: handler.NewTwoFAHandler(
			p.TwoFA.Setup,
			p.TwoFA.VerifyEnable,
			p.TwoFA.Disable,
			p.TwoFA.GetStatus,
		),
		Cache: handler.NewCacheHandler(
			cache.NewInfoHandler(p.AdminCacheSvc),
			cache.NewScanKeysHandler(p.AdminCacheSvc),
			cache.NewGetKeyHandler(p.AdminCacheSvc),
			cache.NewDeleteHandler(p.AdminCacheSvc),
		),
		Operation: handler.NewOperationHandler(),
		Organization: handler.NewOrgHandler(
			p.Organization.Create,
			p.Organization.Update,
			p.Organization.Delete,
			p.Organization.Get,
			p.Organization.List,
		),
		OrgMember: handler.NewOrgMemberHandler(
			p.Organization.MemberAdd,
			p.Organization.MemberRemove,
			p.Organization.MemberUpdateRole,
			p.Organization.MemberList,
		),
		Team: handler.NewTeamHandler(
			p.Organization.TeamCreate,
			p.Organization.TeamUpdate,
			p.Organization.TeamDelete,
			p.Organization.TeamGet,
			p.Organization.TeamList,
		),
		TeamMember: handler.NewTeamMemberHandler(
			p.Organization.TeamMemberAdd,
			p.Organization.TeamMemberRemove,
			p.Organization.TeamMemberList,
		),
		UserOrganization: handler.NewUserOrgHandler(
			p.Organization.UserOrgs,
			p.Organization.UserTeams,
		),
		Task: handler.NewTaskHandler(
			p.Task.Create,
			p.Task.Update,
			p.Task.Delete,
			p.Task.Get,
			p.Task.List,
		),
		Contact: handler.NewContactHandler(
			p.Contact.Create,
			p.Contact.Update,
			p.Contact.Delete,
			p.Contact.Get,
			p.Contact.List,
		),
		Company: handler.NewCompanyHandler(
			p.Company.Create,
			p.Company.Update,
			p.Company.Delete,
			p.Company.Get,
			p.Company.List,
		),
		Lead: handler.NewLeadHandler(
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
		Opportunity: handler.NewOpportunityHandler(
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
	Health      *handler.HealthHandler
	Auth        *handler.AuthHandler
	Captcha     *handler.CaptchaHandler
	AdminUser   *handler.AdminUserHandler
	UserProfile *handler.UserProfileHandler
	Role        *handler.RoleHandler
	Setting     *handler.SettingHandler
	UserSetting *handler.UserSettingHandler
	PAT         *handler.PATHandler
	AuditH      *handler.AuditHandler
	Overview    *handler.OverviewHandler
	TwoFA       *handler.TwoFAHandler
	Cache       *handler.CacheHandler
	Operation   *handler.OperationHandler
	Org         *handler.OrgHandler
	OrgMember   *handler.OrgMemberHandler
	Team        *handler.TeamHandler
	TeamMember  *handler.TeamMemberHandler
	UserOrg     *handler.UserOrgHandler
	TaskHandler *handler.TaskHandler
	Contact     *handler.ContactHandler
	Company     *handler.CompanyHandler
	Lead        *handler.LeadHandler
	Opportunity *handler.OpportunityHandler
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
