package repositories

import (
	"context"
	"ethereum-explorer/db"
	"ethereum-explorer/models"
)



type transactionRepository struct {
	db *db.DB
}

func NewTransactionRepository(db *db.DB) models.TransactionRepository {
	return &transactionRepository {
		db,
	}
}

func (tr *transactionRepository) GetTransactions(c context.Context) ([]models.Transaction, error) {
	documents := tr.db.ReadDocument("transactions", "", "")

	var transactions []models.Transaction

	for _, document := range documents {
		transactions = append(transactions, document.(models.Transaction))
	}
	
	return transactions, nil
}

func (tr *transactionRepository) GetTransactionByHash(c context.Context, hashParam string) (models.Transaction, error) {
	documents := tr.db.ReadDocument("transactions", "hash", hashParam)

	transaction := documents[0].(models.Transaction)
	
	return transaction, nil
}

func (tr *transactionRepository) GetTransactionsByAccount(c context.Context, account string) ([]models.Transaction, error) {
	documents := tr.db.ReadDocument("transactions", "account", account)

	var transactions []models.Transaction

	for _, document := range documents {
		transactions = append(transactions, document.(models.Transaction))
	}
	
	return transactions, nil
}