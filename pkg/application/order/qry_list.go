package order

import (
	"context"

	"github.com/lwmacct/260101-go-pkg-ddd/pkg/domain/order"
)

// ListHandler 订单列表处理器。
type ListHandler struct {
	queryRepo order.QueryRepository
}

// NewListHandler 创建 ListHandler。
func NewListHandler(queryRepo order.QueryRepository) *ListHandler {
	return &ListHandler{queryRepo: queryRepo}
}

// ListResult 列表查询结果。
type ListResult struct {
	Items []OrderDTO
	Total int64
}

// Handle 处理订单列表查询。
func (h *ListHandler) Handle(ctx context.Context, query ListQuery) (*ListResult, error) {
	var (
		entities []*order.Order
		total    int64
		err      error
	)

	if query.UserID > 0 {
		// 用户查询自己的订单
		entities, err = h.queryRepo.ListByUser(ctx, query.UserID, query.Offset, query.Limit)
		if err != nil {
			return nil, err
		}
		total, err = h.queryRepo.CountByUser(ctx, query.UserID)
	} else {
		// 管理员查询所有订单
		entities, err = h.queryRepo.List(ctx, query.Offset, query.Limit)
		if err != nil {
			return nil, err
		}
		total, err = h.queryRepo.Count(ctx)
	}

	if err != nil {
		return nil, err
	}

	return &ListResult{
		Items: ToOrderDTOs(entities),
		Total: total,
	}, nil
}
