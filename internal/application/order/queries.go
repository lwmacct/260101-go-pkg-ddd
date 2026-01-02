package order

// GetQuery 获取订单查询。
type GetQuery struct {
	ID uint
}

// ListQuery 订单列表查询。
type ListQuery struct {
	UserID uint // 0 表示查询所有（管理员）
	Offset int
	Limit  int
}
