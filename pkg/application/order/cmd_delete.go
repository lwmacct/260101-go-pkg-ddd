package order

import (
	"context"

	"github.com/lwmacct/260101-go-pkg-ddd/pkg/domain/order"
)

// DeleteHandler 删除订单处理器。
type DeleteHandler struct {
	cmdRepo   order.CommandRepository
	queryRepo order.QueryRepository
}

// NewDeleteHandler 创建 DeleteHandler。
func NewDeleteHandler(
	cmdRepo order.CommandRepository,
	queryRepo order.QueryRepository,
) *DeleteHandler {
	return &DeleteHandler{
		cmdRepo:   cmdRepo,
		queryRepo: queryRepo,
	}
}

// Handle 处理删除订单命令。
func (h *DeleteHandler) Handle(ctx context.Context, cmd DeleteCommand) error {
	// 先检查订单是否存在
	_, err := h.queryRepo.GetByID(ctx, cmd.ID)
	if err != nil {
		return err
	}

	return h.cmdRepo.Delete(ctx, cmd.ID)
}
