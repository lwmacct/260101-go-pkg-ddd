package invoice

// GetQuery 获取发票查询。
type GetQuery struct {
	ID uint
}

// ListQuery 发票列表查询。
type ListQuery struct {
	UserID  uint // 0 表示查询所有（管理员）
	OrderID uint // 按订单筛选
	Offset  int
	Limit   int
}
