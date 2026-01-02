package modules

import (
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/app/application/audit"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/app/application/captcha"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/app/application/org"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/app/application/setting"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/app/application/stats"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/app/application/task"
)

// AppFacade 提供 App 核心模块的稳定访问接口。
//
// App 模块负责核心治理功能，包括组织管理、配置管理、任务管理、
// 审计日志、验证码和统计数据等。
//
// 注意：Cache 和 Health 是 Transport 层的技术性接口，
// 不属于核心业务用例，因此不在此 Facade 中暴露。
type AppFacade struct {
	Organization *OrganizationUseCases
	Setting      *SettingUseCases
	UserSetting  *UserSettingUseCases
	Task         *TaskUseCases
	Audit        *AuditUseCases
	Captcha      *CaptchaUseCases
	Stats        *StatsUseCases
}

// OrganizationUseCases 组织管理用例处理器。
//
// 提供组织、成员、团队、团队成员的 CRUD 操作，
// 以及用户视角的组织和团队查询功能。
type OrganizationUseCases struct {
	// Organization CRUD
	Create *org.CreateHandler
	Update *org.UpdateHandler
	Delete *org.DeleteHandler
	Get    *org.GetHandler
	List   *org.ListHandler

	// Member Management
	MemberAdd        *org.MemberAddHandler
	MemberRemove     *org.MemberRemoveHandler
	MemberUpdateRole *org.MemberUpdateRoleHandler
	MemberList       *org.MemberListHandler

	// Team CRUD
	TeamCreate *org.TeamCreateHandler
	TeamUpdate *org.TeamUpdateHandler
	TeamDelete *org.TeamDeleteHandler
	TeamGet    *org.TeamGetHandler
	TeamList   *org.TeamListHandler

	// Team Member Management
	TeamMemberAdd    *org.TeamMemberAddHandler
	TeamMemberRemove *org.TeamMemberRemoveHandler
	TeamMemberList   *org.TeamMemberListHandler

	// User View
	UserOrgs  *org.UserOrgsHandler
	UserTeams *org.UserTeamsHandler
}

// SettingUseCases 系统配置管理用例处理器。
//
// 提供系统级配置项和配置分类的 CRUD 操作。
type SettingUseCases struct {
	// Setting CRUD
	Create       *setting.CreateHandler
	Update       *setting.UpdateHandler
	Delete       *setting.DeleteHandler
	BatchUpdate  *setting.BatchUpdateHandler
	Get          *setting.GetHandler
	List         *setting.ListHandler
	ListSettings *setting.ListSettingsHandler

	// Category CRUD
	CreateCategory *setting.CreateCategoryHandler
	UpdateCategory *setting.UpdateCategoryHandler
	DeleteCategory *setting.DeleteCategoryHandler
	GetCategory    *setting.GetCategoryHandler
	ListCategories *setting.ListCategoriesHandler
}

// UserSettingUseCases 用户个人配置管理用例处理器。
//
// 提供用户个人配置项的设置、重置、查询等操作。
type UserSettingUseCases struct {
	Set            *setting.UserSetHandler
	BatchSet       *setting.UserBatchSetHandler
	Reset          *setting.UserResetHandler
	ResetAll       *setting.UserResetAllHandler
	Get            *setting.UserGetHandler
	List           *setting.UserListHandler
	ListSettings   *setting.UserListSettingsHandler
	ListCategories *setting.UserListCategoriesHandler
}

// TaskUseCases 任务管理用例处理器。
//
// 提供任务的 CRUD 操作。
type TaskUseCases struct {
	Create *task.CreateHandler
	Update *task.UpdateHandler
	Delete *task.DeleteHandler
	Get    *task.GetHandler
	List   *task.ListHandler
}

// AuditUseCases 审计日志用例处理器。
//
// 提供审计日志的创建和查询功能。
type AuditUseCases struct {
	CreateLog *audit.CreateHandler
	Get       *audit.GetHandler
	List      *audit.ListHandler
}

// CaptchaUseCases 验证码用例处理器。
//
// 提供图形验证码的生成功能。
type CaptchaUseCases struct {
	Generate *captcha.GenerateHandler
}

// StatsUseCases 统计数据用例处理器。
//
// 提供跨模块的聚合统计数据查询功能。
type StatsUseCases struct {
	GetStats *stats.GetStatsHandler
}
