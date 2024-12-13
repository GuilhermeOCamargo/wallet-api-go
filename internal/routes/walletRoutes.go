package routes

import (
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/presentation/controllers"
	"github.com/gin-gonic/gin"
)

func WalletRoutes(router *gin.RouterGroup, walletController controllers.WalletController) {
	router.POST("/wallet", walletController.CreateWallet)
	router.GET("/wallet/id/:id", walletController.GetWalletById)
}
