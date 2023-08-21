package models

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
)

type Transaction struct {
	TransactionHash string `json:"transactionHash"`
	Status          bool   `json:"status"`
	BlockHeight     string `json:"blockHeight"`
	From            string `json:"from"`
	To              string `json:"to"`
	Value           string `json:"value"`
	TransactionFee  string `json:"transactionFee"`
	GasPrice        string `json:"gasPrice"`
	GasLimit        uint64 `json:"gasLimit"`
	GasUsed         uint64 `json:"gasUsed"`
	Input           string `json:"input"`
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

func MakeTransactionModelFromTypes(receipt *types.Receipt, transaction *types.Transaction, block types.Block) (*Transaction, error) {
	msg, err := core.TransactionToMessage(transaction, types.LatestSignerForChainID(transaction.ChainId()), nil)
	if err != nil {
		return nil, err
	}
	var transactionFee *big.Int
	transactionFee.Mul(receipt.EffectiveGasPrice, new(big.Int).SetUint64(receipt.CumulativeGasUsed)) 
	
	return &Transaction{
		TransactionHash: receipt.TxHash.Hex(),
		Status:          receipt.Status != 0,
		BlockHeight:     receipt.BlockNumber.String(),
		From:            msg.From.Hex(),
		To:              msg.To.String(),
		Value:           msg.Value.String(),
		TransactionFee:  transactionFee.String(),
		GasPrice:        receipt.EffectiveGasPrice.String(),
		GasLimit:        transaction.Gas(),
		GasUsed:         receipt.CumulativeGasUsed,
		Input:           string(transaction.Data()),
	}, nil
}
