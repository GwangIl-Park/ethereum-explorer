package service

import (
	"ethereum-explorer/model"
	"ethereum-explorer/repository"
	"net/http"
)

type TransactionService interface {
	GetTransactions(r *http.Request) ([]model.Transaction, error)
	GetTransactionByHash(r *http.Request) (model.Transaction, error)
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionService(transactionRepository repository.TransactionRepository) TransactionService {
	return &transactionService{
		transactionRepository,
	}
}

func (tu *transactionService) GetTransactions(r *http.Request) ([]model.Transaction, error) {
	return tu.transactionRepository.GetTransactions()
}

func (tu *transactionService) GetTransactionByHash(r *http.Request) (model.Transaction, error) {
	hash := r.RequestURI[len("/transaction/"):]
	return tu.transactionRepository.GetTransactionByHash(hash)
}
