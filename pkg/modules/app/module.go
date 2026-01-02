// Package core 提供核心治理业务模块（Core Bounded Context）。
//
// 本模块包含：
//   - Domain: 组织、设置、统计、任务、审计
//   - Application: 用例处理器
//   - Infrastructure: 仓储实现
//   - Transport: HTTP 适配器
package core

import "go.uber.org/fx"

// Module 返回 Core BC 的 Fx 模块
func Module() fx.Option {
	return fx.Module("appni") // TODO: 添加 providers

}
