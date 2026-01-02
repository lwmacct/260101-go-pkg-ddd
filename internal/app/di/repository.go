package di

import (
	"go.uber.org/fx"
	"gorm.io/gorm"

	"github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/core/application/setting"
	corepersistence "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/core/infrastructure/persistence"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/iam/application/user"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/shared/captcha"

	crmpersistence "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/crm/infrastructure/persistence"
	iampersistence "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/iam/infrastructure/persistence"
)

// RepositoryModule 提供所有仓储实现。
//
// 适当的仓储已装饰缓存层：
//   - User: 缓存查询（GetByIDWithRoles）
//   - Setting: 缓存查询 + 命令，支持多级失效
//   - UserSetting: 缓存查询 + 命令，支持三级失效
var RepositoryModule = fx.Module("repository",
	fx.Provide(
		// 直接使用 persistence 构造函数（无需包装）
		corepersistence.NewAuditRepositories,
		iampersistence.NewRoleRepositories,
		iampersistence.NewPATRepositories,
		iampersistence.NewTwoFARepositories,

		// 组织相关仓储
		corepersistence.NewOrganizationRepositories,
		corepersistence.NewTeamRepositories,
		corepersistence.NewOrgMemberRepositories,
		corepersistence.NewTeamMemberRepositories,

		// 任务仓储
		corepersistence.NewTaskRepositories,

		// CRM 仓储
		crmpersistence.NewContactRepositories,
		crmpersistence.NewCompanyRepositories,
		crmpersistence.NewLeadRepositories,
		crmpersistence.NewOpportunityRepositories,

		// 带缓存装饰的仓储
		newUserRepositoriesWithCache,
		newSettingRepositoriesWithCache,
		newUserSettingRepositoriesWithCache,

		// 特殊仓储
		// newCaptchaRepository, // TODO: implement captcha repository
		// infrastats.NewQueryRepository,
	),
)

// --- 带缓存装饰的仓储构造函数 ---

func newUserRepositoriesWithCache(
	db *gorm.DB,
	userWithRolesCache user.UserWithRolesCacheService,
) iampersistence.UserRepositories {
	rawRepos := iampersistence.NewUserRepositories(db)
	cachedQuery := iampersistence.NewCachedUserQueryRepository(rawRepos.Query, userWithRolesCache)
	cachedCommand := iampersistence.NewCachedUserCommandRepository(rawRepos.Command, userWithRolesCache)
	return iampersistence.UserRepositories{
		Command: cachedCommand,
		Query:   cachedQuery,
	}
}

// newSettingRepositoriesWithCache 创建带缓存装饰的 Setting 仓储。
//
// 简化设计：
//   - Query 直接使用原始仓储，不再缓存（由 Application 层 Settings 缓存覆盖）
//   - Command 装饰器只负责写操作后失效下游缓存（Settings + UserSetting）
//   - CategoryQuery 使用 SettingsCacheService 的 Category 缓存方法
//   - CategoryCommand 直接使用原始仓储，缓存失效在 Handler 层处理
func newSettingRepositoriesWithCache(
	db *gorm.DB,
	userSettingCache setting.UserSettingCacheService,
	settingsCache setting.SettingsCacheService,
) corepersistence.SettingRepositories {
	rawRepos := corepersistence.NewSettingRepositories(db)

	// 查询直接使用原始仓储，不再缓存
	// 写操作装饰器：失效 Settings + UserSetting 缓存
	wrappedCommand := corepersistence.NewSettingCommandWithCacheInvalidation(
		rawRepos.Command,
		userSettingCache,
		settingsCache,
	)

	// Category 查询使用 SettingsCacheService（合并后的 Application 层缓存）
	cachedCategoryQuery := corepersistence.NewCachedSettingCategoryQueryRepository(
		rawRepos.CategoryQuery,
		settingsCache,
	)

	// Category 命令直接使用原始仓储，缓存失效在 Handler 层统一处理
	return corepersistence.SettingRepositories{
		Command:         wrappedCommand,
		Query:           rawRepos.Query,
		CategoryQuery:   cachedCategoryQuery,
		CategoryCommand: rawRepos.CategoryCommand,
	}
}

// userSettingRepositoriesParams 聚合 UserSetting 仓储所需的缓存服务。
type userSettingRepositoriesParams struct {
	fx.In

	DB                    *gorm.DB
	UserSettingQueryCache setting.UserSettingQueryCacheService
	UserSettingCache      setting.UserSettingCacheService
	SettingsCache         setting.SettingsCacheService
}

func newUserSettingRepositoriesWithCache(p userSettingRepositoriesParams) iampersistence.UserSettingRepositories {
	rawRepos := iampersistence.NewUserSettingRepositories(p.DB)

	cachedQuery := iampersistence.NewCachedUserSettingQueryRepository(
		rawRepos.Query,
		p.UserSettingQueryCache,
	)
	cachedCommand := iampersistence.NewCachedUserSettingCommandRepository(
		rawRepos.Command,
		p.UserSettingQueryCache,
		p.UserSettingCache,
		p.SettingsCache,
	)

	return iampersistence.UserSettingRepositories{
		Command: cachedCommand,
		Query:   cachedQuery,
	}
}

// --- 特殊仓储 ---

// CaptchaRepositoryResult 从单个仓储提供 Command 和 Query 两个接口。
type CaptchaRepositoryResult struct {
	fx.Out

	Command captcha.CommandRepository
	Query   captcha.QueryRepository
}

// TODO: implement captcha repository
// func newCaptchaRepository() CaptchaRepositoryResult {
// 	repo := infracaptcha.NewRepository()
// 	return CaptchaRepositoryResult{
// 		Command: repo,
// 		Query:   repo,
// 	}
// }
