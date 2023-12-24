package router

import (
	"ethereum-explorer/controller"
	"net/http"
	"time"
)

func NewTransactionRouter(timeout time.Duration, transactionController controller.TransactionController, router *http.ServeMux) {
	router.HandleFunc("/transactions", transactionController.GetTransactions)
	router.HandleFunc("/transaction/hash/:hash", transactionController.GetTransactionByHash)
}
