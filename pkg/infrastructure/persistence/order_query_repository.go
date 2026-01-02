package persistence

import (
	"context"
	"errors"

	"github.com/lwmacct/260101-go-pkg-ddd/pkg/domain/order"
	"gorm.io/gorm"
)

// orderQueryRepository 订单读仓储实现。
type orderQueryRepository struct {
	db *gorm.DB
}

// NewOrderQueryRepository 创建订单读仓储。
func NewOrderQueryRepository(db *gorm.DB) order.QueryRepository {
	return &orderQueryRepository{db: db}
}

// GetByID 根据 ID 获取订单。
func (r *orderQueryRepository) GetByID(ctx context.Context, id uint) (*order.Order, error) {
	var model OrderModel
	if err := r.db.WithContext(ctx).First(&model, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, order.ErrOrderNotFound
		}
		return nil, err
	}
	return model.ToEntity(), nil
}

// GetByOrderNo 根据订单号获取订单。
func (r *orderQueryRepository) GetByOrderNo(ctx context.Context, orderNo string) (*order.Order, error) {
	var model OrderModel
	if err := r.db.WithContext(ctx).Where("order_no = ?", orderNo).First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, order.ErrOrderNotFound
		}
		return nil, err
	}
	return model.ToEntity(), nil
}

// ListByUser 分页获取用户订单列表。
func (r *orderQueryRepository) ListByUser(ctx context.Context, userID uint, offset, limit int) ([]*order.Order, error) {
	var models []OrderModel
	if err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&models).Error; err != nil {
		return nil, err
	}
	return mapOrderModelsToEntities(models), nil
}

// CountByUser 获取用户订单总数。
func (r *orderQueryRepository) CountByUser(ctx context.Context, userID uint) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&OrderModel{}).Where("user_id = ?", userID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// List 分页获取订单列表（管理员）。
func (r *orderQueryRepository) List(ctx context.Context, offset, limit int) ([]*order.Order, error) {
	var models []OrderModel
	if err := r.db.WithContext(ctx).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&models).Error; err != nil {
		return nil, err
	}
	return mapOrderModelsToEntities(models), nil
}

// Count 获取订单总数。
func (r *orderQueryRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&OrderModel{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
