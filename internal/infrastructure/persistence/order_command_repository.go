package persistence

import (
	"context"

	"github.com/lwmacct/260101-go-pkg-ddd/internal/domain/order"
	"gorm.io/gorm"
)

// orderCommandRepository 订单写仓储实现。
type orderCommandRepository struct {
	db *gorm.DB
}

// NewOrderCommandRepository 创建订单写仓储。
func NewOrderCommandRepository(db *gorm.DB) order.CommandRepository {
	return &orderCommandRepository{db: db}
}

// Create 创建订单。
func (r *orderCommandRepository) Create(ctx context.Context, entity *order.Order) error {
	model := newOrderModelFromEntity(entity)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return err
	}
	entity.ID = model.ID
	entity.CreatedAt = model.CreatedAt
	entity.UpdatedAt = model.UpdatedAt
	return nil
}

// Update 更新订单。
func (r *orderCommandRepository) Update(ctx context.Context, entity *order.Order) error {
	model := newOrderModelFromEntity(entity)
	return r.db.WithContext(ctx).Save(model).Error
}

// Delete 删除订单。
func (r *orderCommandRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&OrderModel{}, id).Error
}
