package container

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"

	"github.com/lwmacct/260101-go-pkg-ddd/internal/bootstrap"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/config"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/platform/health"

	// Application UseCases and Handlers
	appApp "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/app/application"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/app/application/cache"
	corehandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/app/transport/gin/handler"

	// CRM UseCases and Handlers
	crmapplication "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/crm/application"
	crmhandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/crm/transport/gin/handler"

	// IAM UseCases and Handlers
	iamapplication "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/iam/application"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/iam/infrastructure/auth"
	iampersistence "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/iam/infrastructure/persistence"
	iamhandler "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/iam/transport/gin/handler"

	ginHttp "github.com/lwmacct/260101-go-pkg-ddd/pkg/platform/http/gin"
)

// HandlersResult 使用 fx.Out 批量返回所有 HTTP 处理器。
type HandlersResult struct {
	fx.Out

	Health           *corehandler.HealthHandler
	Auth             *iamhandler.AuthHandler
	Captcha          *iamhandler.CaptchaHandler
	AdminUser        *iamhandler.AdminUserHandler
	UserProfile      *iamhandler.UserProfileHandler
	Role             *iamhandler.RoleHandler
	Setting          *corehandler.SettingHandler
	UserSetting      *corehandler.UserSettingHandler
	PAT              *iamhandler.PATHandler
	Audit            *iamhandler.AuditHandler
	Overview         *corehandler.OverviewHandler
	TwoFA            *iamhandler.TwoFAHandler
	Cache            *corehandler.CacheHandler
	Operation        *corehandler.OperationHandler
	Organization     *iamhandler.OrgHandler
	OrgMember        *iamhandler.OrgMemberHandler
	Team             *iamhandler.TeamHandler
	TeamMember       *iamhandler.TeamMemberHandler
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
func newHTTPServer(router *gin.Engine, cfg *config.Config) *ginHttp.Server {
	return ginHttp.NewServer(router, cfg.Server.Addr)
}

// startHTTPServer 注册 HTTP 服务器启动和关闭钩子。
func startHTTPServer(lc fx.Lifecycle, server *ginHttp.Server, cfg *config.Config) {
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
	Auth          *iamapplication.AuthUseCases
	User          *iamapplication.UserUseCases
	Role          *iamapplication.RoleUseCases
	Setting       *appApp.SettingUseCases
	UserSetting   *appApp.UserSettingUseCases
	PAT           *iamapplication.PATUseCases
	Audit         *appApp.AuditUseCases
	Stats         *appApp.StatsUseCases
	Captcha       *appApp.CaptchaUseCases
	TwoFA         *iamapplication.TwoFAUseCases
	Organization  *appApp.OrganizationUseCases
	Task          *appApp.TaskUseCases
	Contact       *crmapplication.ContactUseCases
	Company       *crmapplication.CompanyUseCases
	Lead          *crmapplication.LeadUseCases
	Opportunity   *crmapplication.OpportunityUseCases
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
		Captcha: iamhandler.NewCaptchaHandler(p.Captcha.Generate, p.Config.Auth.DevSecret),
		AdminUser: iamhandler.NewAdminUserHandler(
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
		Audit: iamhandler.NewAuditHandler(
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
		Organization: iamhandler.NewOrgHandler(
			p.Organization.Create,
			p.Organization.Update,
			p.Organization.Delete,
			p.Organization.Get,
			p.Organization.List,
		),
		OrgMember: iamhandler.NewOrgMemberHandler(
			p.Organization.MemberAdd,
			p.Organization.MemberRemove,
			p.Organization.MemberUpdateRole,
			p.Organization.MemberList,
		),
		Team: iamhandler.NewTeamHandler(
			p.Organization.TeamCreate,
			p.Organization.TeamUpdate,
			p.Organization.TeamDelete,
			p.Organization.TeamGet,
			p.Organization.TeamList,
		),
		TeamMember: iamhandler.NewTeamMemberHandler(
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
	Audit *iamapplication.AuditUseCases

	// Repositories (for middleware)
	MemberRepos     iampersistence.OrgMemberRepositories
	TeamRepos       iampersistence.TeamRepositories
	TeamMemberRepos iampersistence.TeamMemberRepositories

	// Handlers
	Health      *corehandler.HealthHandler
	Auth        *iamhandler.AuthHandler
	Captcha     *iamhandler.CaptchaHandler
	AdminUser   *iamhandler.AdminUserHandler
	UserProfile *iamhandler.UserProfileHandler
	Role        *iamhandler.RoleHandler
	Setting     *corehandler.SettingHandler
	UserSetting *corehandler.UserSettingHandler
	PAT         *iamhandler.PATHandler
	AuditH      *iamhandler.AuditHandler
	Overview    *corehandler.OverviewHandler
	TwoFA       *iamhandler.TwoFAHandler
	Cache       *corehandler.CacheHandler
	Operation   *corehandler.OperationHandler
	Org         *iamhandler.OrgHandler
	OrgMember   *iamhandler.OrgMemberHandler
	Team        *iamhandler.TeamHandler
	TeamMember  *iamhandler.TeamMemberHandler
	UserOrg     *iamhandler.UserOrgHandler
	TaskHandler *corehandler.TaskHandler
	Contact     *crmhandler.ContactHandler
	Company     *crmhandler.CompanyHandler
	Lead        *crmhandler.LeadHandler
	Opportunity *crmhandler.OpportunityHandler
}

func newRouter(p routerParams) *gin.Engine {
	// Create Gin Engine using bootstrap
	engine := bootstrap.NewEngine()

	// Get all routes from modules using the new routes function
	allRoutes := AllRoutes(
		// IAM Handlers
		p.Auth,
		p.TwoFA,
		p.UserProfile,
		p.UserOrg,
		p.PAT,
		// Migrated to IAM
		p.AdminUser,
		p.Role,
		p.Captcha,
		p.AuditH,
		p.Org,
		p.OrgMember,
		p.Team,
		p.TeamMember,
		// App Handlers
		p.Setting,
		p.UserSetting,
		p.TaskHandler,
		p.Health,
		p.Cache,
		p.Overview,
		// CRM Handlers
		p.Company,
		p.Contact,
		p.Lead,
		p.Opportunity,
	)

	// Create MiddlewareInjector with all dependencies
	injector := NewMiddlewareInjector(RouterDepsParams{
		Config:             p.Config,
		RedisClient:        p.RedisClient,
		JWTManager:         p.JWTManager,
		PATService:         p.PATService,
		PermissionCache:    p.PermissionCache,
		AuditCreateHandler: p.Audit.CreateLog,
		MemberRepos:        p.MemberRepos,
		TeamRepos:          p.TeamRepos,
		TeamMemberRepos:    p.TeamMemberRepos,
	})

	// Register routes to engine with middleware injection
	RegisterRoutes(engine, allRoutes, injector)

	return engine
}
