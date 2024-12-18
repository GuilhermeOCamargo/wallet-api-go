package requests

import (
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/domain/vo"
)

type OrderRequest struct {
	DebitWalletId  string `json:"debitWalletId"`
	CreditWalletId string `json:"creditWalletId"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	AmountInCents  int    `json:"amountInCents"`
	OrderType      string `json:"orderType"`
}

func (o *OrderRequest) ToVo() *vo.CreateOrderVo {
	return vo.NewCreateOrderVo(o.DebitWalletId, o.CreditWalletId, o.OrderType,
		o.Title, o.Description, o.AmountInCents)
}
