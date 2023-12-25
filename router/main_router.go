package router

import (
	"ethereum-explorer/controller"
	"net/http"
)

func NewMainRouter(mainController controller.MainController, router *http.ServeMux) {
	router.HandleFunc("/", mainController.GetAccountByGetInformationForMainAddress)
}
