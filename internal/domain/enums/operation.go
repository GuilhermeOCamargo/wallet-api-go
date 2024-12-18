package enums

import "github.com/GuilhermeOCamargo/go-wallet-api/internal/appErrors"

type Operation struct {
	name string
}

func (o *Operation) Name() string {
	return o.name
}

var (
	OPERATION_DEBIT  = Operation{"DEBIT"}
	OPERATION_CREDIT = Operation{"CREDIT"}
)

func NewOperation(operation string) (*Operation, error) {
	switch operation {
	case OPERATION_DEBIT.name:
		return &OPERATION_DEBIT, nil
	case OPERATION_CREDIT.name:
		return &OPERATION_CREDIT, nil
	default:
		return nil, appErrors.NewCodeError(422, "Operation Inválido")

	}
}
