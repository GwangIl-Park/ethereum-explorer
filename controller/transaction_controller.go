package controller

import (
	"ethereum-explorer/model"
	"net/http"
)

type TransactionController struct {
	TransactionService model.TransactionService
}

func (tc *TransactionController) GetTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := tc.TransactionService.GetTransactions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

func (tc *TransactionController) GetTransactionByHash(w http.ResponseWriter, r *http.Request) {
	hash := c.Param("hash")
	transaction, err := tc.TransactionService.GetTransactionByHash(c, hash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, transaction)
}

func (tc *TransactionController) GetTransactionsByAccount(w http.ResponseWriter, r *http.Request) {
	account := c.Param("account")
	transaction, err := tc.TransactionService.GetTransactionsByAccount(c, account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, transaction)
}
