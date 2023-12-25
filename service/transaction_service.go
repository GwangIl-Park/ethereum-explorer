package service

import (
	"ethereum-explorer/dto"
	"ethereum-explorer/model"
	"ethereum-explorer/repository"
	"net/http"
)

type TransactionService interface {
	GetTransactions(r *http.Request) (*[]model.Transaction, error)
	GetTransactionByHash(r *http.Request) (*model.Transaction, error)
	GetTransactionsByBlockNumber(r *http.Request) (*[]dto.GetTransactionByBlockNumberDTO, error)
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionService(transactionRepository repository.TransactionRepository) TransactionService {
	return &transactionService{
		transactionRepository,
	}
}

func (tu *transactionService) GetTransactions(r *http.Request) (*[]model.Transaction, error) {
	return tu.transactionRepository.GetTransactions()
}

func (tu *transactionService) GetTransactionByHash(r *http.Request) (*model.Transaction, error) {
	hash := r.RequestURI[len("/transaction/"):]
	return tu.transactionRepository.GetTransactionByHash(hash)
}

func (tu *transactionService) GetTransactionsByBlockNumber(r *http.Request) (*[]dto.GetTransactionByBlockNumberDTO, error) {
	blockNumber := r.URL.Query().Get("blocknumber")
	return tu.transactionRepository.GetTransactionsByBlockNumber(blockNumber)
}