package models

import "context"

type Transaction struct {
	Hash string `json:"hash"`
	BlockHeight string `json:"blockHeight"`
	From string `json:"from"`
	To string `json:"to"`
	Value string `json:"value"`
	TxFee string `json:"txFee"`
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