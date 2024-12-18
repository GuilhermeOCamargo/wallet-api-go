package di

import "github.com/GuilhermeOCamargo/go-wallet-api/internal/presentation/controller"

type controllers struct {
	walletController controller.WalletController
	orderController  controller.OrderController
}

func InitializeControllers(useCases useCases) controllers {
	return controllers{
		walletController: controller.NewWalletController(useCases.createWalletUseCase),
		orderController:  controller.NewOrderController(useCases.createOrderUseCase),
	}
}

func (c *controllers) WalletController() controller.WalletController {
	return c.walletController
}

func (c *controllers) OrderController() controller.OrderController {
	return c.orderController
}
