package requests

import "github.com/GuilhermeOCamargo/go-wallet-api/domain/models"

type WalletRequest struct {
	Id                   string `json:"id"`
	BalanceAmountInCents int    `json:"balanceAmountInCents"`
	Name                 string `json:"name"`
	DocumentNumber       string `json:"documentNumber"`
}

func (w *WalletRequest) ToDomain() models.Wallet {
	owner := models.Owner{
		DocumentNumber: w.DocumentNumber,
		Name:           w.Name,
	}
	return models.Wallet{
		Owner:                owner,
		BalanceAmountInCents: 0,
	}
}
