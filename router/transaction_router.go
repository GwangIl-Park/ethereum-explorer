package router

import (
	"ethereum-explorer/controller"
	"net/http"
)

func NewTransactionRouter(transactionController controller.TransactionController, router *http.ServeMux) {
	router.HandleFunc("/transactions", transactionController.GetTransactions)
	router.HandleFunc("/transaction/hash/:hash", transactionController.GetTransactionByHash)
}
