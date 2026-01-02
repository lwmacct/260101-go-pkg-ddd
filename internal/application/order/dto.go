package order

import "time"

// OrderDTO 订单响应 DTO。
type OrderDTO struct {
	ID          uint      `json:"id"`
	OrderNo     string    `json:"order_no"`
	UserID      uint      `json:"user_id"`
	ProductID   uint      `json:"product_id"`
	Quantity    int       `json:"quantity"`
	TotalAmount float64   `json:"total_amount"`
	Status      string    `json:"status"`
	Remark      string    `json:"remark,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateOrderDTO 创建订单请求 DTO。
type CreateOrderDTO struct {
	ProductID uint   `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required,min=1"`
	Remark    string `json:"remark"`
}

// UpdateOrderDTO 更新订单请求 DTO。
type UpdateOrderDTO struct {
	Remark string `json:"remark"`
}

// UpdateStatusDTO 更新订单状态请求 DTO。
type UpdateStatusDTO struct {
	Status string `json:"status" binding:"required,oneof=pending paid shipped completed canceled"`
}
