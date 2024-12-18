package di

import "github.com/GuilhermeOCamargo/go-wallet-api/internal/service"

type services struct {
	ownerService       service.OwnerService
	walletService      service.WalletService
	orderService       service.OrderService
	transactionService service.TransactionService
}

func InitializeServices(adapters adapters) services {
	return services{
		ownerService:       service.NewOwnerService(adapters.ownerAdapter),
		walletService:      service.NewWalletService(adapters.walletAdapter),
		orderService:       service.NewOrderService(adapters.orderAdapter),
		transactionService: service.NewTransactionService(adapters.transactionAdapter),
	}
}
