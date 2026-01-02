package invoice_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/lwmacct/260101-go-pkg-ddd/internal/manualtest"
	appInvoice "github.com/lwmacct/260101-go-pkg-ddd/pkg/application/invoice"
	appOrder "github.com/lwmacct/260101-go-pkg-ddd/pkg/application/order"
)

// TestMain 在所有测试完成后清理测试数据。
func TestMain(m *testing.M) {
	code := m.Run()

	// 测试结束后，尝试清理残留的测试发票
	if os.Getenv("MANUAL") == "1" {
		cleanupTestInvoices()
	}

	os.Exit(code)
}

// cleanupTestInvoices 清理测试创建的发票。
func cleanupTestInvoices() {
	c := manualtest.NewClient()
	if _, err := c.Login("admin", "admin123"); err != nil {
		return // 登录失败则跳过清理
	}

	// 获取所有发票并删除测试数据（按备注判断）
	invoices, _, _ := manualtest.GetList[appInvoice.InvoiceDTO](c, invoiceBasePath(), map[string]string{"limit": "1000"})
	for _, inv := range invoices {
		if len(inv.Remark) >= 4 && inv.Remark[:4] == "test" {
			// 发票不支持删除，跳过
			continue
		}
	}
}

// invoiceBasePath 返回发票 API 基础路径
func invoiceBasePath() string {
	return "/api/admin/invoices"
}

// invoicePath 返回单个发票 API 路径
func invoicePath(id uint) string {
	return fmt.Sprintf("%s/%d", invoiceBasePath(), id)
}

// invoicePayPath 返回发票支付路径
func invoicePayPath(id uint) string {
	return fmt.Sprintf("%s/%d/pay", invoiceBasePath(), id)
}

// invoiceCancelPath 返回发票取消路径
func invoiceCancelPath(id uint) string {
	return fmt.Sprintf("%s/%d/cancel", invoiceBasePath(), id)
}

// invoiceRefundPath 返回发票退款路径
func invoiceRefundPath(id uint) string {
	return fmt.Sprintf("%s/%d/refund", invoiceBasePath(), id)
}

// createTestOrder 创建测试订单，返回订单 ID。
func createTestOrder(t *testing.T, c *manualtest.Client) uint {
	t.Helper()
	createReq := appOrder.CreateOrderDTO{
		ProductID: 1,
		Quantity:  1,
		Remark:    "test_invoice_order",
	}
	order, err := manualtest.Post[appOrder.OrderDTO](c, "/api/admin/orders", createReq)
	require.NoError(t, err, "创建测试订单失败")
	return order.ID
}

// TestInvoiceCRUD 测试发票完整流程。
//
// 手动运行:
//
//	MANUAL=1 go test -v -run TestInvoiceCRUD ./internal/manualtest/invoice/
func TestInvoiceCRUD(t *testing.T) {
	c := manualtest.LoginAsAdmin(t)

	// 步骤 1: 创建测试订单
	t.Log("\n步骤 1: 创建测试订单")
	orderID := createTestOrder(t, c)
	t.Cleanup(func() {
		_ = c.Delete(fmt.Sprintf("/api/admin/orders/%d", orderID))
	})
	t.Logf("  订单 ID: %d", orderID)

	// 步骤 2: 创建发票
	t.Log("\n步骤 2: 创建发票")
	createReq := appInvoice.CreateInvoiceDTO{
		OrderID: orderID,
		Amount:  100.00,
		DueDate: time.Now().Add(30 * 24 * time.Hour),
		Remark:  "test_invoice_crud",
	}
	createdInvoice, err := manualtest.Post[appInvoice.InvoiceDTO](c, invoiceBasePath(), createReq)
	require.NoError(t, err, "创建发票失败")
	invoiceID := createdInvoice.ID
	t.Logf("  创建成功! 发票 ID: %d, 发票号: %s", invoiceID, createdInvoice.InvoiceNo)
	t.Logf("  金额: %.2f, 状态: %s", createdInvoice.Amount, createdInvoice.Status)

	// 验证字段
	assert.Equal(t, orderID, createdInvoice.OrderID, "订单ID不匹配")
	assert.InDelta(t, createReq.Amount, createdInvoice.Amount, 0.01, "金额不匹配")
	assert.Equal(t, "pending", createdInvoice.Status, "初始状态应为 pending")
	assert.NotEmpty(t, createdInvoice.InvoiceNo, "发票号不应为空")

	// 步骤 3: 获取发票列表
	t.Log("\n步骤 3: 获取发票列表")
	invoices, meta, err := manualtest.GetList[appInvoice.InvoiceDTO](c, invoiceBasePath(), map[string]string{
		"page":  "1",
		"limit": "10",
	})
	require.NoError(t, err, "获取发票列表失败")
	t.Logf("  发票数量: %d", len(invoices))
	if meta != nil {
		t.Logf("  总数: %d, 总页数: %d", meta.Total, meta.TotalPages)
	}

	// 验证列表中包含创建的发票
	invoiceIDs := manualtest.ExtractIDs(invoices, func(inv appInvoice.InvoiceDTO) uint { return inv.ID })
	assert.Contains(t, invoiceIDs, invoiceID, "列表中应包含新创建的发票")

	// 步骤 4: 获取发票详情
	t.Log("\n步骤 4: 获取发票详情")
	invoiceDetail, err := manualtest.Get[appInvoice.InvoiceDTO](c, invoicePath(invoiceID), nil)
	require.NoError(t, err, "获取发票详情失败")
	t.Logf("  详情: 发票号 %s, 状态 %s", invoiceDetail.InvoiceNo, invoiceDetail.Status)
	assert.Equal(t, invoiceID, invoiceDetail.ID, "发票 ID 不匹配")

	t.Log("\n发票 CRUD 流程测试完成!")
}

// TestInvoicePartialPayment 测试发票部分支付流程。
//
// 手动运行:
//
//	MANUAL=1 go test -v -run TestInvoicePartialPayment ./internal/manualtest/invoice/
func TestInvoicePartialPayment(t *testing.T) {
	c := manualtest.LoginAsAdmin(t)

	// 创建测试订单和发票
	t.Log("\n步骤 1: 创建测试订单和发票")
	orderID := createTestOrder(t, c)
	t.Cleanup(func() {
		_ = c.Delete(fmt.Sprintf("/api/admin/orders/%d", orderID))
	})

	createReq := appInvoice.CreateInvoiceDTO{
		OrderID: orderID,
		Amount:  100.00,
		DueDate: time.Now().Add(30 * 24 * time.Hour),
		Remark:  "test_partial_payment",
	}
	createdInvoice, err := manualtest.Post[appInvoice.InvoiceDTO](c, invoiceBasePath(), createReq)
	require.NoError(t, err, "创建发票失败")
	invoiceID := createdInvoice.ID
	t.Logf("  发票 ID: %d, 金额: %.2f", invoiceID, createdInvoice.Amount)

	// 步骤 2: 第一次部分支付
	t.Log("\n步骤 2: 第一次部分支付 30.00")
	payReq := appInvoice.PayInvoiceDTO{
		Amount: 30.00,
	}
	paidInvoice, err := manualtest.Post[appInvoice.InvoiceDTO](c, invoicePayPath(invoiceID), payReq)
	require.NoError(t, err, "第一次支付失败")
	t.Logf("  已付金额: %.2f, 状态: %s", paidInvoice.PaidAmount, paidInvoice.Status)
	assert.InDelta(t, 30.00, paidInvoice.PaidAmount, 0.01, "已付金额应为 30.00")
	assert.Equal(t, "partial", paidInvoice.Status, "状态应为 partial")

	// 步骤 3: 第二次部分支付
	t.Log("\n步骤 3: 第二次部分支付 50.00")
	payReq.Amount = 50.00
	paidInvoice, err = manualtest.Post[appInvoice.InvoiceDTO](c, invoicePayPath(invoiceID), payReq)
	require.NoError(t, err, "第二次支付失败")
	t.Logf("  已付金额: %.2f, 状态: %s", paidInvoice.PaidAmount, paidInvoice.Status)
	assert.InDelta(t, 80.00, paidInvoice.PaidAmount, 0.01, "已付金额应为 80.00")
	assert.Equal(t, "partial", paidInvoice.Status, "状态应为 partial")

	// 步骤 4: 完成支付
	t.Log("\n步骤 4: 完成支付 20.00")
	payReq.Amount = 20.00
	paidInvoice, err = manualtest.Post[appInvoice.InvoiceDTO](c, invoicePayPath(invoiceID), payReq)
	require.NoError(t, err, "完成支付失败")
	t.Logf("  已付金额: %.2f, 状态: %s", paidInvoice.PaidAmount, paidInvoice.Status)
	assert.InDelta(t, 100.00, paidInvoice.PaidAmount, 0.01, "已付金额应为 100.00")
	assert.Equal(t, "paid", paidInvoice.Status, "状态应为 paid")
	assert.NotNil(t, paidInvoice.PaidAt, "支付时间不应为空")

	t.Log("\n部分支付流程测试完成!")
}

// TestInvoiceCancel 测试发票取消。
//
// 手动运行:
//
//	MANUAL=1 go test -v -run TestInvoiceCancel ./internal/manualtest/invoice/
func TestInvoiceCancel(t *testing.T) {
	c := manualtest.LoginAsAdmin(t)

	// 创建测试订单和发票
	t.Log("\n步骤 1: 创建测试订单和发票")
	orderID := createTestOrder(t, c)
	t.Cleanup(func() {
		_ = c.Delete(fmt.Sprintf("/api/admin/orders/%d", orderID))
	})

	createReq := appInvoice.CreateInvoiceDTO{
		OrderID: orderID,
		Amount:  100.00,
		DueDate: time.Now().Add(30 * 24 * time.Hour),
		Remark:  "test_invoice_cancel",
	}
	createdInvoice, err := manualtest.Post[appInvoice.InvoiceDTO](c, invoiceBasePath(), createReq)
	require.NoError(t, err, "创建发票失败")
	invoiceID := createdInvoice.ID
	t.Logf("  发票 ID: %d, 状态: %s", invoiceID, createdInvoice.Status)

	// 步骤 2: 取消发票
	t.Log("\n步骤 2: 取消发票")
	cancelledInvoice, err := manualtest.Post[appInvoice.InvoiceDTO](c, invoiceCancelPath(invoiceID), nil)
	require.NoError(t, err, "取消发票失败")
	t.Logf("  取消成功! 状态: %s", cancelledInvoice.Status)
	assert.Equal(t, "canceled", cancelledInvoice.Status, "状态应为 canceled")

	t.Log("\n发票取消测试完成!")
}

// TestInvoiceRefund 测试发票退款。
//
// 手动运行:
//
//	MANUAL=1 go test -v -run TestInvoiceRefund ./internal/manualtest/invoice/
func TestInvoiceRefund(t *testing.T) {
	c := manualtest.LoginAsAdmin(t)

	// 创建测试订单和发票
	t.Log("\n步骤 1: 创建测试订单和发票")
	orderID := createTestOrder(t, c)
	t.Cleanup(func() {
		_ = c.Delete(fmt.Sprintf("/api/admin/orders/%d", orderID))
	})

	createReq := appInvoice.CreateInvoiceDTO{
		OrderID: orderID,
		Amount:  100.00,
		DueDate: time.Now().Add(30 * 24 * time.Hour),
		Remark:  "test_invoice_refund",
	}
	createdInvoice, err := manualtest.Post[appInvoice.InvoiceDTO](c, invoiceBasePath(), createReq)
	require.NoError(t, err, "创建发票失败")
	invoiceID := createdInvoice.ID

	// 步骤 2: 先完成支付
	t.Log("\n步骤 2: 完成支付")
	payReq := appInvoice.PayInvoiceDTO{
		Amount: 100.00,
	}
	paidInvoice, err := manualtest.Post[appInvoice.InvoiceDTO](c, invoicePayPath(invoiceID), payReq)
	require.NoError(t, err, "支付失败")
	t.Logf("  支付成功! 状态: %s", paidInvoice.Status)
	assert.Equal(t, "paid", paidInvoice.Status, "状态应为 paid")

	// 步骤 3: 退款
	t.Log("\n步骤 3: 退款")
	refundedInvoice, err := manualtest.Post[appInvoice.InvoiceDTO](c, invoiceRefundPath(invoiceID), nil)
	require.NoError(t, err, "退款失败")
	t.Logf("  退款成功! 状态: %s", refundedInvoice.Status)
	assert.Equal(t, "refunded", refundedInvoice.Status, "状态应为 refunded")

	t.Log("\n发票退款测试完成!")
}

// TestInvoiceCancelAfterPartialPayment 测试部分支付后不能取消。
//
// 手动运行:
//
//	MANUAL=1 go test -v -run TestInvoiceCancelAfterPartialPayment ./internal/manualtest/invoice/
func TestInvoiceCancelAfterPartialPayment(t *testing.T) {
	c := manualtest.LoginAsAdmin(t)

	// 创建测试订单和发票
	t.Log("\n步骤 1: 创建测试订单和发票")
	orderID := createTestOrder(t, c)
	t.Cleanup(func() {
		_ = c.Delete(fmt.Sprintf("/api/admin/orders/%d", orderID))
	})

	createReq := appInvoice.CreateInvoiceDTO{
		OrderID: orderID,
		Amount:  100.00,
		DueDate: time.Now().Add(30 * 24 * time.Hour),
		Remark:  "test_cancel_after_partial",
	}
	createdInvoice, err := manualtest.Post[appInvoice.InvoiceDTO](c, invoiceBasePath(), createReq)
	require.NoError(t, err, "创建发票失败")
	invoiceID := createdInvoice.ID

	// 步骤 2: 部分支付
	t.Log("\n步骤 2: 部分支付 30.00")
	payReq := appInvoice.PayInvoiceDTO{
		Amount: 30.00,
	}
	paidInvoice, err := manualtest.Post[appInvoice.InvoiceDTO](c, invoicePayPath(invoiceID), payReq)
	require.NoError(t, err, "部分支付失败")
	t.Logf("  已付金额: %.2f, 状态: %s", paidInvoice.PaidAmount, paidInvoice.Status)
	assert.Equal(t, "partial", paidInvoice.Status)

	// 步骤 3: 尝试取消（应失败）
	t.Log("\n步骤 3: 尝试取消部分支付的发票（应失败）")
	_, err = manualtest.Post[appInvoice.InvoiceDTO](c, invoiceCancelPath(invoiceID), nil)
	require.Error(t, err, "部分支付后取消应该失败")
	t.Logf("  预期失败: %v", err)

	t.Log("\n部分支付后取消测试完成!")
}

// TestInvoiceNotFound 测试发票不存在的情况。
//
// 手动运行:
//
//	MANUAL=1 go test -v -run TestInvoiceNotFound ./internal/manualtest/invoice/
func TestInvoiceNotFound(t *testing.T) {
	c := manualtest.LoginAsAdmin(t)

	t.Log("测试获取不存在的发票...")

	// 使用一个不太可能存在的 ID
	_, err := manualtest.Get[appInvoice.InvoiceDTO](c, invoicePath(999999999), nil)
	require.Error(t, err, "获取不存在的发票应失败")
	t.Logf("  预期失败: %v", err)

	t.Log("\n发票不存在测试完成!")
}
