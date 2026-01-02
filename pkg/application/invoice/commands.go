package invoice

import "time"

// CreateCommand 创建发票命令。
type CreateCommand struct {
	OrderID uint
	UserID  uint
	Amount  float64
	DueDate time.Time
	Remark  string
}

// PayCommand 支付发票命令。
type PayCommand struct {
	ID     uint
	Amount float64 // 支付金额（支持部分支付）
}

// CancelCommand 取消发票命令。
type CancelCommand struct {
	ID uint
}

// RefundCommand 退款命令。
type RefundCommand struct {
	ID uint
}

// UpdateStatusCommand 更新发票状态命令（内部使用）。
type UpdateStatusCommand struct {
	ID     uint
	Status string
}
