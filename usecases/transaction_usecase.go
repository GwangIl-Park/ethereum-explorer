package usecases

import (
	"context"
	"ethereum-explorer/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
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

func (tu *transactionUsecase) GetTransactions(c *gin.Context) ([]models.Transaction, error) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		return nil, err
	}
	show, err := strconv.Atoi(c.DefaultQuery("show", "10"))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.transactionRepository.GetTransactions(ctx, int64(page), int64(show))
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