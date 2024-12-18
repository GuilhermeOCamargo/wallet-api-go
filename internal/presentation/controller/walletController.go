package controller

import (
	"errors"
	"log"
	"net/http"

	"github.com/GuilhermeOCamargo/go-wallet-api/internal/appErrors"
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/presentation/requests"
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/presentation/responses"
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/useCase"
	"github.com/gin-gonic/gin"
)

type WalletController interface {
	CreateWallet(c *gin.Context)
}

type walletControllerImpl struct {
	createWalletUseCase useCase.CreateWalletUseCase
}

func NewWalletController(createWalletUseCase useCase.CreateWalletUseCase) WalletController {
	return &walletControllerImpl{
		createWalletUseCase: createWalletUseCase,
	}
}

func (w *walletControllerImpl) CreateWallet(c *gin.Context) {
	var request requests.WalletRequest
	log.Println("Iniciando fluxo para criação de wallet")
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Erro no bind do corpo da request", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	wallet := request.ToDomain()
	if err := w.createWalletUseCase.Execute(wallet); err != nil {
		log.Println("Erro ao criar nova wallet", err.Error())
		status := 500
		var ce *appErrors.CodeError
		if errors.As(err, &ce) {
			status = ce.Status
		}
		c.JSON(status, err.Error())
		return
	}
	log.Println("Wallet criada com sucesso", wallet)
	c.JSON(http.StatusCreated, responses.NewWalletResponse(wallet))
}
