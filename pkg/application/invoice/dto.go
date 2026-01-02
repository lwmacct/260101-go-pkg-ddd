package invoice

import "time"

// InvoiceDTO 发票数据传输对象。
type InvoiceDTO struct {
	ID              uint       `json:"id"`
	InvoiceNo       string     `json:"invoice_no"`
	OrderID         uint       `json:"order_id"`
	UserID          uint       `json:"user_id"`
	Amount          float64    `json:"amount"`
	PaidAmount      float64    `json:"paid_amount"`
	RemainingAmount float64    `json:"remaining_amount"`
	Status          string     `json:"status"`
	DueDate         time.Time  `json:"due_date"`
	PaidAt          *time.Time `json:"paid_at,omitempty"`
	Remark          string     `json:"remark,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

// CreateInvoiceDTO 创建发票请求。
type CreateInvoiceDTO struct {
	OrderID uint      `json:"order_id" binding:"required"`
	Amount  float64   `json:"amount" binding:"required,gt=0"`
	DueDate time.Time `json:"due_date" binding:"required"`
	Remark  string    `json:"remark"`
}

// PayInvoiceDTO 支付发票请求。
type PayInvoiceDTO struct {
	Amount float64 `json:"amount" binding:"required,gt=0"`
}

// ListResultDTO 发票列表结果。
type ListResultDTO struct {
	Items []*InvoiceDTO
	Total int64
}
