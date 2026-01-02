package order

// CreateCommand 创建订单命令。
type CreateCommand struct {
	UserID    uint
	ProductID uint
	Quantity  int
	Remark    string
}

// UpdateCommand 更新订单命令。
type UpdateCommand struct {
	ID     uint
	Remark string
}

// UpdateStatusCommand 更新订单状态命令。
type UpdateStatusCommand struct {
	ID     uint
	Status string
}

// DeleteCommand 删除订单命令。
type DeleteCommand struct {
	ID uint
}
