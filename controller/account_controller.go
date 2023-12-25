package controller

import (
	"ethereum-explorer/httpResponse"
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

	address := r.RequestURI[len("/address/"):]

	responseData, err := ac.AccountService.GetAccountByAddress(address)
	if err != nil {
		httpResponse.ErrorResponse(w, r, err)
		return
	}

	httpResponse.SendResponse(w, responseData)
}