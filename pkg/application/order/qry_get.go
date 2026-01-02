package order

import (
	"context"

	"github.com/lwmacct/260101-go-pkg-ddd/pkg/domain/order"
)

// GetHandler 获取订单处理器。
type GetHandler struct {
	queryRepo order.QueryRepository
}

// NewGetHandler 创建 GetHandler。
func NewGetHandler(queryRepo order.QueryRepository) *GetHandler {
	return &GetHandler{queryRepo: queryRepo}
}

// Handle 处理获取订单查询。
func (h *GetHandler) Handle(ctx context.Context, query GetQuery) (*OrderDTO, error) {
	entity, err := h.queryRepo.GetByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}
	return ToOrderDTO(entity), nil
}
