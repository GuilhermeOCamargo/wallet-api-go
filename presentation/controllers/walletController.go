package controllers

import (
	"log"
	"net/http"

	"github.com/GuilhermeOCamargo/go-wallet-api/domain/models"
	"github.com/GuilhermeOCamargo/go-wallet-api/presentation/requests"
	"github.com/GuilhermeOCamargo/go-wallet-api/presentation/responses"
	"github.com/GuilhermeOCamargo/go-wallet-api/useCases"
	"github.com/gin-gonic/gin"
)

type WalletController interface {
	CreateWallet(c *gin.Context)
	GetWalletById(c *gin.Context)
}

type walletControllerImpl struct {
	createWalletUseCase useCases.CreateWalletUseCase
}

func NewWalletController(createWalletUseCase useCases.CreateWalletUseCase) WalletController {
	return &walletControllerImpl{
		createWalletUseCase: createWalletUseCase,
	}
}

func (w walletControllerImpl) CreateWallet(c *gin.Context) {
	var request requests.WalletRequest
	log.Println("Iniciando fluxo para criação de wallet")
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
	log.Println("Wallet criada com sucesso", wallet)
	c.JSON(http.StatusCreated, responses.NewWalletResponse(wallet))
}

func (w walletControllerImpl) GetWalletById(c *gin.Context) {
	// id := c.Params.ByName("id")
	wallet := models.NewWallet(*models.NewOwner("name", "document"), 0)
	c.JSON(http.StatusOK, responses.NewWalletResponse(wallet))
}
