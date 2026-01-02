package invoice

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"github.com/oklog/ulid/v2"

	"github.com/lwmacct/260101-go-pkg-ddd/pkg/domain/invoice"
)

var (
	entropyMu sync.Mutex
	//nolint:gosec // ULID 熵源不需要加密级随机性，唯一性由时间戳保证
	entropy = ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
)

// generateInvoiceNo 生成唯一发票号（ULID：时间有序 + 全局唯一）。
func generateInvoiceNo() string {
	entropyMu.Lock()
	defer entropyMu.Unlock()
	return "INV" + ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
}

// CreateHandler 创建发票处理器。
type CreateHandler struct {
	cmdRepo   invoice.CommandRepository
	queryRepo invoice.QueryRepository
}

// NewCreateHandler 创建 CreateHandler。
func NewCreateHandler(
	cmdRepo invoice.CommandRepository,
	queryRepo invoice.QueryRepository,
) *CreateHandler {
	return &CreateHandler{
		cmdRepo:   cmdRepo,
		queryRepo: queryRepo,
	}
}

// Handle 处理创建发票命令。
func (h *CreateHandler) Handle(ctx context.Context, cmd CreateCommand) (*InvoiceDTO, error) {
	if cmd.Amount <= 0 {
		return nil, invoice.ErrInvalidAmount
	}

	entity := &invoice.Invoice{
		InvoiceNo: generateInvoiceNo(),
		OrderID:   cmd.OrderID,
		UserID:    cmd.UserID,
		Amount:    cmd.Amount,
		Status:    invoice.StatusPending,
		DueDate:   cmd.DueDate,
		Remark:    cmd.Remark,
	}

	if err := h.cmdRepo.Create(ctx, entity); err != nil {
		return nil, err
	}

	return ToInvoiceDTO(entity), nil
}
