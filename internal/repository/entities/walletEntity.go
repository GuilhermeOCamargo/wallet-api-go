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
		OwnerID:              w.Owner().Id(),
		Owner:                NewOwner(w.Owner()),
		BalanceAmountInCents: w.BalanceAmountInCents(),
	}
}

func (w *Wallet) Id() uint {
	return w.ID
}

func (w *Wallet) ToDomain() *models.Wallet {
	return models.NewWallet(w.Owner.ToDomain(), w.BalanceAmountInCents, w.ID, w.CreatedAt, w.UpdatedAt)
}
