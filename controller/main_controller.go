package controller

import (
	"ethereum-explorer/httpResponse"
	"ethereum-explorer/logger"
	"ethereum-explorer/service"
	"net/http"
)

type MainController struct {
	MainService service.MainService
}

func NewMainController(mainService service.MainService) MainController {
	return MainController{
		mainService,
	}
}

func (mc *MainController) GetAccountByGetInformationForMainAddress(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		logger.LogMethodNotAllowed(r)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	responseData, err := mc.MainService.GetInformationForMain()
	if err != nil {
		httpResponse.ErrorResponse(w, r, err)
		return
	}

	httpResponse.SendResponse(w, responseData)
}