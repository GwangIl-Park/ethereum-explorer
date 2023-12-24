package controller

import "ethereum-explorer/service"

type MainController struct {
	MainService service.MainService
}

func NewMainController(mainService service.MainService) MainController {
	return MainController{
		mainService,
	}
}