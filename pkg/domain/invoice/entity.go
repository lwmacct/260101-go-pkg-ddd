package invoice

import "time"

// Status 发票状态。
type Status string

const (
	// StatusDraft 草稿。
	StatusDraft Status = "draft"
	// StatusPending 待支付。
	StatusPending Status = "pending"
	// StatusPaid 已支付。
	StatusPaid Status = "paid"
	// StatusPartial 部分支付。
	StatusPartial Status = "partial"
	// StatusCancelled 已取消。
	StatusCancelled Status = "canceled"
	// StatusRefunded 已退款。
	StatusRefunded Status = "refunded"
)

// Invoice 发票实体。
//
// 发票是订单支付的凭证，支持部分支付场景。
// 一个订单可以对应多个发票（分期付款）。
type Invoice struct {
	ID         uint       `json:"id"`
	InvoiceNo  string     `json:"invoice_no"`  // 发票号（唯一）
	OrderID    uint       `json:"order_id"`    // 关联订单
	UserID     uint       `json:"user_id"`     // 用户 ID
	Amount     float64    `json:"amount"`      // 应付金额
	PaidAmount float64    `json:"paid_amount"` // 已付金额
	Status     Status     `json:"status"`
	DueDate    time.Time  `json:"due_date"`          // 到期日期
	PaidAt     *time.Time `json:"paid_at,omitempty"` // 支付完成时间
	Remark     string     `json:"remark,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

// IsPaid 报告发票是否已全额支付。
func (i *Invoice) IsPaid() bool {
	return i.Status == StatusPaid
}

// IsPending 报告发票是否待支付。
func (i *Invoice) IsPending() bool {
	return i.Status == StatusPending
}

// CanPay 报告发票是否可以支付。
func (i *Invoice) CanPay() bool {
	return i.Status == StatusPending || i.Status == StatusPartial
}

// CanCancel 报告发票是否可以取消。
func (i *Invoice) CanCancel() bool {
	return i.Status == StatusDraft || i.Status == StatusPending
}

// CanRefund 报告发票是否可以退款。
func (i *Invoice) CanRefund() bool {
	return i.Status == StatusPaid || i.Status == StatusPartial
}

// Pay 记录支付（支持部分支付）。
func (i *Invoice) Pay(amount float64) {
	i.PaidAmount += amount
	now := time.Now()
	if i.PaidAmount >= i.Amount {
		i.Status = StatusPaid
		i.PaidAt = &now
	} else {
		i.Status = StatusPartial
	}
}

// Submit 提交发票（从草稿变为待支付）。
func (i *Invoice) Submit() {
	if i.Status == StatusDraft {
		i.Status = StatusPending
	}
}

// Cancel 取消发票。
func (i *Invoice) Cancel() {
	i.Status = StatusCancelled
}

// Refund 退款。
func (i *Invoice) Refund() {
	i.Status = StatusRefunded
	i.PaidAmount = 0
}

// RemainingAmount 返回剩余应付金额。
func (i *Invoice) RemainingAmount() float64 {
	remaining := i.Amount - i.PaidAmount
	if remaining < 0 {
		return 0
	}
	return remaining
}
