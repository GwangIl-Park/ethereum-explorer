package models

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Hash        string `json:"hash"`
	BlockHeight string `json:"blockHeight"`
	From        string `json:"from"`
	To          string `json:"to"`
	Value       string `json:"value"`
	TxFee       string `json:"txFee"`
}

type TransactionRepository interface {
	GetTransactions(c context.Context, page int64, show int64) ([]Transaction, error)
	GetTransactionByHash(c context.Context, hash string) (Transaction, error)
	GetTransactionsByAccount(c context.Context, account string) ([]Transaction, error)
	CreateTransaction(c context.Context, transaction *Transaction) error
	CreateTransactions(c context.Context, transactions []*Transaction) error
}

type TransactionUsecase interface {
	GetTransactions(c *gin.Context) ([]Transaction, error)
	GetTransactionByHash(c context.Context, hash string) (Transaction, error)
	GetTransactionsByAccount(c context.Context, account string) ([]Transaction, error)
}

func MakeTransactionModelFromTypes(transaction *types.Transaction, height *big.Int) (*Transaction, error) {
	msg, err := core.TransactionToMessage(transaction, types.LatestSignerForChainID(transaction.ChainId()), nil)
	if err != nil {
		return nil, err
	}
	return &Transaction{
		Hash:        transaction.Hash().String(),
		BlockHeight: height.String(),
		From:        msg.From.String(),
		To:          transaction.To().String(),
		Value:       transaction.Value().String(),
		TxFee:       transaction.Cost().String(),
	}, nil
}
