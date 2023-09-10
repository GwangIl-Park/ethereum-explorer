package router

import (
	"ethereum-explorer/controller"
	"ethereum-explorer/db"
	"ethereum-explorer/repository"
	"ethereum-explorer/service"
	"net/http"
	"time"
)

func NewBlockRouter(timeout time.Duration, db *db.DB, router *http.ServeMux) {
	br := repository.NewBlockRepository(db)
	bc := &controller.BlockController{
		BlockService: service.NewBlockService(br),
	}
	router.HandleFunc("/blocks", bc.GetBlocks)
	router.HandleFunc("/block/:height", bc.GetBlockByHeight)
}
