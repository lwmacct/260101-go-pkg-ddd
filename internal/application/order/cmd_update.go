package order

import (
	"context"

	"github.com/lwmacct/260101-go-pkg-ddd/internal/domain/order"
)

// UpdateHandler 更新订单处理器。
type UpdateHandler struct {
	cmdRepo   order.CommandRepository
	queryRepo order.QueryRepository
}

// NewUpdateHandler 创建 UpdateHandler。
func NewUpdateHandler(
	cmdRepo order.CommandRepository,
	queryRepo order.QueryRepository,
) *UpdateHandler {
	return &UpdateHandler{
		cmdRepo:   cmdRepo,
		queryRepo: queryRepo,
	}
}

// Handle 处理更新订单命令。
func (h *UpdateHandler) Handle(ctx context.Context, cmd UpdateCommand) (*OrderDTO, error) {
	entity, err := h.queryRepo.GetByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}

	entity.Remark = cmd.Remark

	if err := h.cmdRepo.Update(ctx, entity); err != nil {
		return nil, err
	}

	return ToOrderDTO(entity), nil
}

// UpdateStatusHandler 更新订单状态处理器。
type UpdateStatusHandler struct {
	cmdRepo   order.CommandRepository
	queryRepo order.QueryRepository
}

// NewUpdateStatusHandler 创建 UpdateStatusHandler。
func NewUpdateStatusHandler(
	cmdRepo order.CommandRepository,
	queryRepo order.QueryRepository,
) *UpdateStatusHandler {
	return &UpdateStatusHandler{
		cmdRepo:   cmdRepo,
		queryRepo: queryRepo,
	}
}

// Handle 处理更新订单状态命令。
func (h *UpdateStatusHandler) Handle(ctx context.Context, cmd UpdateStatusCommand) (*OrderDTO, error) {
	entity, err := h.queryRepo.GetByID(ctx, cmd.ID)
	if err != nil {
		return nil, err
	}

	newStatus := order.Status(cmd.Status)

	// 验证状态转换
	switch newStatus {
	case order.StatusPaid:
		entity.Pay()
	case order.StatusShipped:
		if !entity.CanShip() {
			return nil, order.ErrCannotShip
		}
		entity.Ship()
	case order.StatusCompleted:
		entity.Complete()
	case order.StatusCancelled:
		if !entity.CanCancel() {
			return nil, order.ErrCannotCancel
		}
		entity.Cancel()
	default:
		return nil, order.ErrInvalidStatus
	}

	if err := h.cmdRepo.Update(ctx, entity); err != nil {
		return nil, err
	}

	return ToOrderDTO(entity), nil
}
