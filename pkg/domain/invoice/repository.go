package invoice

import "context"

// CommandRepository 发票写仓储接口。
type CommandRepository interface {
	// Create 创建发票，成功后回写 ID 到实体。
	Create(ctx context.Context, invoice *Invoice) error
	// Update 更新发票。
	Update(ctx context.Context, invoice *Invoice) error
	// Delete 删除发票。
	Delete(ctx context.Context, id uint) error
}

// QueryRepository 发票读仓储接口。
type QueryRepository interface {
	// GetByID 根据 ID 获取发票。
	GetByID(ctx context.Context, id uint) (*Invoice, error)
	// GetByInvoiceNo 根据发票号获取发票。
	GetByInvoiceNo(ctx context.Context, invoiceNo string) (*Invoice, error)
	// ListByOrder 获取订单下的所有发票。
	ListByOrder(ctx context.Context, orderID uint) ([]*Invoice, error)
	// ListByUser 分页获取用户发票列表。
	ListByUser(ctx context.Context, userID uint, offset, limit int) ([]*Invoice, error)
	// CountByUser 获取用户发票总数。
	CountByUser(ctx context.Context, userID uint) (int64, error)
	// List 分页获取发票列表（管理员）。
	List(ctx context.Context, offset, limit int) ([]*Invoice, error)
	// Count 获取发票总数。
	Count(ctx context.Context) (int64, error)
}
