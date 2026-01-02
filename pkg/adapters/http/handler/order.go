package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/lwmacct/260101-go-pkg-ddd/pkg/adapters/http/ginutil"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/application/order"
	orderDomain "github.com/lwmacct/260101-go-pkg-ddd/pkg/domain/order"
	"github.com/lwmacct/260101-go-pkg-gin/pkg/response"
)

// ListOrdersQuery 订单列表查询参数
type ListOrdersQuery struct {
	response.PaginationQueryDTO

	UserID uint `form:"user_id"`
}

// ToQuery 转换为 Application 层 Query 对象
func (q *ListOrdersQuery) ToQuery() order.ListQuery {
	return order.ListQuery{
		UserID: q.UserID,
		Offset: q.GetOffset(),
		Limit:  q.GetLimit(),
	}
}

// OrderHandler 订单管理 Handler
type OrderHandler struct {
	createHandler       *order.CreateHandler
	updateHandler       *order.UpdateHandler
	updateStatusHandler *order.UpdateStatusHandler
	deleteHandler       *order.DeleteHandler
	getHandler          *order.GetHandler
	listHandler         *order.ListHandler
}

// NewOrderHandler 创建订单管理 Handler
func NewOrderHandler(
	createHandler *order.CreateHandler,
	updateHandler *order.UpdateHandler,
	updateStatusHandler *order.UpdateStatusHandler,
	deleteHandler *order.DeleteHandler,
	getHandler *order.GetHandler,
	listHandler *order.ListHandler,
) *OrderHandler {
	return &OrderHandler{
		createHandler:       createHandler,
		updateHandler:       updateHandler,
		updateStatusHandler: updateStatusHandler,
		deleteHandler:       deleteHandler,
		getHandler:          getHandler,
		listHandler:         listHandler,
	}
}

// Create 创建订单
//
//	@Summary		创建订单
//	@Description	用户创建新订单
//	@Tags			Order Management
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			request	body		order.CreateOrderDTO				true	"订单信息"
//	@Success		201		{object}	response.DataResponse[order.OrderDTO]	"订单创建成功"
//	@Failure		400		{object}	response.ErrorResponse					"参数错误"
//	@Failure		401		{object}	response.ErrorResponse					"未授权"
//	@Failure		500		{object}	response.ErrorResponse					"服务器内部错误"
//	@Router			/api/orders [post]
func (h *OrderHandler) Create(c *gin.Context) {
	userID, ok := ginutil.GetUserID(c)
	if !ok {
		return // getUserID 已处理未授权响应
	}

	var req order.CreateOrderDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err.Error())
		return
	}

	result, err := h.createHandler.Handle(c.Request.Context(), order.CreateCommand{
		UserID:    userID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		Remark:    req.Remark,
	})
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Created(c, result)
}

// List 订单列表
//
//	@Summary		订单列表
//	@Description	分页获取订单列表（管理员可查看所有，普通用户只能查看自己的）
//	@Tags			Order Management
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			params	query		ListOrdersQuery							false	"查询参数"
//	@Success		200		{object}	response.PagedResponse[order.OrderDTO]	"订单列表"
//	@Failure		401		{object}	response.ErrorResponse					"未授权"
//	@Failure		500		{object}	response.ErrorResponse					"服务器内部错误"
//	@Router			/api/orders [get]
func (h *OrderHandler) List(c *gin.Context) {
	var query ListOrdersQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		response.ValidationError(c, err.Error())
		return
	}

	result, err := h.listHandler.Handle(c.Request.Context(), query.ToQuery())
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	meta := response.NewPaginationMeta(int(result.Total), query.GetPage(), query.GetLimit())
	response.List(c, result.Items, meta)
}

// Get 订单详情
//
//	@Summary		订单详情
//	@Description	获取订单详细信息
//	@Tags			Order Management
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		int									true	"订单ID"
//	@Success		200	{object}	response.DataResponse[order.OrderDTO]	"订单详情"
//	@Failure		401	{object}	response.ErrorResponse				"未授权"
//	@Failure		404	{object}	response.ErrorResponse				"订单不存在"
//	@Failure		500	{object}	response.ErrorResponse				"服务器内部错误"
//	@Router			/api/orders/{id} [get]
func (h *OrderHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "无效的订单ID")
		return
	}

	result, err := h.getHandler.Handle(c.Request.Context(), order.GetQuery{
		ID: uint(id),
	})
	if err != nil {
		if errors.Is(err, orderDomain.ErrOrderNotFound) {
			response.NotFoundMessage(c, err.Error())
			return
		}
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c, result)
}

// Update 更新订单
//
//	@Summary		更新订单
//	@Description	更新订单备注信息
//	@Tags			Order Management
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id		path		int									true	"订单ID"
//	@Param			request	body		order.UpdateOrderDTO				true	"更新信息"
//	@Success		200		{object}	response.DataResponse[order.OrderDTO]	"更新成功"
//	@Failure		400		{object}	response.ErrorResponse				"参数错误"
//	@Failure		401		{object}	response.ErrorResponse				"未授权"
//	@Failure		404		{object}	response.ErrorResponse				"订单不存在"
//	@Failure		500		{object}	response.ErrorResponse				"服务器内部错误"
//	@Router			/api/orders/{id} [put]
func (h *OrderHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "无效的订单ID")
		return
	}

	var req order.UpdateOrderDTO
	if err = c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err.Error())
		return
	}

	result, err := h.updateHandler.Handle(c.Request.Context(), order.UpdateCommand{
		ID:     uint(id),
		Remark: req.Remark,
	})
	if err != nil {
		if errors.Is(err, orderDomain.ErrOrderNotFound) {
			response.NotFoundMessage(c, err.Error())
			return
		}
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c, result)
}

// UpdateStatus 更新订单状态
//
//	@Summary		更新订单状态
//	@Description	更新订单状态（管理员操作）
//	@Tags			Order Management
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id		path		int									true	"订单ID"
//	@Param			request	body		order.UpdateStatusDTO				true	"状态信息"
//	@Success		200		{object}	response.DataResponse[order.OrderDTO]	"更新成功"
//	@Failure		400		{object}	response.ErrorResponse				"参数错误或状态转换无效"
//	@Failure		401		{object}	response.ErrorResponse				"未授权"
//	@Failure		404		{object}	response.ErrorResponse				"订单不存在"
//	@Failure		500		{object}	response.ErrorResponse				"服务器内部错误"
//	@Router			/api/orders/{id}/status [patch]
func (h *OrderHandler) UpdateStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "无效的订单ID")
		return
	}

	var req order.UpdateStatusDTO
	if err = c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err.Error())
		return
	}

	result, err := h.updateStatusHandler.Handle(c.Request.Context(), order.UpdateStatusCommand{
		ID:     uint(id),
		Status: req.Status,
	})
	if err != nil {
		if errors.Is(err, orderDomain.ErrOrderNotFound) {
			response.NotFoundMessage(c, err.Error())
			return
		}
		if errors.Is(err, orderDomain.ErrInvalidStatus) ||
			errors.Is(err, orderDomain.ErrCannotCancel) ||
			errors.Is(err, orderDomain.ErrCannotShip) {
			response.BadRequest(c, err.Error())
			return
		}
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c, result)
}

// Delete 删除订单
//
//	@Summary		删除订单
//	@Description	删除订单
//	@Tags			Order Management
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		int						true	"订单ID"
//	@Success		200	{object}	response.MessageResponse	"删除成功"
//	@Failure		401	{object}	response.ErrorResponse	"未授权"
//	@Failure		404	{object}	response.ErrorResponse	"订单不存在"
//	@Failure		500	{object}	response.ErrorResponse	"服务器内部错误"
//	@Router			/api/orders/{id} [delete]
func (h *OrderHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "无效的订单ID")
		return
	}

	if err := h.deleteHandler.Handle(c.Request.Context(), order.DeleteCommand{
		ID: uint(id),
	}); err != nil {
		if errors.Is(err, orderDomain.ErrOrderNotFound) {
			response.NotFoundMessage(c, err.Error())
			return
		}
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c, nil)
}
