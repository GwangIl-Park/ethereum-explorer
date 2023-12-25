package router

import (
	"ethereum-explorer/controller"
	"net/http"
	"time"
)

func NewMainRouter(timeout time.Duration, mainController controller.MainController, router *http.ServeMux) {
	router.HandleFunc("/", mainController.GetAccountByGetInformationForMainAddress)
}
