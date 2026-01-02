package invoice

import (
	"github.com/lwmacct/260101-go-pkg-ddd/pkg/domain/invoice"
)

// ToInvoiceDTO 将发票实体转换为 DTO。
func ToInvoiceDTO(entity *invoice.Invoice) *InvoiceDTO {
	if entity == nil {
		return nil
	}
	return &InvoiceDTO{
		ID:              entity.ID,
		InvoiceNo:       entity.InvoiceNo,
		OrderID:         entity.OrderID,
		UserID:          entity.UserID,
		Amount:          entity.Amount,
		PaidAmount:      entity.PaidAmount,
		RemainingAmount: entity.RemainingAmount(),
		Status:          string(entity.Status),
		DueDate:         entity.DueDate,
		PaidAt:          entity.PaidAt,
		Remark:          entity.Remark,
		CreatedAt:       entity.CreatedAt,
		UpdatedAt:       entity.UpdatedAt,
	}
}

// ToInvoiceDTOs 将发票实体列表转换为 DTO 列表。
func ToInvoiceDTOs(entities []*invoice.Invoice) []*InvoiceDTO {
	if len(entities) == 0 {
		return nil
	}
	dtos := make([]*InvoiceDTO, 0, len(entities))
	for _, entity := range entities {
		if dto := ToInvoiceDTO(entity); dto != nil {
			dtos = append(dtos, dto)
		}
	}
	return dtos
}
