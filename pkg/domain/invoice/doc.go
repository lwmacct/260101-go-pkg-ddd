// Package invoice 定义发票领域模型和仓储接口。
//
// 本包定义了：
//   - [Invoice]: 发票实体（支持部分支付）
//   - [CommandRepository]: 写仓储接口
//   - [QueryRepository]: 读仓储接口
//
// 业务规则：
//   - 一个订单可对应多个发票（1:N 关系，支持分期付款）
//   - 发票支持部分支付，PaidAmount 累计直到达到 Amount
//   - 状态流转：draft → pending → partial → paid / canceled / refunded
//
// 依赖倒置：本包仅定义接口，实现位于 infrastructure/persistence。
package invoice
