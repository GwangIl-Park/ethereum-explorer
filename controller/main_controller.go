package controller

import (
	"encoding/json"
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

	main, err := mc.MainService.GetInformationForMain(r)
	if err != nil {
		logger.LogInternalServerError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(main)

	w.Write(jsonData)
}