package persistence

import (
	"time"

	"github.com/lwmacct/260101-go-pkg-ddd/pkg/domain/invoice"
	"gorm.io/gorm"
)

// InvoiceModel 定义发票的 GORM 持久化模型。
//
//nolint:recvcheck // TableName uses value receiver per GORM convention
type InvoiceModel struct {
	ID         uint    `gorm:"primaryKey"`
	InvoiceNo  string  `gorm:"size:64;uniqueIndex;not null"`
	OrderID    uint    `gorm:"index;not null"`
	UserID     uint    `gorm:"index;not null"`
	Amount     float64 `gorm:"type:decimal(10,2);not null"`
	PaidAmount float64 `gorm:"type:decimal(10,2);not null;default:0"`
	Status     string  `gorm:"size:20;default:'pending';not null;index"`
	DueDate    time.Time
	PaidAt     *time.Time
	Remark     string `gorm:"size:500"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// TableName 指定发票表名。
func (InvoiceModel) TableName() string {
	return "invoices"
}

func newInvoiceModelFromEntity(entity *invoice.Invoice) *InvoiceModel {
	if entity == nil {
		return nil
	}

	return &InvoiceModel{
		ID:         entity.ID,
		InvoiceNo:  entity.InvoiceNo,
		OrderID:    entity.OrderID,
		UserID:     entity.UserID,
		Amount:     entity.Amount,
		PaidAmount: entity.PaidAmount,
		Status:     string(entity.Status),
		DueDate:    entity.DueDate,
		PaidAt:     entity.PaidAt,
		Remark:     entity.Remark,
		CreatedAt:  entity.CreatedAt,
		UpdatedAt:  entity.UpdatedAt,
	}
}

// ToEntity 将 GORM Model 转换为 Domain Entity。
func (m *InvoiceModel) ToEntity() *invoice.Invoice {
	if m == nil {
		return nil
	}

	return &invoice.Invoice{
		ID:         m.ID,
		InvoiceNo:  m.InvoiceNo,
		OrderID:    m.OrderID,
		UserID:     m.UserID,
		Amount:     m.Amount,
		PaidAmount: m.PaidAmount,
		Status:     invoice.Status(m.Status),
		DueDate:    m.DueDate,
		PaidAt:     m.PaidAt,
		Remark:     m.Remark,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
}

func mapInvoiceModelsToEntities(models []InvoiceModel) []*invoice.Invoice {
	if len(models) == 0 {
		return nil
	}

	invoices := make([]*invoice.Invoice, 0, len(models))
	for i := range models {
		if entity := models[i].ToEntity(); entity != nil {
			invoices = append(invoices, entity)
		}
	}
	return invoices
}
