package invoice

import "errors"

// 发票相关错误。
var (
	// ErrInvoiceNotFound 发票不存在。
	ErrInvoiceNotFound = errors.New("invoice not found")

	// ErrInvoiceNoExists 发票号已存在。
	ErrInvoiceNoExists = errors.New("invoice number already exists")
)

// 业务约束相关错误。
var (
	// ErrCannotPay 无法支付（状态不允许）。
	ErrCannotPay = errors.New("invoice cannot be paid")

	// ErrCannotCancel 无法取消（状态不允许）。
	ErrCannotCancel = errors.New("invoice cannot be canceled")

	// ErrCannotRefund 无法退款（状态不允许）。
	ErrCannotRefund = errors.New("invoice cannot be refunded")

	// ErrInvalidAmount 无效金额。
	ErrInvalidAmount = errors.New("invalid amount")
)
