package service

import (
	"ethereum-explorer/dto"
	"ethereum-explorer/repository"
)

type TransactionService interface {
	GetTransactions() (*dto.GetTransactionsDTO, error)
	GetTransactionByHash(txHash string) (*dto.GetTransactionsByHashDTO, error)
	GetTransactionsByBlockNumber(blockNumber string) (*dto.GetTransactionsByBlockNumberDTO, error)
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionService(transactionRepository repository.TransactionRepository) TransactionService {
	return &transactionService{
		transactionRepository,
	}
}

func (tu *transactionService) GetTransactions() (*dto.GetTransactionsDTO, error) {
	return tu.transactionRepository.GetTransactions()
}

func (tu *transactionService) GetTransactionByHash(txHash string) (*dto.GetTransactionsByHashDTO, error) {
	return tu.transactionRepository.GetTransactionByHash(txHash)
}

func (tu *transactionService) GetTransactionsByBlockNumber(blockNumber string) (*dto.GetTransactionsByBlockNumberDTO, error) {
	return tu.transactionRepository.GetTransactionsByBlockNumber(blockNumber)
}