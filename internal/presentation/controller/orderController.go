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

type OrderController interface {
	CreateOrder(c *gin.Context)
}

type orderControllerImpl struct {
	createOrderUseCase useCase.CreateOrderUseCase
}

func NewOrderController(createOrderUseCase useCase.CreateOrderUseCase) OrderController {
	return &orderControllerImpl{
		createOrderUseCase: createOrderUseCase,
	}
}

func (oc *orderControllerImpl) CreateOrder(c *gin.Context) {
	var request requests.OrderRequest
	log.Println("Inciando fluxo de criação de ordem")
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println("Erro no bind do corpo da request", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	orderVo := request.ToVo()
	order, err := oc.createOrderUseCase.Execute(orderVo)
	if err != nil {
		log.Println("Erro ao criar ordem", err.Error())
		status := 500
		var ce *appErrors.CodeError
		if errors.As(err, &ce) {
			status = ce.Status
		}
		c.JSON(status, err.Error())
		return
	}
	log.Println("Ordem criada com sucesso", order)
	c.JSON(http.StatusCreated, responses.NewOrderResponse(order))

}
