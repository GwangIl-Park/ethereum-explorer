package repository

import (
	"ethereum-explorer/db"
	"ethereum-explorer/dto"
	"ethereum-explorer/model"
	"fmt"
)

type TransactionRepository interface {
	GetTransactions() (*dto.GetTransactionsDTO, error)
	GetTransactionByHash(hash string) (*dto.GetTransactionsByHashDTO, error)
	GetTransactionsByBlockNumber(blockNumber string) (*dto.GetTransactionsByBlockNumberDTO, error)
	CreateTransaction(transaction *model.Transaction) error
	CreateTransactions(transactions []*model.Transaction) error
}

type transactionRepository struct {
	db *db.DB
}

func NewTransactionRepository(db *db.DB) TransactionRepository {
	return &transactionRepository{
		db,
	}
}

func (tr *transactionRepository) GetTransactions() (*dto.GetTransactionsDTO, error) {
	rows, err := tr.db.Client.Query(`SELECT * FROM Transaction`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var getTransactionsDTO *dto.GetTransactionsDTO
	for rows.Next() {
		var transaction dto.GetTransactionsResult
		err = rows.Scan(&transaction)
		if err != nil {
			return nil, err
		}
		getTransactionsDTO.GetTransactionsResult = append(getTransactionsDTO.GetTransactionsResult, transaction)
	}

	return getTransactionsDTO, nil
}

func (tr *transactionRepository) GetTransactionByHash(hashParam string) (*dto.GetTransactionsByHashDTO, error) {
	var getTransactionsByHashDTO *dto.GetTransactionsByHashDTO
	err := tr.db.Client.QueryRow(`SELECT * FROM Transaction WHERE hash=%s`, hashParam).Scan(&getTransactionsByHashDTO.GetTransactionsResult)
	if err != nil {
		return nil, err
	}
	return getTransactionsByHashDTO, nil
}

func (tr *transactionRepository) GetTransactionsByBlockNumber(blockNumber string) (*dto.GetTransactionsByBlockNumberDTO, error) {
	rows, err := tr.db.Client.Query(`SELECT txHash, timestamp, from, to, value, txFee FROM transaction WHERE blockNumber = %s`, blockNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var getTransactionsByBlockNumberDTO *dto.GetTransactionsByBlockNumberDTO
	for rows.Next() {
		var getTransactionsByBlockNumberResult dto.GetTransactionsByBlockNumberResult
		err = rows.Scan(&getTransactionsByBlockNumberResult)
		if err != nil {
			return nil, err
		}
		getTransactionsByBlockNumberDTO.GetTransactionsByBlockNumberResult = append(getTransactionsByBlockNumberDTO.GetTransactionsByBlockNumberResult, getTransactionsByBlockNumberResult)
	}

	return getTransactionsByBlockNumberDTO, nil
}

func (tr *transactionRepository) CreateTransaction(transaction *model.Transaction) error {
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

func (tr *transactionRepository) CreateTransactions(transactions []*model.Transaction) error {
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
