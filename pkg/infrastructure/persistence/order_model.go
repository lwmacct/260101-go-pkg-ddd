package persistence

import (
	"time"

	"github.com/lwmacct/260101-go-pkg-ddd/pkg/domain/order"
	"gorm.io/gorm"
)

// OrderModel 定义订单的 GORM 持久化模型。
//
//nolint:recvcheck // TableName uses value receiver per GORM convention
type OrderModel struct {
	ID          uint    `gorm:"primaryKey"`
	OrderNo     string  `gorm:"size:64;uniqueIndex;not null"`
	UserID      uint    `gorm:"index;not null"`
	ProductID   uint    `gorm:"index;not null"`
	Quantity    int     `gorm:"not null;default:1"`
	TotalAmount float64 `gorm:"type:decimal(10,2);not null"`
	Status      string  `gorm:"size:20;default:'pending';not null;index"`
	Remark      string  `gorm:"size:500"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// TableName 指定订单表名。
func (OrderModel) TableName() string {
	return "orders"
}

func newOrderModelFromEntity(entity *order.Order) *OrderModel {
	if entity == nil {
		return nil
	}

	return &OrderModel{
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

// ToEntity 将 GORM Model 转换为 Domain Entity。
func (m *OrderModel) ToEntity() *order.Order {
	if m == nil {
		return nil
	}

	return &order.Order{
		ID:          m.ID,
		OrderNo:     m.OrderNo,
		UserID:      m.UserID,
		ProductID:   m.ProductID,
		Quantity:    m.Quantity,
		TotalAmount: m.TotalAmount,
		Status:      order.Status(m.Status),
		Remark:      m.Remark,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func mapOrderModelsToEntities(models []OrderModel) []*order.Order {
	if len(models) == 0 {
		return nil
	}

	orders := make([]*order.Order, 0, len(models))
	for i := range models {
		if entity := models[i].ToEntity(); entity != nil {
			orders = append(orders, entity)
		}
	}
	return orders
}
