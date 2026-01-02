package order

import "context"

// CommandRepository 订单写仓储接口。
type CommandRepository interface {
	// Create 创建订单，成功后回写 ID 到实体。
	Create(ctx context.Context, order *Order) error
	// Update 更新订单。
	Update(ctx context.Context, order *Order) error
	// Delete 删除订单。
	Delete(ctx context.Context, id uint) error
}

// QueryRepository 订单读仓储接口。
type QueryRepository interface {
	// GetByID 根据 ID 获取订单。
	GetByID(ctx context.Context, id uint) (*Order, error)
	// GetByOrderNo 根据订单号获取订单。
	GetByOrderNo(ctx context.Context, orderNo string) (*Order, error)
	// ListByUser 分页获取用户订单列表。
	ListByUser(ctx context.Context, userID uint, offset, limit int) ([]*Order, error)
	// CountByUser 获取用户订单总数。
	CountByUser(ctx context.Context, userID uint) (int64, error)
	// List 分页获取订单列表（管理员）。
	List(ctx context.Context, offset, limit int) ([]*Order, error)
	// Count 获取订单总数。
	Count(ctx context.Context) (int64, error)
}
