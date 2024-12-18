package enums

import "github.com/GuilhermeOCamargo/go-wallet-api/internal/appErrors"

type OrderType struct {
	name string
}

func (t *OrderType) Name() string {
	return t.name
}

var (
	ORDER_TRANSFER_BETWEEN_WALLETS = OrderType{"TRANSFER_BETWEEN_WALLETS"}
	ORDER_DEPOSIT                  = OrderType{"DEPOSIT"}
	ORDER_WITHDRAW                 = OrderType{"WITHDRAW"}
)

func NewOrderType(orderType string) (*OrderType, error) {
	switch orderType {
	case ORDER_TRANSFER_BETWEEN_WALLETS.name:
		return &ORDER_TRANSFER_BETWEEN_WALLETS, nil
	case ORDER_DEPOSIT.name:
		return &ORDER_DEPOSIT, nil
	case ORDER_WITHDRAW.name:
		return &ORDER_WITHDRAW, nil
	default:
		return nil, appErrors.NewCodeError(422, "OrderType Inválido")
	}
}
