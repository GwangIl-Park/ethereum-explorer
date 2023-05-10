package models

import (
	"context"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Hash string `json:"hash"`
	BlockHeight string `json:"blockHeight"`
	From string `json:"from"`
	To string `json:"to"`
	Value string `json:"value"`
	TxFee string `json:"txFee"`
}

type TransactionRepository interface{
	GetTransactions(c context.Context, page int64, show int64) ([]Transaction, error)
	GetTransactionByHash(c context.Context, hash string) (Transaction, error)
	GetTransactionsByAccount(c context.Context, account string) ([]Transaction, error)
}

type TransactionUsecase interface{
	GetTransactions(c *gin.Context) ([]Transaction, error)
	GetTransactionByHash(c context.Context, hash string) (Transaction, error)
	GetTransactionsByAccount(c context.Context, account string) ([]Transaction, error)
}