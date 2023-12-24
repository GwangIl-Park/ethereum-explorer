package router

import (
	"ethereum-explorer/controller"
	"net/http"
	"time"
)

func NewAccountRouter(timeout time.Duration, accountController controller.AccountController, router *http.ServeMux) {
}
