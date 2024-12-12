package responses

import "github.com/GuilhermeOCamargo/go-wallet-api/domain/models"

type WalletResponse struct {
	Id                   string `json:"id"`
	BalanceAmountInCents int    `json:"balanceAmountInCents"`
	Name                 string `json:"name"`
	DocumentNumber       string `json:"documentNumber"`
}

func NewWalletResponse(w models.Wallet) *WalletResponse {
	return &WalletResponse{
		Id:                   w.GetId(),
		BalanceAmountInCents: w.GetBalanceAmountInCents(),
		Name:                 w.GetOwner().GetName(),
		DocumentNumber:       w.GetOwner().GetDocumentNumber(),
	}
}
