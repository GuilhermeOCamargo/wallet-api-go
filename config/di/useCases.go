package di

import "github.com/GuilhermeOCamargo/go-wallet-api/internal/useCase"

type useCases struct {
	createWalletUseCase useCase.CreateWalletUseCase
	createOrderUseCase  useCase.CreateOrderUseCase
}

func InitializeUseCases(services services) useCases {
	return useCases{
		createWalletUseCase: useCase.NewCreateWalletUseCase(services.walletService, services.ownerService),
		createOrderUseCase:  useCase.NewCreateOrderUseCase(services.orderService, services.transactionService, services.walletService),
	}
}
