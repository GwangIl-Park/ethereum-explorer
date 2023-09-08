package controller

import (
	"ethereum-explorer/model"
	"net/http"
)

type TransactionController struct {
	TransactionUsecase model.TransactionUsecase
}

func (tc *TransactionController) GetTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := tc.TransactionUsecase.GetTransactions(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, transactions)
}

func (tc *TransactionController) GetTransactionByHash(w http.ResponseWriter, r *http.Request) {
	hash := c.Param("hash")
	transaction, err := tc.TransactionUsecase.GetTransactionByHash(c, hash)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, transaction)
}

func (tc *TransactionController) GetTransactionsByAccount(w http.ResponseWriter, r *http.Request) {
	account := c.Param("account")
	transaction, err := tc.TransactionUsecase.GetTransactionsByAccount(c, account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, transaction)
}
