package service

import (
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/appErrors"
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/domain/models"
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/repository/adapter"
)

type WalletService interface {
	CreateWallet(w *models.Wallet) error
	FindWalletById(id string) (*models.Wallet, error)
}

type walletServiceImpl struct {
	adapter adapter.WalletAdapter
}

func NewWalletService(adapter adapter.WalletAdapter) WalletService {
	return &walletServiceImpl{
		adapter: adapter,
	}
}

func (s *walletServiceImpl) CreateWallet(w *models.Wallet) error {
	s.adapter.CreateWallet(w)
	return nil
}

func (s *walletServiceImpl) FindWalletById(id string) (*models.Wallet, error) {
	wallet, err := s.adapter.FindWalletById(id)
	if err != nil {
		return nil, err
	}

	if wallet.Id() == 0 {
		return nil, appErrors.NewCodeError(404, "Carteira não encontrada")
	}

	return wallet, nil
}
