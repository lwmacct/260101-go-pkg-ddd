package persistence

import (
	"context"

	"github.com/lwmacct/260101-go-pkg-ddd/pkg/domain/invoice"
	"gorm.io/gorm"
)

type invoiceCommandRepository struct {
	db *gorm.DB
}

// NewInvoiceCommandRepository 创建发票写仓储。
func NewInvoiceCommandRepository(db *gorm.DB) invoice.CommandRepository {
	return &invoiceCommandRepository{db: db}
}

func (r *invoiceCommandRepository) Create(ctx context.Context, entity *invoice.Invoice) error {
	model := newInvoiceModelFromEntity(entity)
	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return err
	}
	entity.ID = model.ID
	entity.CreatedAt = model.CreatedAt
	entity.UpdatedAt = model.UpdatedAt
	return nil
}

func (r *invoiceCommandRepository) Update(ctx context.Context, entity *invoice.Invoice) error {
	model := newInvoiceModelFromEntity(entity)
	return r.db.WithContext(ctx).Save(model).Error
}

func (r *invoiceCommandRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&InvoiceModel{}, id).Error
}
