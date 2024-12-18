package vo

type CreateOrderVo struct {
	debitWalletId, creditWalletId string
	title, description            string
	amountInCents                 int
	orderType                     string
}

func NewCreateOrderVo(debitWalletId, creditWalletId, orderType, title, description string, amountInCents int) *CreateOrderVo {
	return &CreateOrderVo{
		debitWalletId:  debitWalletId,
		creditWalletId: creditWalletId,
		title:          title,
		description:    description,
		amountInCents:  amountInCents,
		orderType:      orderType,
	}
}

func (vo *CreateOrderVo) DebitWalletId() string {
	return vo.debitWalletId
}

func (vo *CreateOrderVo) CreditWalletId() string {
	return vo.creditWalletId
}

func (vo *CreateOrderVo) Title() string {
	return vo.title
}

func (vo *CreateOrderVo) Description() string {
	return vo.description
}

func (vo *CreateOrderVo) AmountInCents() int {
	return vo.amountInCents
}

func (vo *CreateOrderVo) OrderType() string {
	return vo.orderType
}

func (vo *CreateOrderVo) IsDepositDataValid() bool {
	return vo.creditWalletId != "" && vo.amountInCents > 0
}

func (vo *CreateOrderVo) IsWithdrawDataValid(balance int) bool {
	return vo.debitWalletId != "" && vo.amountInCents > 0 && balance >= vo.amountInCents
}

func (vo *CreateOrderVo) IsTransferDataValid(balance int) bool {
	return vo.IsDepositDataValid() && vo.IsWithdrawDataValid(balance)
}
