package useCases

import (
	"log"

	"github.com/GuilhermeOCamargo/go-wallet-api/internal/domain/models"
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/services"
)

type CreateWalletUseCase interface {
	Execute(wallet *models.Wallet) error
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

func (w *createWalletUseCaseImpl) Execute(wallet *models.Wallet) error {
	log.Println("createWalletUseCaseImpl - Iniciando Execute", wallet)
	owner := wallet.Owner()
	if err := w.ownerService.CreateOwner(owner); err != nil {

		return err
	}

	log.Println("createWalletUseCaseImpl - owner criado. iniciando criação da wallet", owner)
	if err := w.walletService.CreateWallet(wallet); err != nil {
		return err
	}

	log.Println("createWalletUseCaseImpl - Wallet criada com sucesso", wallet)
	return nil
}
