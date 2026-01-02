package invoice

import (
	"context"

	"github.com/lwmacct/260101-go-pkg-ddd/pkg/domain/invoice"
)

// GetHandler 获取发票详情处理器。
type GetHandler struct {
	queryRepo invoice.QueryRepository
}

// NewGetHandler 创建 GetHandler。
func NewGetHandler(queryRepo invoice.QueryRepository) *GetHandler {
	return &GetHandler{queryRepo: queryRepo}
}

// Handle 处理获取发票查询。
func (h *GetHandler) Handle(ctx context.Context, query GetQuery) (*InvoiceDTO, error) {
	entity, err := h.queryRepo.GetByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}
	return ToInvoiceDTO(entity), nil
}
