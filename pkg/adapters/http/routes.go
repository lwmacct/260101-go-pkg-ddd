package http

import (
	"github.com/gin-gonic/gin"
	permissionHttp "github.com/lwmacct/260101-go-pkg-ddd/pkg/adapters/http/permission"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/domain/permission"
)

// RouteBinding 路由绑定，将 Operation 与 Handler 关联
type RouteBinding struct {
	Op      permission.Operation
	Handler gin.HandlerFunc
}

// AllRouteBindings 返回所有路由绑定
// 绑定顺序决定路由注册顺序，对于同路径前缀的路由需注意顺序
func (deps *RouterDependencies) AllRouteBindings() []RouteBinding {
	return []RouteBinding{
		// ==================== Public 域（公开） ====================
		{permissionHttp.PublicAuthRegister, deps.AuthHandler.Register},
		{permissionHttp.PublicAuthLogin, deps.AuthHandler.Login},
		{permissionHttp.PublicAuthLogin2FA, deps.AuthHandler.Login2FA},
		{permissionHttp.PublicAuthRefresh, deps.AuthHandler.RefreshToken},
		{permissionHttp.PublicAuthCaptcha, deps.CaptchaHandler.GetCaptcha},

		// ==================== Self 域 - 2FA ====================
		{permissionHttp.Self2FASetup, deps.TwoFAHandler.Setup},
		{permissionHttp.Self2FAVerify, deps.TwoFAHandler.VerifyAndEnable},
		{permissionHttp.Self2FADisable, deps.TwoFAHandler.Disable},
		{permissionHttp.Self2FAStatus, deps.TwoFAHandler.GetStatus},

		// ==================== Admin 域 - 用户管理 ====================
		{permissionHttp.AdminUsersCreate, deps.AdminUserHandler.CreateUser},
		{permissionHttp.AdminUsersBatchCreate, deps.AdminUserHandler.BatchCreateUsers},
		{permissionHttp.AdminUsersList, deps.AdminUserHandler.ListUsers},
		{permissionHttp.AdminUsersGet, deps.AdminUserHandler.GetUser},
		{permissionHttp.AdminUsersUpdate, deps.AdminUserHandler.UpdateUser},
		{permissionHttp.AdminUsersDelete, deps.AdminUserHandler.DeleteUser},
		{permissionHttp.AdminUsersAssignRoles, deps.AdminUserHandler.AssignRoles},

		// ==================== Admin 域 - 角色管理 ====================
		{permissionHttp.AdminRolesCreate, deps.RoleHandler.CreateRole},
		{permissionHttp.AdminRolesList, deps.RoleHandler.ListRoles},
		{permissionHttp.AdminRolesGet, deps.RoleHandler.GetRole},
		{permissionHttp.AdminRolesUpdate, deps.RoleHandler.UpdateRole},
		{permissionHttp.AdminRolesDelete, deps.RoleHandler.DeleteRole},
		{permissionHttp.AdminRolesSetPermissions, deps.RoleHandler.SetPermissions},

		// ==================== Admin 域 - 操作列表 ====================
		{permissionHttp.AdminOperationsList, deps.OperationHandler.ListOperations},

		// ==================== Admin 域 - 审计日志 ====================
		// 注意：actions 路由必须在 :id 路由之前
		{permissionHttp.AdminAuditActions, deps.AuditHandler.GetActions},
		{permissionHttp.AdminAuditList, deps.AuditHandler.ListLogs},
		{permissionHttp.AdminAuditGet, deps.AuditHandler.GetLog},

		// ==================== Admin 域 - 系统概览 ====================
		{permissionHttp.AdminOverviewStats, deps.OverviewHandler.GetStats},

		// ==================== Admin 域 - 配置分类（必须在 :key 之前） ====================
		{permissionHttp.AdminSettingCategoriesList, deps.SettingHandler.GetCategories},
		{permissionHttp.AdminSettingCategoriesGet, deps.SettingHandler.GetCategory},
		{permissionHttp.AdminSettingCategoriesCreate, deps.SettingHandler.CreateCategory},
		{permissionHttp.AdminSettingCategoriesUpdate, deps.SettingHandler.UpdateCategory},
		{permissionHttp.AdminSettingCategoriesDelete, deps.SettingHandler.DeleteCategory},

		// ==================== Admin 域 - 系统配置 ====================
		// 注意：batch 路由必须在 :key 路由之前
		{permissionHttp.AdminSettingsBatchUpdate, deps.SettingHandler.BatchUpdateSettings},
		{permissionHttp.AdminSettingsCreate, deps.SettingHandler.CreateSetting},
		{permissionHttp.AdminSettingsList, deps.SettingHandler.GetSettings},
		{permissionHttp.AdminSettingsGet, deps.SettingHandler.GetSetting},
		{permissionHttp.AdminSettingsUpdate, deps.SettingHandler.UpdateSetting},
		{permissionHttp.AdminSettingsDelete, deps.SettingHandler.DeleteSetting},

		// ==================== Admin 域 - 缓存管理 ====================
		{permissionHttp.AdminCacheInfo, deps.CacheHandler.Info},
		{permissionHttp.AdminCacheScanKeys, deps.CacheHandler.ScanKeys},
		{permissionHttp.AdminCacheGetKey, deps.CacheHandler.GetKey},
		{permissionHttp.AdminCacheDeleteKey, deps.CacheHandler.DeleteKey},
		{permissionHttp.AdminCacheDeletePattern, deps.CacheHandler.DeleteByPattern},

		// ==================== Self 域 - 个人资料 ====================
		{permissionHttp.SelfProfileGet, deps.UserProfileHandler.GetProfile},
		{permissionHttp.SelfProfileUpdate, deps.UserProfileHandler.UpdateProfile},
		{permissionHttp.SelfPasswordUpdate, deps.UserProfileHandler.ChangePassword},
		{permissionHttp.SelfAccountDelete, deps.UserProfileHandler.DeleteAccount},

		// ==================== Self 域 - 访问令牌 ====================
		// 注意：scopes 必须在 :id 路由之前
		{permissionHttp.SelfTokensCreate, deps.PATHandler.CreateToken},
		{permissionHttp.SelfTokensList, deps.PATHandler.ListTokens},
		{permissionHttp.SelfTokensScopes, deps.PATHandler.ListScopes},
		{permissionHttp.SelfTokensGet, deps.PATHandler.GetToken},
		{permissionHttp.SelfTokensDelete, deps.PATHandler.DeleteToken},
		{permissionHttp.SelfTokensDisable, deps.PATHandler.DisableToken},
		{permissionHttp.SelfTokensEnable, deps.PATHandler.EnableToken},

		// ==================== Self 域 - 用户配置 ====================
		// 注意：categories 和 batch 路由必须在 :key 路由之前
		{permissionHttp.SelfSettingsCategoriesList, deps.UserSettingHandler.ListUserSettingCategories},
		{permissionHttp.SelfSettingsBatchSet, deps.UserSettingHandler.BatchSetUserSettings},
		{permissionHttp.SelfSettingsList, deps.UserSettingHandler.GetUserSettings},
		{permissionHttp.SelfSettingsGet, deps.UserSettingHandler.GetUserSetting},
		{permissionHttp.SelfSettingsSet, deps.UserSettingHandler.SetUserSetting},
		{permissionHttp.SelfSettingsReset, deps.UserSettingHandler.ResetUserSetting},

		// ==================== Self 域 - 用户组织/团队 ====================
		{permissionHttp.SelfOrgsList, deps.UserOrgHandler.ListMyOrganizations},
		{permissionHttp.SelfTeamsList, deps.UserOrgHandler.ListMyTeams},

		// ==================== Admin 域 - 组织管理 ====================
		{permissionHttp.AdminOrgsCreate, deps.OrgHandler.Create},
		{permissionHttp.AdminOrgsList, deps.OrgHandler.List},
		{permissionHttp.AdminOrgsGet, deps.OrgHandler.Get},
		{permissionHttp.AdminOrgsUpdate, deps.OrgHandler.Update},
		{permissionHttp.AdminOrgsDelete, deps.OrgHandler.Delete},

		// ==================== Org 域 - 组织成员管理 ====================
		{permissionHttp.OrgMembersList, deps.OrgMemberHandler.List},
		{permissionHttp.OrgMembersAdd, deps.OrgMemberHandler.Add},
		{permissionHttp.OrgMembersRemove, deps.OrgMemberHandler.Remove},
		{permissionHttp.OrgMembersUpdateRole, deps.OrgMemberHandler.UpdateRole},

		// ==================== Org 域 - 团队管理 ====================
		{permissionHttp.OrgTeamsCreate, deps.TeamHandler.Create},
		{permissionHttp.OrgTeamsList, deps.TeamHandler.List},
		{permissionHttp.OrgTeamsGet, deps.TeamHandler.Get},
		{permissionHttp.OrgTeamsUpdate, deps.TeamHandler.Update},
		{permissionHttp.OrgTeamsDelete, deps.TeamHandler.Delete},

		// ==================== Org 域 - 团队成员管理 ====================
		{permissionHttp.OrgTeamMembersList, deps.TeamMemberHandler.List},
		{permissionHttp.OrgTeamMembersAdd, deps.TeamMemberHandler.Add},
		{permissionHttp.OrgTeamMembersRemove, deps.TeamMemberHandler.Remove},

		// ==================== Admin 域 - 产品管理 ====================
		{permissionHttp.AdminProductsCreate, deps.ProductHandler.Create},
		{permissionHttp.AdminProductsList, deps.ProductHandler.List},
		{permissionHttp.AdminProductsGet, deps.ProductHandler.Get},
		{permissionHttp.AdminProductsUpdate, deps.ProductHandler.Update},
		{permissionHttp.AdminProductsDelete, deps.ProductHandler.Delete},

		// ==================== Org 域 - 团队任务管理 ====================
		{permissionHttp.OrgTasksCreate, deps.TaskHandler.Create},
		{permissionHttp.OrgTasksList, deps.TaskHandler.List},
		{permissionHttp.OrgTasksGet, deps.TaskHandler.Get},
		{permissionHttp.OrgTasksUpdate, deps.TaskHandler.Update},
		{permissionHttp.OrgTasksDelete, deps.TaskHandler.Delete},
	}
}
