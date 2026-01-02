package persistence

import (
	"github.com/lwmacct/260101-go-pkg-ddd/internal/domain/order"
	"gorm.io/gorm"
)

// OrderRepositories 聚合订单仓储。
type OrderRepositories struct {
	Command order.CommandRepository
	Query   order.QueryRepository
}

// NewOrderRepositories 创建订单仓储聚合。
func NewOrderRepositories(db *gorm.DB) OrderRepositories {
	return OrderRepositories{
		Command: NewOrderCommandRepository(db),
		Query:   NewOrderQueryRepository(db),
	}
}
