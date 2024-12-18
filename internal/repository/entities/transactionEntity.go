package entities

import (
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/domain/models"
	"gorm.io/gorm"
)

type TransactionEntity struct {
	gorm.Model
	WalletID      uint
	Wallet        *Wallet `gorm:""foreignKey:WalletID`
	AmountInCents int
	Operation     string
	OrderID       uint
	OrderEntity   *OrderEntity `gorm:"foreignKey:OrderID"`
}

func NewTransactionEntity(t *models.Transaction) *TransactionEntity {
	return &TransactionEntity{
		WalletID:      t.Wallet().Id(),
		AmountInCents: t.AmountInCents(),
		Operation:     t.Operation().Name(),
		OrderID:       t.Order().Id(),
	}
}
