package routes

import (
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/presentation/controller"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.RouterGroup, controller controller.OrderController) {
	router.POST("/order", controller.CreateOrder)
}
