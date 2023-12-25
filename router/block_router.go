package router

import (
	"ethereum-explorer/controller"
	"net/http"
)

func NewBlockRouter(blockController controller.BlockController, router *http.ServeMux) {
	router.HandleFunc("/blocks", blockController.GetBlocks)
	router.HandleFunc("/block/:height", blockController.GetBlockByHeight)
}
