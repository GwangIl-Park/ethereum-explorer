package models

import "context"

type Transaction struct {
	Id int `json:"id"`
	Hash string `json:"hash"`
	BlockHeight int `json:"blockHeight"`
	From string `json:"from"`
	To string `json:"to"`
	Value float32 `json:"value"`
	TxFee float32 `json:"txFee"`
}

type TransactionRepository interface{
	GetTransactions(c context.Context) ([]Transaction, error)
	GetTransactionByHash(c context.Context, hash string) (Transaction, error)
	GetTransactionsByAccount(c context.Context, account string) ([]Transaction, error)
}

type TransactionUsecase interface{
	GetTransactions(c context.Context) ([]Transaction, error)
	GetTransactionByHash(c context.Context, hash string) (Transaction, error)
	GetTransactionsByAccount(c context.Context, account string) ([]Transaction, error)
}