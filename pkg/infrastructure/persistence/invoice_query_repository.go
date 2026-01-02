package persistence

import (
	"context"
	"errors"

	"github.com/lwmacct/260101-go-pkg-ddd/pkg/domain/invoice"
	"gorm.io/gorm"
)

type invoiceQueryRepository struct {
	db *gorm.DB
}

// NewInvoiceQueryRepository 创建发票读仓储。
func NewInvoiceQueryRepository(db *gorm.DB) invoice.QueryRepository {
	return &invoiceQueryRepository{db: db}
}

func (r *invoiceQueryRepository) GetByID(ctx context.Context, id uint) (*invoice.Invoice, error) {
	var model InvoiceModel
	if err := r.db.WithContext(ctx).First(&model, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, invoice.ErrInvoiceNotFound
		}
		return nil, err
	}
	return model.ToEntity(), nil
}

func (r *invoiceQueryRepository) GetByInvoiceNo(ctx context.Context, invoiceNo string) (*invoice.Invoice, error) {
	var model InvoiceModel
	if err := r.db.WithContext(ctx).Where("invoice_no = ?", invoiceNo).First(&model).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, invoice.ErrInvoiceNotFound
		}
		return nil, err
	}
	return model.ToEntity(), nil
}

func (r *invoiceQueryRepository) ListByOrder(ctx context.Context, orderID uint) ([]*invoice.Invoice, error) {
	var models []InvoiceModel
	if err := r.db.WithContext(ctx).
		Where("order_id = ?", orderID).
		Order("created_at DESC").
		Find(&models).Error; err != nil {
		return nil, err
	}
	return mapInvoiceModelsToEntities(models), nil
}

func (r *invoiceQueryRepository) ListByUser(ctx context.Context, userID uint, offset, limit int) ([]*invoice.Invoice, error) {
	var models []InvoiceModel
	if err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&models).Error; err != nil {
		return nil, err
	}
	return mapInvoiceModelsToEntities(models), nil
}

func (r *invoiceQueryRepository) CountByUser(ctx context.Context, userID uint) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).
		Model(&InvoiceModel{}).
		Where("user_id = ?", userID).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *invoiceQueryRepository) List(ctx context.Context, offset, limit int) ([]*invoice.Invoice, error) {
	var models []InvoiceModel
	if err := r.db.WithContext(ctx).
		Order("created_at DESC").
		Offset(offset).
		Limit(limit).
		Find(&models).Error; err != nil {
		return nil, err
	}
	return mapInvoiceModelsToEntities(models), nil
}

func (r *invoiceQueryRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	if err := r.db.WithContext(ctx).Model(&InvoiceModel{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
