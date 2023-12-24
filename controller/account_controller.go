package controller

import "ethereum-explorer/service"

type AccountController struct {
	AccountService service.AccountService
}

func NewAccountController(accountService service.AccountService) AccountController {
	return AccountController{
		accountService,
	}
}