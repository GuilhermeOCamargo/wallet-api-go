package postgres

import (
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/repository/entities"
	"gorm.io/gorm"
)

type WalletRepository interface {
	CreateWallet(*entities.Wallet)
}

type walletRepositoryImpl struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
	return &walletRepositoryImpl{db: db}
}

func (w *walletRepositoryImpl) CreateWallet(e *entities.Wallet) {
	w.db.Create(&e)
}
