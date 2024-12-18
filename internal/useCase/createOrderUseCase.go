package useCase

import (
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/appErrors"
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/domain/enums"
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/domain/models"
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/domain/vo"
	"github.com/GuilhermeOCamargo/go-wallet-api/internal/service"
)

type CreateOrderUseCase interface {
	Execute(o *vo.CreateOrderVo) (*models.Order, error)
}

type createOrderUseCaseImpl struct {
	orderService       service.OrderService
	transactionService service.TransactionService
	walletService      service.WalletService
}

func NewCreateOrderUseCase(orderService service.OrderService,
	transactionService service.TransactionService,
	walletService service.WalletService) CreateOrderUseCase {
	return &createOrderUseCaseImpl{
		orderService:       orderService,
		transactionService: transactionService,
		walletService:      walletService,
	}
}

func (uc *createOrderUseCaseImpl) Execute(o *vo.CreateOrderVo) (*models.Order, error) {
	orderType, err := enums.NewOrderType(o.OrderType())
	if err != nil {
		return nil, err
	}

	handlers := map[*enums.OrderType]func(*vo.CreateOrderVo, *enums.OrderType) (*models.Order, error){
		&enums.ORDER_DEPOSIT:                  uc.handleDeposit,
		&enums.ORDER_WITHDRAW:                 uc.handleWithdraw,
		&enums.ORDER_TRANSFER_BETWEEN_WALLETS: uc.handleTransfer,
	}

	handler, exists := handlers[orderType]
	if !exists {
		return nil, appErrors.NewCodeError(422, "OrderType Inválido")
	}

	return handler(o, orderType)

}

func (uc *createOrderUseCaseImpl) handleDeposit(o *vo.CreateOrderVo, orderType *enums.OrderType) (*models.Order, error) {
	if !o.IsDepositDataValid() {
		return nil, appErrors.NewCodeError(422, "Dados inválidos para deposito")
	}

	wallet, err := uc.walletService.FindWalletById(o.CreditWalletId())
	if err != nil {
		return nil, err
	}

	return uc.processOrder(o, orderType, wallet, nil, &enums.OPERATION_CREDIT)
}

func (uc *createOrderUseCaseImpl) handleWithdraw(o *vo.CreateOrderVo, orderType *enums.OrderType) (*models.Order, error) {
	wallet, err := uc.walletService.FindWalletById(o.DebitWalletId())
	if err != nil {
		return nil, err
	}
	if !o.IsWithdrawDataValid(wallet.BalanceAmountInCents()) {
		return &models.Order{}, appErrors.NewCodeError(422, "Dados inválidos para saque")
	}

	return uc.processOrder(o, orderType, nil, wallet, &enums.OPERATION_DEBIT)

}

func (uc *createOrderUseCaseImpl) handleTransfer(o *vo.CreateOrderVo, orderType *enums.OrderType) (*models.Order, error) {
	debitWallet, err := uc.walletService.FindWalletById(o.DebitWalletId())
	if err != nil {
		return nil, err
	}

	creditWallet, err := uc.walletService.FindWalletById(o.CreditWalletId())
	if err != nil {
		return nil, err
	}

	if !o.IsTransferDataValid(debitWallet.BalanceAmountInCents()) {
		return &models.Order{}, appErrors.NewCodeError(422, "Dados inválidos para transferencia")
	}

	order, err := uc.createOrder(o, creditWallet, debitWallet, orderType)
	if err != nil {
		return nil, err
	}

	debitTransaction, err := uc.createTransaction(o, debitWallet, order, &enums.OPERATION_DEBIT)
	if err != nil {
		//TODO Remover Ordem
		return nil, err
	}
	creditTransaction, err := uc.createTransaction(o, creditWallet, order, &enums.OPERATION_DEBIT)
	if err != nil {
		//TODO Remover Ordem e transação
		return nil, err
	}
	order.AddTransaction(debitTransaction, creditTransaction)
	return order, nil
}

func (uc *createOrderUseCaseImpl) processOrder(o *vo.CreateOrderVo, orderType *enums.OrderType, creditWallet *models.Wallet, debitWallet *models.Wallet, operation *enums.Operation) (*models.Order, error) {
	order, err := uc.createOrder(o, creditWallet, debitWallet, orderType)
	if err != nil {
		return nil, err
	}
	var wallet *models.Wallet
	if operation == &enums.OPERATION_DEBIT {
		wallet = debitWallet
	} else {
		wallet = creditWallet
	}

	transaction, err := uc.createTransaction(o, wallet, order, operation)
	if err != nil {
		return nil, err
	}

	order.AddTransaction(transaction)
	return order, nil
}

func (uc *createOrderUseCaseImpl) createOrder(o *vo.CreateOrderVo, creditWallet *models.Wallet, debitWallet *models.Wallet, orderType *enums.OrderType) (*models.Order, error) {
	order := models.NewOrder(o, debitWallet, creditWallet, orderType)
	if err := uc.orderService.CreateOrder(order); err != nil {
		return nil, err
	}
	return order, nil
}

func (uc *createOrderUseCaseImpl) createTransaction(o *vo.CreateOrderVo, wallet *models.Wallet, order *models.Order, operation *enums.Operation) (*models.Transaction, error) {
	transaction := models.NewTransaction(wallet, o.AmountInCents(), operation, order)
	if err := uc.transactionService.CreateTransaction(transaction); err != nil {
		return nil, err
	}

	return transaction, nil
}
