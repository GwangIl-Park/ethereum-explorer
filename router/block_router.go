package router

import (
	"ethereum-explorer/controller"
	"net/http"
	"time"
)

func NewBlockRouter(timeout time.Duration, blockController controller.BlockController, router *http.ServeMux) {
	router.HandleFunc("/blocks", blockController.GetBlocks)
	router.HandleFunc("/block/:height", blockController.GetBlockByHeight)
}
