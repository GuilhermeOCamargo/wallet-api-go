package usecases

import (
	"github.com/GuilhermeOCamargo/go-wallet-api/domain/models"
	"github.com/GuilhermeOCamargo/go-wallet-api/services"
)

type CreateWalletUseCase interface {
	Execute(wallet models.Wallet) (models.Wallet, error)
}

type createWalletUseCaseImpl struct {
	walletService services.WalletService
	ownerService  services.OwnerService
}

func NewCreateWalletUseCase(walletService services.WalletService, ownerService services.OwnerService) CreateWalletUseCase {
	return &createWalletUseCaseImpl{
		walletService: walletService,
		ownerService:  ownerService,
	}
}

func (w *createWalletUseCaseImpl) Execute(wallet models.Wallet) (models.Wallet, error) {
	//TODO alguma coisa
	return models.Wallet{}, nil
}
