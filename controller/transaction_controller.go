package controller

import (
	"encoding/json"
	"ethereum-explorer/logger"
	"ethereum-explorer/service"
	"net/http"
)

type TransactionController struct {
	TransactionService service.TransactionService
}

func (tc *TransactionController) GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		logger.LogMethodNotAllowed(r)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	block, err := tc.TransactionService.GetTransactions(r)
	if err != nil {
		logger.LogInternalServerError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(block)

	w.Write(jsonData)
}

func (tc *TransactionController) GetTransactionByHash(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		logger.LogMethodNotAllowed(r)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	block, err := tc.TransactionService.GetTransactionByHash(r)
	if err != nil {
		logger.LogInternalServerError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(block)

	w.Write(jsonData)
}
