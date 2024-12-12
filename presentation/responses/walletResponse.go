package responses

import "github.com/GuilhermeOCamargo/go-wallet-api/domain/models"

type WalletResponse struct {
	Id                   uint   `json:"id"`
	BalanceAmountInCents int    `json:"balanceAmountInCents"`
	Name                 string `json:"name"`
	DocumentNumber       string `json:"documentNumber"`
}

func NewWalletResponse(w models.Wallet) *WalletResponse {
	return &WalletResponse{
		Id:                   w.Id(),
		BalanceAmountInCents: w.BalanceAmountInCents(),
		Name:                 w.Owner().Name(),
		DocumentNumber:       w.Owner().DocumentNumber(),
	}
}
