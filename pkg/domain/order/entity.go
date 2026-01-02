package order

import (
	"time"
)

// Status 订单状态。
type Status string

const (
	// StatusPending 待支付。
	StatusPending Status = "pending"
	// StatusPaid 已支付。
	StatusPaid Status = "paid"
	// StatusShipped 已发货。
	StatusShipped Status = "shipped"
	// StatusCompleted 已完成。
	StatusCompleted Status = "completed"
	// StatusCancelled 已取消。
	StatusCancelled Status = "canceled"
)

// Order 订单实体。
type Order struct {
	ID          uint      `json:"id"`
	OrderNo     string    `json:"order_no"`
	UserID      uint      `json:"user_id"`
	ProductID   uint      `json:"product_id"`
	Quantity    int       `json:"quantity"`
	TotalAmount float64   `json:"total_amount"`
	Status      Status    `json:"status"`
	Remark      string    `json:"remark,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// IsPending 报告订单是否待支付。
func (o *Order) IsPending() bool {
	return o.Status == StatusPending
}

// CanCancel 报告订单是否可以取消。
func (o *Order) CanCancel() bool {
	return o.Status == StatusPending || o.Status == StatusPaid
}

// CanShip 报告订单是否可以发货。
func (o *Order) CanShip() bool {
	return o.Status == StatusPaid
}

// Pay 将订单状态变更为已支付。
func (o *Order) Pay() {
	o.Status = StatusPaid
}

// Ship 将订单状态变更为已发货。
func (o *Order) Ship() {
	o.Status = StatusShipped
}

// Complete 将订单标记为完成。
func (o *Order) Complete() {
	o.Status = StatusCompleted
}

// Cancel 将订单标记为取消。
func (o *Order) Cancel() {
	o.Status = StatusCancelled
}
