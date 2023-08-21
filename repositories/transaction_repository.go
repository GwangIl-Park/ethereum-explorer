package repositories

import (
	"context"
	"ethereum-explorer/db"
	"ethereum-explorer/models"
	"fmt"
)

type transactionRepository struct {
	db *db.DB
}

func NewTransactionRepository(db *db.DB) models.TransactionRepository {
	return &transactionRepository{
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
	for rows.Next() {
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
	for rows.Next() {
		var transaction models.Transaction
		err = rows.Scan(&transaction)
		if err != nil {
			panic(err)
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (tr *transactionRepository) CreateTransaction(c context.Context, transaction *models.Transaction) error {
	valuesStr := fmt.Sprintf("(%s,%t,%s,%s,%s,%s,%s,%s,%v,%v,%s)",
		transaction.TransactionHash,
		transaction.Status,
		transaction.BlockHeight,
		transaction.From,
		transaction.To,
		transaction.Value,
		transaction.TransactionFee,
		transaction.GasPrice,
		transaction.GasLimit,
		transaction.GasUsed,
		transaction.Input,
	)

	_, err := tr.db.Client.Exec(`INSERT INTO Transaction VALUES %s`, valuesStr)
	if err != nil {
		return err
	}
	return nil
}

func (tr *transactionRepository) CreateTransactions(c context.Context, transactions []*models.Transaction) error {
	var valuesStr string

	for _, transaction := range transactions {
		valueStr := fmt.Sprintf("(%s,%t,%s,%s,%s,%s,%s,%s,%v,%v,%s)",
		transaction.TransactionHash,
		transaction.Status,
		transaction.BlockHeight,
		transaction.From,
		transaction.To,
		transaction.Value,
		transaction.TransactionFee,
		transaction.GasPrice,
		transaction.GasLimit,
		transaction.GasUsed,
		transaction.Input,
	)
		valuesStr = fmt.Sprintf("%s%s", valuesStr, valueStr)
	}

	_, err := tr.db.Client.Exec(`INSERT INTO Transaction VALUES %s`, valuesStr[:len(valuesStr)-1])
	if err != nil {
		return err
	}
	return nil
}
