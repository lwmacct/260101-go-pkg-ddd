// Package order 定义订单领域模型和仓储接口。
//
// 本包定义了：
//   - [Order]: 订单实体
//   - [CommandRepository]: 写仓储接口
//   - [QueryRepository]: 读仓储接口
//
// 依赖倒置：本包仅定义接口，实现位于 infrastructure/persistence。
package order
