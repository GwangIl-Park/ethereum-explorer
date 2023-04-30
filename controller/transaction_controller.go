package controller

import (
	"ethereum-explorer/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	TransactionUsecase models.TransactionUsecase
}

func (tc *TransactionController) GetTransactions(c *gin.Context) {
	transactions, err := tc.TransactionUsecase.GetTransactions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

func (tc *TransactionController) GetTransactionByHash(c *gin.Context) {
	hash := c.Param("hash")
	transaction, err := tc.TransactionUsecase.GetTransactionByHash(c, hash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, transaction)
}

func (tc *TransactionController) GetTransactionsByAccount(c *gin.Context) {
	account := c.Param("account")
	transaction, err := tc.TransactionUsecase.GetTransactionsByAccount(c, account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, transaction)
}