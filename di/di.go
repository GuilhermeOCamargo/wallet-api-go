package di

import (
	"github.com/GuilhermeOCamargo/go-wallet-api/presentation/controllers"
	"github.com/GuilhermeOCamargo/go-wallet-api/services"
	usecases "github.com/GuilhermeOCamargo/go-wallet-api/useCases"
)

func Initialize() controllers.WalletController {
	walletService := services.NewWalletService()
	ownerService := services.NewOwnerService()

	createWalletUseCase := usecases.NewCreateWalletUseCase(walletService, ownerService)
	walletController := controllers.NewWalletController(createWalletUseCase)
	return walletController
}
