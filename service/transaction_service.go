package service

import (
	"ethereum-explorer/dto"
	"ethereum-explorer/model"
	"ethereum-explorer/repository"
)

type TransactionService interface {
	GetTransactions() (*dto.GetTransactionsDTO, error)
	GetTransactionByHash(txHash string) (*dto.GetTransactionsByHashDTO, error)
	GetTransactionsByBlockNumber(blockNumber string) (*dto.GetTransactionsByBlockNumberDTO, error)
	CreateTransaction(transaction *model.Transaction) error
	CreateTransactions(transactions []*model.Transaction) error
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionService(transactionRepository repository.TransactionRepository) TransactionService {
	return &transactionService{
		transactionRepository,
	}
}

func (ts *transactionService) GetTransactions() (*dto.GetTransactionsDTO, error) {
	return ts.transactionRepository.GetTransactions()
}

func (ts *transactionService) GetTransactionByHash(txHash string) (*dto.GetTransactionsByHashDTO, error) {
	return ts.transactionRepository.GetTransactionByHash(txHash)
}

func (ts *transactionService) GetTransactionsByBlockNumber(blockNumber string) (*dto.GetTransactionsByBlockNumberDTO, error) {
	return ts.transactionRepository.GetTransactionsByBlockNumber(blockNumber)
}

func (ts *transactionService) CreateTransaction(transaction *model.Transaction) error {
	return ts.transactionRepository.CreateTransaction(transaction)
}

func (ts *transactionService) CreateTransactions(transactions []*model.Transaction) error {
	return ts.transactionRepository.CreateTransactions(transactions)
}