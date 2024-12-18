package responses

import (
	"time"

	"github.com/GuilhermeOCamargo/go-wallet-api/internal/domain/models"
)

type OrderResponse struct {
	Id             uint                  `json:"id"`
	DebitWalletId  uint                  `json:"debitWalletId"`
	CreditWalletId uint                  `json:"creditWalletId"`
	CreatedAt      time.Time             `json:"createdAt"`
	UpdatedAt      time.Time             `json:"updatedAt"`
	Title          string                `json:"title"`
	Description    string                `json:"description"`
	AmountInCents  int                   `json:"amountInCents"`
	Transactions   []TransactionResponse `json:"transactions"`
	OrderType      string                `json:"orderType"`
}

func NewOrderResponse(o *models.Order) *OrderResponse {
	var transactions []TransactionResponse
	for _, t := range o.Transactions() {
		transactions = append(transactions, *NewTransactionResponse(t))
	}
	return &OrderResponse{
		Id:             o.Id(),
		DebitWalletId:  o.DebitWallet().Id(),
		CreditWalletId: o.CreditWallet().Id(),
		CreatedAt:      o.CreatedAt(),
		UpdatedAt:      o.UpdatedAt(),
		Title:          o.Title(),
		Description:    o.Description(),
		AmountInCents:  o.AmountInCents(),
		OrderType:      o.OrderType().Name(),
		Transactions:   transactions,
	}
}
