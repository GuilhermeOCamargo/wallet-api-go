package controllers

import (
	"log"
	"net/http"

	"github.com/GuilhermeOCamargo/go-wallet-api/presentation/requests"
	usecases "github.com/GuilhermeOCamargo/go-wallet-api/useCases"
	"github.com/gin-gonic/gin"
)

type WalletController interface {
	CreateWallet(c *gin.Context)
}

type walletControllerImpl struct {
	createWalletUseCase usecases.CreateWalletUseCase
}

func NewWalletController(createWalletUseCase usecases.CreateWalletUseCase) WalletController {
	return &walletControllerImpl{
		createWalletUseCase: createWalletUseCase,
	}
}

func (w walletControllerImpl) CreateWallet(c *gin.Context) {
	var request requests.WalletRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Panicln("Erro no bind do corpo da request", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	wallet, err := w.createWalletUseCase.Execute(request.ToDomain())
	if err != nil {
		log.Panicln("Erro ao criar nova wallet", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, wallet)
}
