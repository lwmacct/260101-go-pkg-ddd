package order

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"github.com/oklog/ulid/v2"

	"github.com/lwmacct/260101-go-pkg-ddd/internal/domain/order"
	productDomain "github.com/lwmacct/260101-go-pkg-ddd/pkg/domain/product"
)

var (
	entropyMu sync.Mutex
	//nolint:gosec // ULID 熵源不需要加密级随机性，唯一性由时间戳保证
	entropy = ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
)

// generateOrderNo 生成唯一订单号（ULID：时间有序 + 全局唯一）。
func generateOrderNo() string {
	entropyMu.Lock()
	defer entropyMu.Unlock()
	return "ORD" + ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
}

// CreateHandler 创建订单处理器。
type CreateHandler struct {
	cmdRepo      order.CommandRepository
	productQuery productDomain.QueryRepository
}

// NewCreateHandler 创建 CreateHandler。
func NewCreateHandler(
	cmdRepo order.CommandRepository,
	productQuery productDomain.QueryRepository,
) *CreateHandler {
	return &CreateHandler{
		cmdRepo:      cmdRepo,
		productQuery: productQuery,
	}
}

// Handle 处理创建订单命令。
func (h *CreateHandler) Handle(ctx context.Context, cmd CreateCommand) (*OrderDTO, error) {
	// 获取产品信息计算价格
	product, err := h.productQuery.GetByID(ctx, cmd.ProductID)
	if err != nil {
		return nil, err
	}

	entity := &order.Order{
		OrderNo:     generateOrderNo(),
		UserID:      cmd.UserID,
		ProductID:   cmd.ProductID,
		Quantity:    cmd.Quantity,
		TotalAmount: product.Price * float64(cmd.Quantity),
		Status:      order.StatusPending,
		Remark:      cmd.Remark,
	}

	if err := h.cmdRepo.Create(ctx, entity); err != nil {
		return nil, err
	}

	return ToOrderDTO(entity), nil
}
