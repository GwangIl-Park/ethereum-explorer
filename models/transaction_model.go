package models

import (
	"context"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
)

type Transaction struct {
	TransactionHash string `json:"transactionHash"`
	Status          bool   `json:"status"`
	BlockHeight     string `json:"blockHeight"`
	Timestamp       string `json:"timestamp"`
	From            string `json:"from"`
	To              string `json:"to"`
	Value           string `json:"value"`
	TransactionFee  string `json:"transactionFee"`
	GasPrice        string `json:"gasPrice"`
	GasLimit        string `json:"gasLimit"`
	GasUsed         string `json:"gasUsed"`
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

func MakeTransactionModelFromTypes(receipt *types.Receipt, transaction types.Transaction, block types.Block) *Transaction {
	msg, err := core.TransactionToMessage(transaction, types.LatestSignerForChainID(transaction.ChainId()), nil)
	if err != nil {
		return nil, err
	}
	return &Transaction{
		TransactionHash: receipt.TxHash.Hex(),
		Status:          receipt.Status != 0,
		BlockHeight:     receipt.BlockNumber.String(),
		Timestamp:       block.ReceivedAt.String(),
		From:            msg.From.Hex(),
		To:              msg.To.String(),
		Value:           msg.Value.String(),
		TransactionFee:  (Add(receipt.EffectiveGasPrice, receipt.CumulativeGasUsed)),
		GasPrice:        receipt.EffectiveGasPrice.String(),
		GasLimit:        transaction.EffectiveGasTip(),
		GasUsed:         receipt.CumulativeGasUsed,
		Input:           string(transaction.Data()),
	}, nil
}
