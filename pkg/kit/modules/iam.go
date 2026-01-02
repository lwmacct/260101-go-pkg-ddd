package modules

import (
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/iam/application/auth"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/iam/application/pat"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/iam/application/role"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/iam/application/twofa"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/iam/application/user"
)

// IAMFacade 提供 IAM 模块的稳定访问接口。
//
// IAM (Identity and Access Management) 模块负责身份认证、授权和用户管理。
// 包含五个子模块：Auth、User、Role、PAT、TwoFA。
type IAMFacade struct {
	Auth  *AuthUseCases
	User  *UserUseCases
	Role  *RoleUseCases
	PAT   *PATUseCases
	TwoFA *TwoFAUseCases
}

// AuthUseCases 认证相关用例处理器。
//
// 提供登录、注册、令牌刷新等认证操作。
type AuthUseCases struct {
	Login        *auth.LoginHandler
	Login2FA     *auth.Login2FAHandler
	Register     *auth.RegisterHandler
	RefreshToken *auth.RefreshTokenHandler
}

// UserUseCases 用户管理用例处理器。
//
// 提供用户的 CRUD 操作、角色分配、密码管理等功能。
type UserUseCases struct {
	Create         *user.CreateHandler
	Update         *user.UpdateHandler
	Delete         *user.DeleteHandler
	Get            *user.GetHandler
	List           *user.ListHandler
	AssignRoles    *user.AssignRolesHandler
	ChangePassword *user.ChangePasswordHandler
	BatchCreate    *user.BatchCreateHandler
}

// RoleUseCases 角色管理用例处理器。
//
// 提供角色的 CRUD 操作和权限分配功能。
type RoleUseCases struct {
	Create         *role.CreateHandler
	Update         *role.UpdateHandler
	Delete         *role.DeleteHandler
	Get            *role.GetHandler
	List           *role.ListHandler
	SetPermissions *role.SetPermissionsHandler
}

// PATUseCases 个人访问令牌（Personal Access Token）用例处理器。
//
// 提供 PAT 的创建、删除、启用/禁用、查询等功能。
type PATUseCases struct {
	Create  *pat.CreateHandler
	Delete  *pat.DeleteHandler
	Disable *pat.DisableHandler
	Enable  *pat.EnableHandler
	Get     *pat.GetHandler
	List    *pat.ListHandler
}

// TwoFAUseCases 双因素认证（Two-Factor Authentication）用例处理器。
//
// 提供 2FA 的设置、验证、启用/禁用、状态查询等功能。
type TwoFAUseCases struct {
	Setup        *twofa.SetupHandler
	VerifyEnable *twofa.VerifyEnableHandler
	Disable      *twofa.DisableHandler
	GetStatus    *twofa.GetStatusHandler
}
