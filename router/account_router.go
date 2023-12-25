package router

import (
	"ethereum-explorer/controller"
	"net/http"
)

func NewAccountRouter(accountController controller.AccountController, router *http.ServeMux) {
	router.HandleFunc("/address/", accountController.GetAccountByAddress)
}
