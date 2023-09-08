package routes

import (
	"ethereum-explorer/controller"
	"ethereum-explorer/db"
	"ethereum-explorer/repositories"
	"ethereum-explorer/usecases"
	"net/http"
	"time"
)

func NewTransactionRouter(timeout time.Duration, db *db.DB, router *http.ServeMux) {
	tr := repositories.NewTransactionRepository(db)
	tc := &controller.TransactionController{
		TransactionUsecase: usecases.NewTransactionUsecase(tr, timeout),
	}
	router.HandleFunc("/transactions", tc.GetTransactions)
	router.HandleFunc("/transaction/hash/:hash", tc.GetTransactionByHash)
	router.HandleFunc("/transaction/account/:account", tc.GetTransactionsByAccount)
}
