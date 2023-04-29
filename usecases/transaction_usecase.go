package usecases

import (
	"context"
	"ethereum-explorer/models"
	"time"
)

type transactionUsecase struct {
	transactionRepository models.TransactionRepository
	contextTimeout time.Duration
}

func NewTransactionUsecase(transactionRepository models.TransactionRepository, timeout time.Duration) models.TransactionUsecase {
	return &transactionUsecase{
		transactionRepository,
		timeout,
	}
}

func (tu *transactionUsecase) CreateTransaction(c context.Context, transaction *models.Transaction) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.transactionRepository.CreateTransaction(ctx, transaction)
}

func (tu *transactionUsecase) GetTransactions(c context.Context) ([]models.Transaction, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.transactionRepository.GetTransactions(ctx)
}

func (tu *transactionUsecase) GetTransactionByHash(c context.Context, hash string) (models.Transaction, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.transactionRepository.GetTransactionByHash(ctx, hash)
}

func (tu *transactionUsecase) GetTransactionsByAccount(c context.Context, account string) ([]models.Transaction, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.transactionRepository.GetTransactionsByAccount(ctx, account)
}