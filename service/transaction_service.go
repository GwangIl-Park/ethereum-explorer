package service

import (
	"context"
	"ethereum-explorer/model"
	"ethereum-explorer/repository"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionService interface {
	GetTransactions(c *gin.Context) ([]model.Transaction, error)
	GetTransactionByHash(c context.Context, hash string) (model.Transaction, error)
	GetTransactionsByAccount(c context.Context, account string) ([]model.Transaction, error)
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
	contextTimeout        time.Duration
}

func NewTransactionService(transactionRepository repository.TransactionRepository, timeout time.Duration) TransactionService {
	return &transactionService{
		transactionRepository,
		timeout,
	}
}

func (tu *transactionService) GetTransactions(c *gin.Context) ([]model.Transaction, error) {
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

func (tu *transactionService) GetTransactionByHash(c context.Context, hash string) (model.Transaction, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.transactionRepository.GetTransactionByHash(ctx, hash)
}

func (tu *transactionService) GetTransactionsByAccount(c context.Context, account string) ([]model.Transaction, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.transactionRepository.GetTransactionsByAccount(ctx, account)
}
