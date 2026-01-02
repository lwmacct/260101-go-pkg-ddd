package invoice

import (
	"context"

	"github.com/lwmacct/260101-go-pkg-ddd/pkg/domain/invoice"
)

// CancelHandler 取消发票处理器。
type CancelHandler struct {
	cmdRepo   invoice.CommandRepository
	queryRepo invoice.QueryRepository
}

// NewCancelHandler 创建 CancelHandler。
func NewCancelHandler(
	cmdRepo invoice.CommandRepository,
	queryRepo invoice.QueryRepository,
) *CancelHandler {
	return &CancelHandler{
		cmdRepo:   cmdRepo,
		queryRepo: queryRepo,
	}
}

// Handle 处理取消发票命令。
func (h *CancelHandler) Handle(ctx context.Context, cmd CancelCommand) (*InvoiceDTO, error) {
	entity, err := h.queryRepo.GetByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}

	if !entity.CanCancel() {
		return nil, invoice.ErrCannotCancel
	}

	entity.Cancel()

	if err := h.cmdRepo.Update(ctx, entity); err != nil {
		return nil, err
	}

	return ToInvoiceDTO(entity), nil
}
