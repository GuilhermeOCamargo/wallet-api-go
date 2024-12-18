package entities

import (
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/domain/models"
	"gorm.io/gorm"
)

type OrderEntity struct {
	gorm.Model

	DebitWalletID      uint
	DebitWallet        *Wallet `gorm:"foreignKey:DebitWalletID;references:ID"`
	CreditWalletID     uint
	CreditWallet       *Wallet `gorm:"foreignKey:CreditWalletID;references:ID"`
	Title, Description string
	AmountInCents      int
}

func NewOrder(o *models.Order) *OrderEntity {
	return &OrderEntity{
		DebitWalletID:  o.DebitWallet().Id(),
		CreditWalletID: o.CreditWallet().Id(),
		Title:          o.Title(),
		Description:    o.Description(),
		AmountInCents:  o.AmountInCents(),
	}
}
