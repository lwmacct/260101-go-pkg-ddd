package invoice

import (
	"context"

	"github.com/lwmacct/260101-go-pkg-ddd/pkg/domain/invoice"
)

// RefundHandler 退款处理器。
type RefundHandler struct {
	cmdRepo   invoice.CommandRepository
	queryRepo invoice.QueryRepository
}

// NewRefundHandler 创建 RefundHandler。
func NewRefundHandler(
	cmdRepo invoice.CommandRepository,
	queryRepo invoice.QueryRepository,
) *RefundHandler {
	return &RefundHandler{
		cmdRepo:   cmdRepo,
		queryRepo: queryRepo,
	}
}

// Handle 处理退款命令。
func (h *RefundHandler) Handle(ctx context.Context, cmd RefundCommand) (*InvoiceDTO, error) {
	entity, err := h.queryRepo.GetByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}

	if !entity.CanRefund() {
		return nil, invoice.ErrCannotRefund
	}

	entity.Refund()

	if err := h.cmdRepo.Update(ctx, entity); err != nil {
		return nil, err
	}

	return ToInvoiceDTO(entity), nil
}
