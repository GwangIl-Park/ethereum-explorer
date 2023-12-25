package controller

import (
	"ethereum-explorer/httpResponse"
	"ethereum-explorer/logger"
	"ethereum-explorer/service"
	"net/http"
)

type TransactionController struct {
	TransactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) TransactionController {
	return TransactionController{
		transactionService,
	}
}

func (tc *TransactionController) GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		logger.LogMethodNotAllowed(r)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	blockNumber := r.URL.Query().Get("blocknumber")
	if blockNumber != "" {
		responseData, err := tc.TransactionService.GetTransactionsByBlockNumber(blockNumber)
		if err != nil {
			httpResponse.ErrorResponse(w, r, err)
			return
		}

		httpResponse.SendResponse(w, responseData)
	} else {
		responseData, err := tc.TransactionService.GetTransactions()
		if err != nil {
			httpResponse.ErrorResponse(w, r, err)
			return
		}
		httpResponse.SendResponse(w, responseData)
	}
}

func (tc *TransactionController) GetTransactionByHash(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		logger.LogMethodNotAllowed(r)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	txHash := r.RequestURI[len("/transaction/"):]

	responseData, err := tc.TransactionService.GetTransactionByHash(txHash)
	if err != nil {
		httpResponse.ErrorResponse(w, r, err)
	}
	
	httpResponse.SendResponse(w, responseData)
}