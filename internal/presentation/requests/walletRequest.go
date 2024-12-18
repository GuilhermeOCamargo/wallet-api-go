package requests

import (
	"time"

	"github.com/GuilhermeOCamargo/go-wallet-api/internal/domain/models"
)

type WalletRequest struct {
	Id                   uint      `json:"id"`
	BalanceAmountInCents int       `json:"balanceAmountInCents"`
	Name                 string    `json:"name"`
	DocumentNumber       string    `json:"documentNumber"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}

func (w *WalletRequest) ToDomain() *models.Wallet {
	owner := models.NewOwner(w.DocumentNumber, w.Name, w.Id, w.CreatedAt, w.UpdatedAt)
	return models.NewWalletWithOwnerAndBalance(owner, 0)
}
