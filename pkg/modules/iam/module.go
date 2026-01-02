// Package iam 提供身份与访问管理业务模块（IAM Bounded Context）。
//
// 本模块包含：
//   - Domain: 用户、角色、权限、认证、双因素认证、PAT
//   - Application: 用例处理器
//   - Infrastructure: JWT、TOTP、仓储实现
//   - Transport: HTTP 适配器
package iam

import "go.uber.org/fx"

// Module 返回 IAM BC 的 Fx 模块
func Module() fx.Option {
	return fx.Module("iam") // TODO: 添加 providers
	// fx.Provide(
	//     NewRepositories,
	//     NewUseCases,
	//     NewHandlers,
	// ),

}
