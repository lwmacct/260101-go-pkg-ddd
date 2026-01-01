package http

import (
	"github.com/gin-gonic/gin"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/domain/permission"
	ginpermission "github.com/lwmacct/260101-go-pkg-ddd/starter/gin/permission"
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
		{ginpermission.PublicAuthRegister, deps.AuthHandler.Register},
		{ginpermission.PublicAuthLogin, deps.AuthHandler.Login},
		{ginpermission.PublicAuthLogin2FA, deps.AuthHandler.Login2FA},
		{ginpermission.PublicAuthRefresh, deps.AuthHandler.RefreshToken},
		{ginpermission.PublicAuthCaptcha, deps.CaptchaHandler.GetCaptcha},

		// ==================== Self 域 - 2FA ====================
		{ginpermission.Self2FASetup, deps.TwoFAHandler.Setup},
		{ginpermission.Self2FAVerify, deps.TwoFAHandler.VerifyAndEnable},
		{ginpermission.Self2FADisable, deps.TwoFAHandler.Disable},
		{ginpermission.Self2FAStatus, deps.TwoFAHandler.GetStatus},

		// ==================== Admin 域 - 用户管理 ====================
		{ginpermission.AdminUsersCreate, deps.AdminUserHandler.CreateUser},
		{ginpermission.AdminUsersBatchCreate, deps.AdminUserHandler.BatchCreateUsers},
		{ginpermission.AdminUsersList, deps.AdminUserHandler.ListUsers},
		{ginpermission.AdminUsersGet, deps.AdminUserHandler.GetUser},
		{ginpermission.AdminUsersUpdate, deps.AdminUserHandler.UpdateUser},
		{ginpermission.AdminUsersDelete, deps.AdminUserHandler.DeleteUser},
		{ginpermission.AdminUsersAssignRoles, deps.AdminUserHandler.AssignRoles},

		// ==================== Admin 域 - 角色管理 ====================
		{ginpermission.AdminRolesCreate, deps.RoleHandler.CreateRole},
		{ginpermission.AdminRolesList, deps.RoleHandler.ListRoles},
		{ginpermission.AdminRolesGet, deps.RoleHandler.GetRole},
		{ginpermission.AdminRolesUpdate, deps.RoleHandler.UpdateRole},
		{ginpermission.AdminRolesDelete, deps.RoleHandler.DeleteRole},
		{ginpermission.AdminRolesSetPermissions, deps.RoleHandler.SetPermissions},

		// ==================== Admin 域 - 操作列表 ====================
		{ginpermission.AdminOperationsList, deps.OperationHandler.ListOperations},

		// ==================== Admin 域 - 审计日志 ====================
		// 注意：actions 路由必须在 :id 路由之前
		{ginpermission.AdminAuditActions, deps.AuditHandler.GetActions},
		{ginpermission.AdminAuditList, deps.AuditHandler.ListLogs},
		{ginpermission.AdminAuditGet, deps.AuditHandler.GetLog},

		// ==================== Admin 域 - 系统概览 ====================
		{ginpermission.AdminOverviewStats, deps.OverviewHandler.GetStats},

		// ==================== Admin 域 - 配置分类（必须在 :key 之前） ====================
		{ginpermission.AdminSettingCategoriesList, deps.SettingHandler.GetCategories},
		{ginpermission.AdminSettingCategoriesGet, deps.SettingHandler.GetCategory},
		{ginpermission.AdminSettingCategoriesCreate, deps.SettingHandler.CreateCategory},
		{ginpermission.AdminSettingCategoriesUpdate, deps.SettingHandler.UpdateCategory},
		{ginpermission.AdminSettingCategoriesDelete, deps.SettingHandler.DeleteCategory},

		// ==================== Admin 域 - 系统配置 ====================
		// 注意：batch 路由必须在 :key 路由之前
		{ginpermission.AdminSettingsBatchUpdate, deps.SettingHandler.BatchUpdateSettings},
		{ginpermission.AdminSettingsCreate, deps.SettingHandler.CreateSetting},
		{ginpermission.AdminSettingsList, deps.SettingHandler.GetSettings},
		{ginpermission.AdminSettingsGet, deps.SettingHandler.GetSetting},
		{ginpermission.AdminSettingsUpdate, deps.SettingHandler.UpdateSetting},
		{ginpermission.AdminSettingsDelete, deps.SettingHandler.DeleteSetting},

		// ==================== Admin 域 - 缓存管理 ====================
		{ginpermission.AdminCacheInfo, deps.CacheHandler.Info},
		{ginpermission.AdminCacheScanKeys, deps.CacheHandler.ScanKeys},
		{ginpermission.AdminCacheGetKey, deps.CacheHandler.GetKey},
		{ginpermission.AdminCacheDeleteKey, deps.CacheHandler.DeleteKey},
		{ginpermission.AdminCacheDeletePattern, deps.CacheHandler.DeleteByPattern},

		// ==================== Self 域 - 个人资料 ====================
		{ginpermission.SelfProfileGet, deps.UserProfileHandler.GetProfile},
		{ginpermission.SelfProfileUpdate, deps.UserProfileHandler.UpdateProfile},
		{ginpermission.SelfPasswordUpdate, deps.UserProfileHandler.ChangePassword},
		{ginpermission.SelfAccountDelete, deps.UserProfileHandler.DeleteAccount},

		// ==================== Self 域 - 访问令牌 ====================
		// 注意：scopes 必须在 :id 路由之前
		{ginpermission.SelfTokensCreate, deps.PATHandler.CreateToken},
		{ginpermission.SelfTokensList, deps.PATHandler.ListTokens},
		{ginpermission.SelfTokensScopes, deps.PATHandler.ListScopes},
		{ginpermission.SelfTokensGet, deps.PATHandler.GetToken},
		{ginpermission.SelfTokensDelete, deps.PATHandler.DeleteToken},
		{ginpermission.SelfTokensDisable, deps.PATHandler.DisableToken},
		{ginpermission.SelfTokensEnable, deps.PATHandler.EnableToken},

		// ==================== Self 域 - 用户配置 ====================
		// 注意：categories 和 batch 路由必须在 :key 路由之前
		{ginpermission.SelfSettingsCategoriesList, deps.UserSettingHandler.ListUserSettingCategories},
		{ginpermission.SelfSettingsBatchSet, deps.UserSettingHandler.BatchSetUserSettings},
		{ginpermission.SelfSettingsList, deps.UserSettingHandler.GetUserSettings},
		{ginpermission.SelfSettingsGet, deps.UserSettingHandler.GetUserSetting},
		{ginpermission.SelfSettingsSet, deps.UserSettingHandler.SetUserSetting},
		{ginpermission.SelfSettingsReset, deps.UserSettingHandler.ResetUserSetting},

		// ==================== Self 域 - 用户组织/团队 ====================
		{ginpermission.SelfOrgsList, deps.UserOrgHandler.ListMyOrganizations},
		{ginpermission.SelfTeamsList, deps.UserOrgHandler.ListMyTeams},

		// ==================== Admin 域 - 组织管理 ====================
		{ginpermission.AdminOrgsCreate, deps.OrgHandler.Create},
		{ginpermission.AdminOrgsList, deps.OrgHandler.List},
		{ginpermission.AdminOrgsGet, deps.OrgHandler.Get},
		{ginpermission.AdminOrgsUpdate, deps.OrgHandler.Update},
		{ginpermission.AdminOrgsDelete, deps.OrgHandler.Delete},

		// ==================== Org 域 - 组织成员管理 ====================
		{ginpermission.OrgMembersList, deps.OrgMemberHandler.List},
		{ginpermission.OrgMembersAdd, deps.OrgMemberHandler.Add},
		{ginpermission.OrgMembersRemove, deps.OrgMemberHandler.Remove},
		{ginpermission.OrgMembersUpdateRole, deps.OrgMemberHandler.UpdateRole},

		// ==================== Org 域 - 团队管理 ====================
		{ginpermission.OrgTeamsCreate, deps.TeamHandler.Create},
		{ginpermission.OrgTeamsList, deps.TeamHandler.List},
		{ginpermission.OrgTeamsGet, deps.TeamHandler.Get},
		{ginpermission.OrgTeamsUpdate, deps.TeamHandler.Update},
		{ginpermission.OrgTeamsDelete, deps.TeamHandler.Delete},

		// ==================== Org 域 - 团队成员管理 ====================
		{ginpermission.OrgTeamMembersList, deps.TeamMemberHandler.List},
		{ginpermission.OrgTeamMembersAdd, deps.TeamMemberHandler.Add},
		{ginpermission.OrgTeamMembersRemove, deps.TeamMemberHandler.Remove},

		// ==================== Admin 域 - 产品管理 ====================
		{ginpermission.AdminProductsCreate, deps.ProductHandler.Create},
		{ginpermission.AdminProductsList, deps.ProductHandler.List},
		{ginpermission.AdminProductsGet, deps.ProductHandler.Get},
		{ginpermission.AdminProductsUpdate, deps.ProductHandler.Update},
		{ginpermission.AdminProductsDelete, deps.ProductHandler.Delete},

		// ==================== Org 域 - 团队任务管理 ====================
		{ginpermission.OrgTasksCreate, deps.TaskHandler.Create},
		{ginpermission.OrgTasksList, deps.TaskHandler.List},
		{ginpermission.OrgTasksGet, deps.TaskHandler.Get},
		{ginpermission.OrgTasksUpdate, deps.TaskHandler.Update},
		{ginpermission.OrgTasksDelete, deps.TaskHandler.Delete},
	}
}
