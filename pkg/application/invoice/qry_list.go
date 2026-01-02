package invoice

import (
	"context"

	"github.com/lwmacct/260101-go-pkg-ddd/pkg/domain/invoice"
)

// ListHandler 发票列表处理器。
type ListHandler struct {
	queryRepo invoice.QueryRepository
}

// NewListHandler 创建 ListHandler。
func NewListHandler(queryRepo invoice.QueryRepository) *ListHandler {
	return &ListHandler{queryRepo: queryRepo}
}

// Handle 处理发票列表查询。
func (h *ListHandler) Handle(ctx context.Context, query ListQuery) (*ListResultDTO, error) {
	var (
		items []*invoice.Invoice
		total int64
		err   error
	)

	// 按条件筛选
	switch {
	case query.OrderID > 0:
		// 按订单筛选
		items, err = h.queryRepo.ListByOrder(ctx, query.OrderID)
		if err != nil {
			return nil, err
		}
		total = int64(len(items))
	case query.UserID > 0:
		// 按用户筛选
		items, err = h.queryRepo.ListByUser(ctx, query.UserID, query.Offset, query.Limit)
		if err != nil {
			return nil, err
		}
		total, err = h.queryRepo.CountByUser(ctx, query.UserID)
		if err != nil {
			return nil, err
		}
	default:
		// 管理员查询所有
		items, err = h.queryRepo.List(ctx, query.Offset, query.Limit)
		if err != nil {
			return nil, err
		}
		total, err = h.queryRepo.Count(ctx)
		if err != nil {
			return nil, err
		}
	}

	return &ListResultDTO{
		Items: ToInvoiceDTOs(items),
		Total: total,
	}, nil
}
