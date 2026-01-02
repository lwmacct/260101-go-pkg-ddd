package order

import "errors"

// 订单相关错误。
var (
	// ErrOrderNotFound 订单不存在。
	ErrOrderNotFound = errors.New("order not found")

	// ErrOrderNoExists 订单号已存在。
	ErrOrderNoExists = errors.New("order number already exists")

	// ErrInvalidStatus 无效的订单状态。
	ErrInvalidStatus = errors.New("invalid order status")

	// ErrCannotCancel 订单无法取消。
	ErrCannotCancel = errors.New("order cannot be canceled")

	// ErrCannotShip 订单无法发货。
	ErrCannotShip = errors.New("order cannot be shipped")
)
