// Package permission 定义 HTTP 层的操作常量。
//
// 这些常量是 HTTP 特定的，用于路由注册和权限检查。
// 格式：{scope}:{type}:{action}
//
// Scope 划分：
//   - public: 公开域（无需认证）
//   - self:   用户自服务域（当前用户权限）
//   - admin:  系统管理域（需管理员权限）
//   - org:    组织域（多租户场景）
package permission

import "github.com/lwmacct/260101-go-pkg-ddd/pkg/domain/permission"

// Public 域（公开接口）
const (
	PublicAuthRegister permission.Operation = "public:auth:register"
	PublicAuthLogin    permission.Operation = "public:auth:login"
	PublicAuthLogin2FA permission.Operation = "public:auth:login2fa"
	PublicAuthRefresh  permission.Operation = "public:auth:refresh"
	PublicAuthCaptcha  permission.Operation = "public:auth:captcha"
)

// Self 域 - 2FA（需认证）
const (
	Self2FASetup   permission.Operation = "self:2fa:setup"
	Self2FAVerify  permission.Operation = "self:2fa:verify"
	Self2FADisable permission.Operation = "self:2fa:disable"
	Self2FAStatus  permission.Operation = "self:2fa:status"
)

// Admin 域 - 用户管理
const (
	AdminUsersCreate      permission.Operation = "admin:users:create"
	AdminUsersBatchCreate permission.Operation = "admin:users:batch-create"
	AdminUsersList        permission.Operation = "admin:users:list"
	AdminUsersGet         permission.Operation = "admin:users:get"
	AdminUsersUpdate      permission.Operation = "admin:users:update"
	AdminUsersDelete      permission.Operation = "admin:users:delete"
	AdminUsersAssignRoles permission.Operation = "admin:users:assign-roles"
)

// Admin 域 - 角色管理
const (
	AdminRolesCreate         permission.Operation = "admin:roles:create"
	AdminRolesList           permission.Operation = "admin:roles:list"
	AdminRolesGet            permission.Operation = "admin:roles:get"
	AdminRolesUpdate         permission.Operation = "admin:roles:update"
	AdminRolesDelete         permission.Operation = "admin:roles:delete"
	AdminRolesSetPermissions permission.Operation = "admin:roles:set-permissions"
)

// Admin 域 - 操作列表（供前端权限配置使用）
const (
	AdminOperationsList permission.Operation = "admin:operations:list"
)

// Admin 域 - 审计日志
const (
	AdminAuditList    permission.Operation = "admin:audit:list"
	AdminAuditGet     permission.Operation = "admin:audit:get"
	AdminAuditActions permission.Operation = "admin:audit:actions"
)

// Admin 域 - 系统概览
const (
	AdminOverviewStats permission.Operation = "admin:overview:stats"
)

// Admin 域 - 系统配置
const (
	AdminSettingsCreate      permission.Operation = "admin:settings:create"
	AdminSettingsList        permission.Operation = "admin:settings:list"
	AdminSettingsGet         permission.Operation = "admin:settings:get"
	AdminSettingsUpdate      permission.Operation = "admin:settings:update"
	AdminSettingsDelete      permission.Operation = "admin:settings:delete"
	AdminSettingsBatchUpdate permission.Operation = "admin:settings:batch-update"
)

// Admin 域 - 配置分类
const (
	AdminSettingCategoriesList   permission.Operation = "admin:settings-categories:list"
	AdminSettingCategoriesGet    permission.Operation = "admin:settings-categories:get"
	AdminSettingCategoriesCreate permission.Operation = "admin:settings-categories:create"
	AdminSettingCategoriesUpdate permission.Operation = "admin:settings-categories:update"
	AdminSettingCategoriesDelete permission.Operation = "admin:settings-categories:delete"
)

// Admin 域 - 缓存管理
const (
	AdminCacheInfo          permission.Operation = "admin:cache:info"
	AdminCacheScanKeys      permission.Operation = "admin:cache:scan-keys"
	AdminCacheGetKey        permission.Operation = "admin:cache:get-key"
	AdminCacheDeleteKey     permission.Operation = "admin:cache:delete-key"
	AdminCacheDeletePattern permission.Operation = "admin:cache:delete-pattern"
)

// Self 域 - 个人资料
const (
	SelfProfileGet     permission.Operation = "self:profile:get"
	SelfProfileUpdate  permission.Operation = "self:profile:update"
	SelfPasswordUpdate permission.Operation = "self:password:update"
	SelfAccountDelete  permission.Operation = "self:account:delete"
)

// Self 域 - 访问令牌
//
//nolint:gosec // G101: 这些是操作标识符，非硬编码凭证
const (
	SelfTokensCreate  permission.Operation = "self:tokens:create"
	SelfTokensList    permission.Operation = "self:tokens:list"
	SelfTokensGet     permission.Operation = "self:tokens:get"
	SelfTokensDelete  permission.Operation = "self:tokens:delete"
	SelfTokensDisable permission.Operation = "self:tokens:disable"
	SelfTokensEnable  permission.Operation = "self:tokens:enable"
	SelfTokensScopes  permission.Operation = "self:tokens:scopes"
)

// Self 域 - 用户配置
const (
	SelfSettingsCategoriesList permission.Operation = "self:settings-categories:list"
	SelfSettingsList           permission.Operation = "self:settings:list"
	SelfSettingsGet            permission.Operation = "self:settings:get"
	SelfSettingsSet            permission.Operation = "self:settings:set"
	SelfSettingsReset          permission.Operation = "self:settings:reset"
	SelfSettingsBatchSet       permission.Operation = "self:settings:batch-set"
)

// Self 域 - 用户组织/团队
const (
	SelfOrgsList  permission.Operation = "self:orgs:list"
	SelfTeamsList permission.Operation = "self:teams:list"
)

// Admin 域 - 组织管理
const (
	AdminOrgsCreate permission.Operation = "admin:orgs:create"
	AdminOrgsList   permission.Operation = "admin:orgs:list"
	AdminOrgsGet    permission.Operation = "admin:orgs:get"
	AdminOrgsUpdate permission.Operation = "admin:orgs:update"
	AdminOrgsDelete permission.Operation = "admin:orgs:delete"
)

// Org 域 - 组织成员管理
const (
	OrgMembersList       permission.Operation = "org:members:list"
	OrgMembersAdd        permission.Operation = "org:members:add"
	OrgMembersRemove     permission.Operation = "org:members:remove"
	OrgMembersUpdateRole permission.Operation = "org:members:update-role"
)

// Org 域 - 团队管理
const (
	OrgTeamsCreate permission.Operation = "org:teams:create"
	OrgTeamsList   permission.Operation = "org:teams:list"
	OrgTeamsGet    permission.Operation = "org:teams:get"
	OrgTeamsUpdate permission.Operation = "org:teams:update"
	OrgTeamsDelete permission.Operation = "org:teams:delete"
)

// Org 域 - 团队成员管理
const (
	OrgTeamMembersList   permission.Operation = "org:team-members:list"
	OrgTeamMembersAdd    permission.Operation = "org:team-members:add"
	OrgTeamMembersRemove permission.Operation = "org:team-members:remove"
)

// Admin 域 - 产品管理
const (
	AdminProductsCreate permission.Operation = "admin:products:create"
	AdminProductsList   permission.Operation = "admin:products:list"
	AdminProductsGet    permission.Operation = "admin:products:get"
	AdminProductsUpdate permission.Operation = "admin:products:update"
	AdminProductsDelete permission.Operation = "admin:products:delete"
)

// Org 域 - 团队任务管理
const (
	OrgTasksCreate permission.Operation = "org:tasks:create"
	OrgTasksList   permission.Operation = "org:tasks:list"
	OrgTasksGet    permission.Operation = "org:tasks:get"
	OrgTasksUpdate permission.Operation = "org:tasks:update"
	OrgTasksDelete permission.Operation = "org:tasks:delete"
)
