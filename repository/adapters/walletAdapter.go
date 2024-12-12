package adapters

import (
	"log"

	"github.com/GuilhermeOCamargo/go-wallet-api/domain/models"
	"github.com/GuilhermeOCamargo/go-wallet-api/repository/entities"
	"github.com/GuilhermeOCamargo/go-wallet-api/repository/postgres"
)

type WalletAdapter interface {
	CreateWallet(w *models.Wallet)
}

type walletAdapterImpl struct {
	repository postgres.WalletRepository
}

func NewWalletRepository(repository postgres.WalletRepository) WalletAdapter {
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
