package router

import (
	"ethereum-explorer/controller"
	"ethereum-explorer/db"
	"ethereum-explorer/repository"
	"ethereum-explorer/usecase"
	"net/http"
	"time"
)

func NewTransactionRouter(timeout time.Duration, db *db.DB, router *http.ServeMux) {
	tr := repository.NewTransactionRepository(db)
	tc := &controller.TransactionController{
		TransactionUsecase: usecase.NewTransactionUsecase(tr, timeout),
	}
	router.HandleFunc("/transactions", tc.GetTransactions)
	router.HandleFunc("/transaction/hash/:hash", tc.GetTransactionByHash)
	router.HandleFunc("/transaction/account/:account", tc.GetTransactionsByAccount)
}
