package postgres

import (
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/repository/entities"
	"gorm.io/gorm"
)

type WalletRepository interface {
	CreateWallet(*entities.Wallet)
	FindWalletById(string) (*entities.Wallet, error)
}

type walletRepositoryImpl struct {
	db *gorm.DB
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
	return &walletRepositoryImpl{db: db}
}

func (r *walletRepositoryImpl) CreateWallet(e *entities.Wallet) {
	r.db.Create(&e)
}

func (r *walletRepositoryImpl) FindWalletById(id string) (*entities.Wallet, error) {
	var wallet entities.Wallet
	r.db.First(&wallet, id)
	return &wallet, nil
}
