package controller

import (
	"encoding/json"
	"ethereum-explorer/logger"
	"ethereum-explorer/service"
	"net/http"
)

type AccountController struct {
	AccountService service.AccountService
}

func NewAccountController(accountService service.AccountService) AccountController {
	return AccountController{
		accountService,
	}
}

func (ac *AccountController) GetAccountByAddress(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		logger.LogMethodNotAllowed(r)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	account, err := ac.AccountService.GetAccountByAddress(r)
	if err != nil {
		logger.LogInternalServerError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(account)

	w.Write(jsonData)
}