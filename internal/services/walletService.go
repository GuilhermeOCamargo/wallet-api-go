package services

import (
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/domain/models"
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/repository/adapters"
)

type WalletService interface {
	CreateWallet(w *models.Wallet) error
}

type walletServiceImpl struct {
	adapter adapters.WalletAdapter
}

func NewWalletService(adapter adapters.WalletAdapter) WalletService {
	return &walletServiceImpl{
		adapter: adapter,
	}
}

func (s walletServiceImpl) CreateWallet(w *models.Wallet) error {
	s.adapter.CreateWallet(w)
	return nil
}
