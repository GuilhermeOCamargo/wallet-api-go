package adapter

import (
	"log"

	"github.com/GuilhermeOCamargo/go-wallet-api/internal/domain/models"
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/repository/entities"
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/repository/postgres"
)

type WalletAdapter interface {
	CreateWallet(w *models.Wallet)
	FindWalletById(id string) (*models.Wallet, error)
}

type walletAdapterImpl struct {
	repository postgres.WalletRepository
}

func NewWalletAdapter(repository postgres.WalletRepository) WalletAdapter {
	return &walletAdapterImpl{
		repository: repository,
	}
}

func (a *walletAdapterImpl) CreateWallet(w *models.Wallet) {
	entity := entities.NewWallet(w)

	a.repository.CreateWallet(entity)
	log.Println("Wallet criada no banco de dados", entity)

	w.SetId(entity.ID)
	w.SetCreatedAt(entity.CreatedAt)
	w.SetUpdatedAt(entity.UpdatedAt)
}

func (a *walletAdapterImpl) FindWalletById(id string) (*models.Wallet, error) {
	wallet, err := a.repository.FindWalletById(id)
	if err != nil {
		return nil, err
	}
	return wallet.ToDomain(), nil
}
