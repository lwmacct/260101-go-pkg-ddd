// Package crm 提供客户关系管理业务模块（CRM Bounded Context）。
//
// 本模块包含：
//   - Domain: 公司、联系人、线索、商机
//   - Application: 用例处理器
//   - Infrastructure: 仓储实现
//   - Transport: HTTP 适配器
package crm

import "go.uber.org/fx"

// Module 返回 CRM BC 的 Fx 模块
func Module() fx.Option {
	return fx.Module("crm") // TODO: 添加 providers

}
