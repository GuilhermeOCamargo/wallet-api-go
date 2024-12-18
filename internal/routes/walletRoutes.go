package routes

import (
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/presentation/controller"
	"github.com/gin-gonic/gin"
)

func WalletRoutes(router *gin.RouterGroup, walletController controller.WalletController) {
	router.POST("/wallet", walletController.CreateWallet)
}
