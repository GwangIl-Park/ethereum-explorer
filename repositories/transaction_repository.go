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

func (tr *transactionRepository) GetTransactions(c context.Context, page int64, show int64) ([]models.Transaction, error) {
	//opts := options.Find().SetSort(bson.D{{Key: "blockheight",Value: -1}}).SetSkip((page-1) * show).SetLimit(show)
	rows, err := tr.db.Client.Query(`SELECT * FROM Transaction`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var transactions []models.Transaction
	for (rows.Next()) {
		var transaction models.Transaction
		err = rows.Scan(&transaction)
		if err != nil {
			panic(err)
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (tr *transactionRepository) GetTransactionByHash(c context.Context, hashParam string) (models.Transaction, error) {
	//cursor := tr.db.Collections["transactions"].FindOne(c, bson.M{"hash":hashParam})
	rows, err := tr.db.Client.Query(`SELECT * FROM Transaction WHERE hash=%s`, hashParam)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var transaction models.Transaction
	err = rows.Scan(&transaction)
	if err != nil {
		panic(err)
	}
	return transaction, nil
}

func (tr *transactionRepository) GetTransactionsByAccount(c context.Context, account string) ([]models.Transaction, error) {
	//cursor, err := tr.db.Collections["transactions"].Find(c, bson.M{"account":account})
	rows, err := tr.db.Client.Query(`SELECT * FROM Transaction WHERE from=%s OR to=%s`, account)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var transactions []models.Transaction
	for (rows.Next()) {
		var transaction models.Transaction
		err = rows.Scan(&transaction)
		if err != nil {
			panic(err)
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}