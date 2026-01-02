package order_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/lwmacct/260101-go-pkg-ddd/internal/manualtest"
	appOrder "github.com/lwmacct/260101-go-pkg-ddd/pkg/application/order"
)

// TestMain 在所有测试完成后清理测试数据。
func TestMain(m *testing.M) {
	code := m.Run()

	// 测试结束后，尝试清理残留的测试订单
	if os.Getenv("MANUAL") == "1" {
		cleanupTestOrders()
	}

	os.Exit(code)
}

// cleanupTestOrders 清理测试创建的订单。
func cleanupTestOrders() {
	c := manualtest.NewClient()
	if _, err := c.Login("admin", "admin123"); err != nil {
		return // 登录失败则跳过清理
	}

	// 获取所有订单并删除测试数据
	orders, _, _ := manualtest.GetList[appOrder.OrderDTO](c, orderBasePath(), map[string]string{"limit": "1000"})
	for _, o := range orders {
		// 删除备注包含 "test" 的订单
		if len(o.Remark) >= 4 && o.Remark[:4] == "test" {
			_ = c.Delete(orderPath(o.ID))
		}
	}
}

// orderBasePath 返回订单 API 基础路径
func orderBasePath() string {
	return "/api/admin/orders"
}

// orderPath 返回单个订单 API 路径
func orderPath(id uint) string {
	return fmt.Sprintf("%s/%d", orderBasePath(), id)
}

// orderStatusPath 返回订单状态更新路径
func orderStatusPath(id uint) string {
	return fmt.Sprintf("%s/%d/status", orderBasePath(), id)
}

// TestOrderCRUD 测试订单完整 CRUD 流程。
//
// 手动运行:
//
//	MANUAL=1 go test -v -run TestOrderCRUD ./internal/manualtest/order/
func TestOrderCRUD(t *testing.T) {
	c := manualtest.LoginAsAdmin(t)

	// 测试 1: 创建订单
	t.Log("\n测试 1: 创建订单")
	createReq := appOrder.CreateOrderDTO{
		ProductID: 1, // 假设产品 ID 1 存在
		Quantity:  2,
		Remark:    "test_order_crud",
	}
	createdOrder, err := manualtest.Post[appOrder.OrderDTO](c, orderBasePath(), createReq)
	require.NoError(t, err, "创建订单失败")
	orderID := createdOrder.ID
	t.Cleanup(func() {
		if orderID > 0 {
			_ = c.Delete(orderPath(orderID))
		}
	})
	t.Logf("  创建成功! 订单 ID: %d, 订单号: %s", orderID, createdOrder.OrderNo)
	t.Logf("  产品ID: %d, 数量: %d, 总金额: %.2f", createdOrder.ProductID, createdOrder.Quantity, createdOrder.TotalAmount)

	// 验证字段
	assert.Equal(t, createReq.ProductID, createdOrder.ProductID, "产品ID不匹配")
	assert.Equal(t, createReq.Quantity, createdOrder.Quantity, "数量不匹配")
	assert.Equal(t, "pending", createdOrder.Status, "初始状态应为 pending")
	assert.NotEmpty(t, createdOrder.OrderNo, "订单号不应为空")

	// 测试 2: 获取订单列表
	t.Log("\n测试 2: 获取订单列表")
	orders, meta, err := manualtest.GetList[appOrder.OrderDTO](c, orderBasePath(), map[string]string{
		"page":  "1",
		"limit": "10",
	})
	require.NoError(t, err, "获取订单列表失败")
	t.Logf("  订单数量: %d", len(orders))
	if meta != nil {
		t.Logf("  总数: %d, 总页数: %d", meta.Total, meta.TotalPages)
	}

	// 验证列表中包含创建的订单
	orderIDs := manualtest.ExtractIDs(orders, func(o appOrder.OrderDTO) uint { return o.ID })
	assert.Contains(t, orderIDs, orderID, "列表中应包含新创建的订单")

	// 测试 3: 获取订单详情
	t.Log("\n测试 3: 获取订单详情")
	orderDetail, err := manualtest.Get[appOrder.OrderDTO](c, orderPath(orderID), nil)
	require.NoError(t, err, "获取订单详情失败")
	t.Logf("  详情: 订单号 %s, 状态 %s", orderDetail.OrderNo, orderDetail.Status)
	assert.Equal(t, orderID, orderDetail.ID, "订单 ID 不匹配")

	// 测试 4: 更新订单（备注）
	t.Log("\n测试 4: 更新订单备注")
	updateReq := appOrder.UpdateOrderDTO{
		Remark: "test_order_updated_remark",
	}
	updatedOrder, err := manualtest.Put[appOrder.OrderDTO](c, orderPath(orderID), updateReq)
	require.NoError(t, err, "更新订单失败")
	t.Logf("  更新成功! 新备注: %s", updatedOrder.Remark)
	assert.Equal(t, updateReq.Remark, updatedOrder.Remark, "备注未更新")

	// 测试 5: 删除订单
	t.Log("\n测试 5: 删除订单")
	err = c.Delete(orderPath(orderID))
	require.NoError(t, err, "删除订单失败")
	t.Log("  删除成功!")

	// 标记已删除，避免 t.Cleanup 重复删除
	orderID = 0

	t.Log("\n订单 CRUD 流程测试完成!")
}

// TestOrderStatusTransition 测试订单状态转换。
//
// 手动运行:
//
//	MANUAL=1 go test -v -run TestOrderStatusTransition ./internal/manualtest/order/
func TestOrderStatusTransition(t *testing.T) {
	c := manualtest.LoginAsAdmin(t)

	// 创建测试订单
	t.Log("\n步骤 1: 创建测试订单")
	createReq := appOrder.CreateOrderDTO{
		ProductID: 1,
		Quantity:  1,
		Remark:    "test_status_transition",
	}
	createdOrder, err := manualtest.Post[appOrder.OrderDTO](c, orderBasePath(), createReq)
	require.NoError(t, err, "创建订单失败")
	orderID := createdOrder.ID
	t.Cleanup(func() {
		if orderID > 0 {
			_ = c.Delete(orderPath(orderID))
		}
	})
	t.Logf("  创建成功! 初始状态: %s", createdOrder.Status)
	assert.Equal(t, "pending", createdOrder.Status, "初始状态应为 pending")

	// 测试 2: pending -> paid
	t.Log("\n步骤 2: 更新状态 pending -> paid")
	statusReq := appOrder.UpdateStatusDTO{
		Status: "paid",
	}
	updatedOrder, err := manualtest.Patch[appOrder.OrderDTO](c, orderStatusPath(orderID), statusReq)
	require.NoError(t, err, "更新状态为 paid 失败")
	t.Logf("  状态更新为: %s", updatedOrder.Status)
	assert.Equal(t, "paid", updatedOrder.Status, "状态应为 paid")

	// 测试 3: paid -> shipped
	t.Log("\n步骤 3: 更新状态 paid -> shipped")
	statusReq.Status = "shipped"
	updatedOrder, err = manualtest.Patch[appOrder.OrderDTO](c, orderStatusPath(orderID), statusReq)
	require.NoError(t, err, "更新状态为 shipped 失败")
	t.Logf("  状态更新为: %s", updatedOrder.Status)
	assert.Equal(t, "shipped", updatedOrder.Status, "状态应为 shipped")

	// 测试 4: shipped -> completed
	t.Log("\n步骤 4: 更新状态 shipped -> completed")
	statusReq.Status = "completed"
	updatedOrder, err = manualtest.Patch[appOrder.OrderDTO](c, orderStatusPath(orderID), statusReq)
	require.NoError(t, err, "更新状态为 completed 失败")
	t.Logf("  状态更新为: %s", updatedOrder.Status)
	assert.Equal(t, "completed", updatedOrder.Status, "状态应为 completed")

	t.Log("\n状态转换测试完成!")
}

// TestOrderCancel 测试订单取消。
//
// 手动运行:
//
//	MANUAL=1 go test -v -run TestOrderCancel ./internal/manualtest/order/
func TestOrderCancel(t *testing.T) {
	c := manualtest.LoginAsAdmin(t)

	// 创建测试订单
	t.Log("\n步骤 1: 创建测试订单")
	createReq := appOrder.CreateOrderDTO{
		ProductID: 1,
		Quantity:  1,
		Remark:    "test_order_cancel",
	}
	createdOrder, err := manualtest.Post[appOrder.OrderDTO](c, orderBasePath(), createReq)
	require.NoError(t, err, "创建订单失败")
	orderID := createdOrder.ID
	t.Cleanup(func() {
		if orderID > 0 {
			_ = c.Delete(orderPath(orderID))
		}
	})
	t.Logf("  创建成功! 初始状态: %s", createdOrder.Status)

	// 测试 2: pending 状态可以取消
	t.Log("\n步骤 2: 取消 pending 状态的订单")
	statusReq := appOrder.UpdateStatusDTO{
		Status: "canceled",
	}
	updatedOrder, err := manualtest.Patch[appOrder.OrderDTO](c, orderStatusPath(orderID), statusReq)
	require.NoError(t, err, "取消订单失败")
	t.Logf("  状态更新为: %s", updatedOrder.Status)
	assert.Equal(t, "canceled", updatedOrder.Status, "状态应为 canceled")

	t.Log("\n订单取消测试完成!")
}

// TestOrderCancelAfterShip 测试发货后不能取消。
//
// 手动运行:
//
//	MANUAL=1 go test -v -run TestOrderCancelAfterShip ./internal/manualtest/order/
func TestOrderCancelAfterShip(t *testing.T) {
	c := manualtest.LoginAsAdmin(t)

	// 创建测试订单
	t.Log("\n步骤 1: 创建测试订单")
	createReq := appOrder.CreateOrderDTO{
		ProductID: 1,
		Quantity:  1,
		Remark:    "test_cancel_after_ship",
	}
	createdOrder, err := manualtest.Post[appOrder.OrderDTO](c, orderBasePath(), createReq)
	require.NoError(t, err, "创建订单失败")
	orderID := createdOrder.ID
	t.Cleanup(func() {
		if orderID > 0 {
			_ = c.Delete(orderPath(orderID))
		}
	})

	// 推进到 shipped 状态
	t.Log("\n步骤 2: 推进订单状态到 shipped")
	statusReq := appOrder.UpdateStatusDTO{Status: "paid"}
	_, err = manualtest.Patch[appOrder.OrderDTO](c, orderStatusPath(orderID), statusReq)
	require.NoError(t, err)

	statusReq.Status = "shipped"
	_, err = manualtest.Patch[appOrder.OrderDTO](c, orderStatusPath(orderID), statusReq)
	require.NoError(t, err)
	t.Log("  订单已发货")

	// 测试 3: 发货后不能取消
	t.Log("\n步骤 3: 尝试取消已发货订单（应失败）")
	statusReq.Status = "canceled"
	_, err = manualtest.Patch[appOrder.OrderDTO](c, orderStatusPath(orderID), statusReq)
	require.Error(t, err, "发货后取消应该失败")
	t.Logf("  预期失败: %v", err)

	t.Log("\n发货后取消测试完成!")
}

// TestOrderListPagination 测试订单列表分页。
//
// 手动运行:
//
//	MANUAL=1 go test -v -run TestOrderListPagination ./internal/manualtest/order/
func TestOrderListPagination(t *testing.T) {
	c := manualtest.LoginAsAdmin(t)

	t.Log("测试订单列表分页...")

	// 第一页
	page1, meta1, err := manualtest.GetList[appOrder.OrderDTO](c, orderBasePath(), map[string]string{
		"page":  "1",
		"limit": "5",
	})
	require.NoError(t, err, "获取第一页失败")
	t.Logf("  第 1 页: %d 条记录", len(page1))
	if meta1 != nil {
		t.Logf("  总数: %d, 总页数: %d", meta1.Total, meta1.TotalPages)
	}

	// 验证总数
	if meta1 != nil {
		assert.GreaterOrEqual(t, meta1.Total, 0, "总数应有效")
	}

	// 如果有第二页，获取并验证
	if meta1 != nil && meta1.TotalPages > 1 {
		page2, meta2, err := manualtest.GetList[appOrder.OrderDTO](c, orderBasePath(), map[string]string{
			"page":  "2",
			"limit": "5",
		})
		require.NoError(t, err, "获取第二页失败")
		t.Logf("  第 2 页: %d 条记录", len(page2))
		if meta2 != nil {
			assert.Equal(t, meta1.Total, meta2.Total, "两页总数应一致")
		}

		// 验证无重复 ID
		ids1 := manualtest.ExtractIDs(page1, func(o appOrder.OrderDTO) uint { return o.ID })
		ids2 := manualtest.ExtractIDs(page2, func(o appOrder.OrderDTO) uint { return o.ID })
		for _, id := range ids1 {
			assert.NotContains(t, ids2, id, "两页不应有相同 ID")
		}
	}

	t.Log("\n分页测试完成!")
}

// TestOrderNotFound 测试订单不存在的情况。
//
// 手动运行:
//
//	MANUAL=1 go test -v -run TestOrderNotFound ./internal/manualtest/order/
func TestOrderNotFound(t *testing.T) {
	c := manualtest.LoginAsAdmin(t)

	t.Log("测试获取不存在的订单...")

	// 使用一个不太可能存在的 ID
	_, err := manualtest.Get[appOrder.OrderDTO](c, orderPath(999999999), nil)
	require.Error(t, err, "获取不存在的订单应失败")
	t.Logf("  预期失败: %v", err)

	t.Log("\n订单不存在测试完成!")
}
