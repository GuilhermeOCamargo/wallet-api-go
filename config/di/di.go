package di

import (
	"github.com/GuilhermeOCamargo/go-wallet-api/config/database"
	"github.com/GuilhermeOCamargo/go-wallet-api/presentation/controllers"
	"github.com/GuilhermeOCamargo/go-wallet-api/repository/adapters"
	"github.com/GuilhermeOCamargo/go-wallet-api/repository/postgres"
	"github.com/GuilhermeOCamargo/go-wallet-api/services"
	"github.com/GuilhermeOCamargo/go-wallet-api/useCases"
)

func Initialize() controllers.WalletController {
	ownerRepository, walletRepository := initializeRepositories()
	ownerAdapter, walletAdapter := initializeAdapters(ownerRepository, walletRepository)
	ownerService, walletService := initializeServices(ownerAdapter, walletAdapter)

	createWalletUseCase := useCases.NewCreateWalletUseCase(walletService, ownerService)
	walletController := controllers.NewWalletController(createWalletUseCase)

	return walletController
}

func initializeRepositories() (postgres.OwnerRepository, postgres.WalletRepository) {
	db := database.DB

	return postgres.NewOwnerRepository(db), postgres.NewWalletRepository(db)
}

func initializeAdapters(or postgres.OwnerRepository, wr postgres.WalletRepository) (adapters.OwnerAdapter, adapters.WalletAdapter) {
	return adapters.NewOwnerAdapter(or), adapters.NewWalletRepository(wr)
}

func initializeServices(oa adapters.OwnerAdapter, wa adapters.WalletAdapter) (services.OwnerService, services.WalletService) {
	return services.NewOwnerService(oa), services.NewWalletService(wa)
}
