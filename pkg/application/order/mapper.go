package order

import (
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/domain/order"
)

// ToOrderDTO 将实体转换为 DTO。
func ToOrderDTO(entity *order.Order) *OrderDTO {
	if entity == nil {
		return nil
	}

	return &OrderDTO{
		ID:          entity.ID,
		OrderNo:     entity.OrderNo,
		UserID:      entity.UserID,
		ProductID:   entity.ProductID,
		Quantity:    entity.Quantity,
		TotalAmount: entity.TotalAmount,
		Status:      string(entity.Status),
		Remark:      entity.Remark,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

// ToOrderDTOs 将实体列表转换为 DTO 列表。
func ToOrderDTOs(entities []*order.Order) []OrderDTO {
	if len(entities) == 0 {
		return []OrderDTO{}
	}

	dtos := make([]OrderDTO, 0, len(entities))
	for _, entity := range entities {
		if dto := ToOrderDTO(entity); dto != nil {
			dtos = append(dtos, *dto)
		}
	}
	return dtos
}
