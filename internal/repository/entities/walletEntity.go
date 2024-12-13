package entities

import (
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/domain/models"
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	OwnerID              uint
	Owner                *Owner `gorm:"foreignKey:OwnerID"`
	BalanceAmountInCents int
}

func NewWallet(w *models.Wallet) *Wallet {
	return &Wallet{
		Owner:                NewOwner(w.Owner()),
		BalanceAmountInCents: w.BalanceAmountInCents(),
	}
}
