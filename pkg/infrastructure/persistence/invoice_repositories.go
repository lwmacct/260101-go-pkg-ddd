package persistence

import (
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/domain/invoice"
	"gorm.io/gorm"
)

// InvoiceRepositories 发票仓储聚合。
type InvoiceRepositories struct {
	Command invoice.CommandRepository
	Query   invoice.QueryRepository
}

// NewInvoiceRepositories 创建发票仓储聚合。
func NewInvoiceRepositories(db *gorm.DB) InvoiceRepositories {
	return InvoiceRepositories{
		Command: NewInvoiceCommandRepository(db),
		Query:   NewInvoiceQueryRepository(db),
	}
}
