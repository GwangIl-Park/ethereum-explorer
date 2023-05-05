package repositories

import (
	"context"
	"database/sql"
	"ethereum-explorer/models"
)



type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) models.TransactionRepository {
	return &transactionRepository {
		db,
	}
}

func (tr *transactionRepository) GetTransactions(c context.Context) ([]models.Transaction, error) {
	rows, err := tr.db.Query("SELECT * FROM transactions")
	if err != nil {
		return nil, err
	}

	var id, blockHeight int
	var hash, from, to string
	var value, txFee float32
	var transactions []models.Transaction

	for rows.Next() {
		_ = rows.Scan(&id, &hash, &blockHeight, &from, &to, &value, &txFee)
		transactions = append(transactions, models.Transaction{Id:id, Hash: hash, BlockHeight: blockHeight, From: from, To: to, Value: value, TxFee: txFee})
	}

	return transactions, nil
}

func (tr *transactionRepository) GetTransactionByHash(c context.Context, hashParam string) (models.Transaction, error) {
	row := tr.db.QueryRow("SELECT * FROM transactions WHERE hash=?", hashParam)

	var id, blockHeight int
	var hash, from, to string
	var value, txFee float32

	row.Scan(&id, &hash, &blockHeight, &from, &to, &value, &txFee)

	return models.Transaction{Id: id, Hash: hash, BlockHeight: blockHeight, From: from, To: to, Value: value, TxFee: txFee}, nil
}

func (tr *transactionRepository) GetTransactionsByAccount(c context.Context, account string) ([]models.Transaction, error) {
	rows, err := tr.db.Query("SELECT * FROM transactions WHERE from=? OR to=?", account, account)
	if err != nil {
		return nil, err
	}

	var id, blockHeight int
	var hash, from, to string
	var value, txFee float32
	var transactions []models.Transaction

	for rows.Next() {
		_ = rows.Scan(&id, &hash, &blockHeight, &from, &to, &value, &txFee)
		transactions = append(transactions, models.Transaction{Id:id, Hash: hash, BlockHeight: blockHeight, From: from, To: to, Value: value, TxFee: txFee})
	}

	return transactions, nil
}