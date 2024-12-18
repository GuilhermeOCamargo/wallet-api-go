package responses

import (
	"time"

	"github.com/GuilhermeOCamargo/go-wallet-api/internal/domain/models"
)

type TransactionResponse struct {
	Id            uint      `json:"id"`
	WalletId      uint      `json:"walletId"`
	AmountInCents int       `json:"amountInCents"`
	Operation     string    `json:"operation"`
	OrderId       uint      `json:"orderId"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func NewTransactionResponse(t *models.Transaction) *TransactionResponse {
	return &TransactionResponse{
		Id:            t.Id(),
		WalletId:      t.Wallet().Id(),
		AmountInCents: t.AmountInCents(),
		Operation:     t.Operation().Name(),
		OrderId:       t.Order().Id(),
		CreatedAt:     t.CreatedAt(),
		UpdatedAt:     t.UpdatedAt(),
	}
}
