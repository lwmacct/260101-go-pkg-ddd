package handler

import (
	"go.uber.org/fx"

	appapplication "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/app/application"
	iamapplication "github.com/lwmacct/260101-go-pkg-ddd/pkg/modules/iam/application"
)

// HandlersResult 使用 fx.Out 批量返回 IAM 模块的所有 HTTP 处理器。
type HandlersResult struct {
	fx.Out

	Auth        *AuthHandler
	UserProfile *UserProfileHandler
	Role        *RoleHandler
	PAT         *PATHandler
	TwoFA       *TwoFAHandler
	UserOrg     *UserOrgHandler
}

// HandlerModule 提供 IAM 模块的所有 HTTP 处理器。
var HandlerModule = fx.Module("iam.handler",
	fx.Provide(newAllHandlers),
)

// handlersParams 聚合创建 Handler 所需的依赖。
type handlersParams struct {
	fx.In

	// IAM 模块用例
	Auth  *iamapplication.AuthUseCases
	User  *iamapplication.UserUseCases
	Role  *iamapplication.RoleUseCases
	PAT   *iamapplication.PATUseCases
	TwoFA *iamapplication.TwoFAUseCases

	// App 模块用例（跨模块依赖）
	Organization *appapplication.OrganizationUseCases
}

func newAllHandlers(p handlersParams) HandlersResult {
	return HandlersResult{
		Auth: NewAuthHandler(
			p.Auth.Login,
			p.Auth.Login2FA,
			p.Auth.Register,
			p.Auth.RefreshToken,
		),
		UserProfile: NewUserProfileHandler(
			p.User.Get,
			p.User.Update,
			p.User.ChangePassword,
			p.User.Delete,
		),
		Role: NewRoleHandler(
			p.Role.Create,
			p.Role.Update,
			p.Role.Delete,
			p.Role.SetPermissions,
			p.Role.Get,
			p.Role.List,
		),
		PAT: NewPATHandler(
			p.PAT.Create,
			p.PAT.Delete,
			p.PAT.Disable,
			p.PAT.Enable,
			p.PAT.Get,
			p.PAT.List,
		),
		TwoFA: NewTwoFAHandler(
			p.TwoFA.Setup,
			p.TwoFA.VerifyEnable,
			p.TwoFA.Disable,
			p.TwoFA.GetStatus,
		),
		UserOrg: NewUserOrgHandler(
			p.Organization.UserOrgs,
			p.Organization.UserTeams,
		),
	}
}
