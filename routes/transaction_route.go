package routes

import (
	"ethereum-explorer/controller"
	"ethereum-explorer/db"
	"ethereum-explorer/repositories"
	"ethereum-explorer/usecases"
	"time"

	"github.com/gin-gonic/gin"
)
func NewTransactionRouter(timeout time.Duration, db *db.DB, group *gin.RouterGroup) {
	tr := repositories.NewTransactionRepository(db)
	tc := &controller.TransactionController{
		TransactionUsecase: usecases.NewTransactionUsecase(tr, timeout),
	}
	group.GET("/transactions", tc.GetTransactions)
	group.GET("/transaction/hash/:hash", tc.GetTransactionByHash)
	group.GET("/transaction/account/:account", tc.GetTransactionsByAccount)
}