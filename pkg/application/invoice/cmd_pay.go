package invoice

import (
	"context"

	"github.com/lwmacct/260101-go-pkg-ddd/pkg/domain/invoice"
)

// PayHandler 支付发票处理器。
type PayHandler struct {
	cmdRepo   invoice.CommandRepository
	queryRepo invoice.QueryRepository
}

// NewPayHandler 创建 PayHandler。
func NewPayHandler(
	cmdRepo invoice.CommandRepository,
	queryRepo invoice.QueryRepository,
) *PayHandler {
	return &PayHandler{
		cmdRepo:   cmdRepo,
		queryRepo: queryRepo,
	}
}

// Handle 处理支付发票命令。
func (h *PayHandler) Handle(ctx context.Context, cmd PayCommand) (*InvoiceDTO, error) {
	if cmd.Amount <= 0 {
		return nil, invoice.ErrInvalidAmount
	}

	entity, err := h.queryRepo.GetByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}

	if !entity.CanPay() {
		return nil, invoice.ErrCannotPay
	}

	entity.Pay(cmd.Amount)

	if err := h.cmdRepo.Update(ctx, entity); err != nil {
		return nil, err
	}

	return ToInvoiceDTO(entity), nil
}
