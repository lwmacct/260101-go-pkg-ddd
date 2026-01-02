package handler

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/lwmacct/260101-go-pkg-ddd/pkg/adapters/http/ginutil"
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/application/invoice"
	invoiceDomain "github.com/lwmacct/260101-go-pkg-ddd/pkg/domain/invoice"
	"github.com/lwmacct/260101-go-pkg-gin/pkg/response"
)

// ListInvoicesQuery 发票列表查询参数
type ListInvoicesQuery struct {
	response.PaginationQueryDTO

	UserID  uint `form:"user_id"`
	OrderID uint `form:"order_id"`
}

// ToQuery 转换为 Application 层 Query 对象
func (q *ListInvoicesQuery) ToQuery() invoice.ListQuery {
	return invoice.ListQuery{
		UserID:  q.UserID,
		OrderID: q.OrderID,
		Offset:  q.GetOffset(),
		Limit:   q.GetLimit(),
	}
}

// InvoiceHandler 发票管理 Handler
type InvoiceHandler struct {
	createHandler *invoice.CreateHandler
	payHandler    *invoice.PayHandler
	cancelHandler *invoice.CancelHandler
	refundHandler *invoice.RefundHandler
	getHandler    *invoice.GetHandler
	listHandler   *invoice.ListHandler
}

// NewInvoiceHandler 创建发票管理 Handler
func NewInvoiceHandler(
	createHandler *invoice.CreateHandler,
	payHandler *invoice.PayHandler,
	cancelHandler *invoice.CancelHandler,
	refundHandler *invoice.RefundHandler,
	getHandler *invoice.GetHandler,
	listHandler *invoice.ListHandler,
) *InvoiceHandler {
	return &InvoiceHandler{
		createHandler: createHandler,
		payHandler:    payHandler,
		cancelHandler: cancelHandler,
		refundHandler: refundHandler,
		getHandler:    getHandler,
		listHandler:   listHandler,
	}
}

// Create 创建发票
//
//	@Summary		创建发票
//	@Description	管理员为订单创建发票
//	@Tags			Admin - Invoice Management
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			request	body		invoice.CreateInvoiceDTO					true	"发票信息"
//	@Success		201		{object}	response.DataResponse[invoice.InvoiceDTO]	"发票创建成功"
//	@Failure		400		{object}	response.ErrorResponse						"参数错误"
//	@Failure		401		{object}	response.ErrorResponse						"未授权"
//	@Failure		500		{object}	response.ErrorResponse						"服务器内部错误"
//	@Router			/api/admin/invoices [post]
func (h *InvoiceHandler) Create(c *gin.Context) {
	userID, ok := ginutil.GetUserID(c)
	if !ok {
		return
	}

	var req invoice.CreateInvoiceDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err.Error())
		return
	}

	result, err := h.createHandler.Handle(c.Request.Context(), invoice.CreateCommand{
		OrderID: req.OrderID,
		UserID:  userID,
		Amount:  req.Amount,
		DueDate: req.DueDate,
		Remark:  req.Remark,
	})
	if err != nil {
		if errors.Is(err, invoiceDomain.ErrInvalidAmount) {
			response.BadRequest(c, err.Error())
			return
		}
		response.InternalError(c, err.Error())
		return
	}

	response.Created(c, result)
}

// List 发票列表
//
//	@Summary		发票列表
//	@Description	分页获取发票列表
//	@Tags			Admin - Invoice Management
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			params	query		ListInvoicesQuery							false	"查询参数"
//	@Success		200		{object}	response.PagedResponse[invoice.InvoiceDTO]	"发票列表"
//	@Failure		401		{object}	response.ErrorResponse						"未授权"
//	@Failure		500		{object}	response.ErrorResponse						"服务器内部错误"
//	@Router			/api/admin/invoices [get]
func (h *InvoiceHandler) List(c *gin.Context) {
	var query ListInvoicesQuery
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

// Get 发票详情
//
//	@Summary		发票详情
//	@Description	获取发票详细信息
//	@Tags			Admin - Invoice Management
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		int											true	"发票ID"
//	@Success		200	{object}	response.DataResponse[invoice.InvoiceDTO]	"发票详情"
//	@Failure		401	{object}	response.ErrorResponse						"未授权"
//	@Failure		404	{object}	response.ErrorResponse						"发票不存在"
//	@Failure		500	{object}	response.ErrorResponse						"服务器内部错误"
//	@Router			/api/admin/invoices/{id} [get]
func (h *InvoiceHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "无效的发票ID")
		return
	}

	result, err := h.getHandler.Handle(c.Request.Context(), invoice.GetQuery{
		ID: uint(id),
	})
	if err != nil {
		if errors.Is(err, invoiceDomain.ErrInvoiceNotFound) {
			response.NotFoundMessage(c, err.Error())
			return
		}
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c, result)
}

// Pay 支付发票
//
//	@Summary		支付发票
//	@Description	支付发票（支持部分支付）
//	@Tags			Admin - Invoice Management
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id		path		int											true	"发票ID"
//	@Param			request	body		invoice.PayInvoiceDTO						true	"支付信息"
//	@Success		200		{object}	response.DataResponse[invoice.InvoiceDTO]	"支付成功"
//	@Failure		400		{object}	response.ErrorResponse						"参数错误或状态不允许"
//	@Failure		401		{object}	response.ErrorResponse						"未授权"
//	@Failure		404		{object}	response.ErrorResponse						"发票不存在"
//	@Failure		500		{object}	response.ErrorResponse						"服务器内部错误"
//	@Router			/api/admin/invoices/{id}/pay [post]
func (h *InvoiceHandler) Pay(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "无效的发票ID")
		return
	}

	var req invoice.PayInvoiceDTO
	if err = c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, err.Error())
		return
	}

	result, err := h.payHandler.Handle(c.Request.Context(), invoice.PayCommand{
		ID:     uint(id),
		Amount: req.Amount,
	})
	if err != nil {
		if errors.Is(err, invoiceDomain.ErrInvoiceNotFound) {
			response.NotFoundMessage(c, err.Error())
			return
		}
		if errors.Is(err, invoiceDomain.ErrCannotPay) ||
			errors.Is(err, invoiceDomain.ErrInvalidAmount) {
			response.BadRequest(c, err.Error())
			return
		}
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c, result)
}

// Cancel 取消发票
//
//	@Summary		取消发票
//	@Description	取消发票
//	@Tags			Admin - Invoice Management
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		int											true	"发票ID"
//	@Success		200	{object}	response.DataResponse[invoice.InvoiceDTO]	"取消成功"
//	@Failure		400	{object}	response.ErrorResponse						"状态不允许"
//	@Failure		401	{object}	response.ErrorResponse						"未授权"
//	@Failure		404	{object}	response.ErrorResponse						"发票不存在"
//	@Failure		500	{object}	response.ErrorResponse						"服务器内部错误"
//	@Router			/api/admin/invoices/{id}/cancel [post]
func (h *InvoiceHandler) Cancel(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "无效的发票ID")
		return
	}

	result, err := h.cancelHandler.Handle(c.Request.Context(), invoice.CancelCommand{
		ID: uint(id),
	})
	if err != nil {
		if errors.Is(err, invoiceDomain.ErrInvoiceNotFound) {
			response.NotFoundMessage(c, err.Error())
			return
		}
		if errors.Is(err, invoiceDomain.ErrCannotCancel) {
			response.BadRequest(c, err.Error())
			return
		}
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c, result)
}

// Refund 退款
//
//	@Summary		退款
//	@Description	发票退款
//	@Tags			Admin - Invoice Management
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		int											true	"发票ID"
//	@Success		200	{object}	response.DataResponse[invoice.InvoiceDTO]	"退款成功"
//	@Failure		400	{object}	response.ErrorResponse						"状态不允许"
//	@Failure		401	{object}	response.ErrorResponse						"未授权"
//	@Failure		404	{object}	response.ErrorResponse						"发票不存在"
//	@Failure		500	{object}	response.ErrorResponse						"服务器内部错误"
//	@Router			/api/admin/invoices/{id}/refund [post]
func (h *InvoiceHandler) Refund(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.ValidationError(c, "无效的发票ID")
		return
	}

	result, err := h.refundHandler.Handle(c.Request.Context(), invoice.RefundCommand{
		ID: uint(id),
	})
	if err != nil {
		if errors.Is(err, invoiceDomain.ErrInvoiceNotFound) {
			response.NotFoundMessage(c, err.Error())
			return
		}
		if errors.Is(err, invoiceDomain.ErrCannotRefund) {
			response.BadRequest(c, err.Error())
			return
		}
		response.InternalError(c, err.Error())
		return
	}

	response.OK(c, result)
}
