package routes

import (
	"ethereum-explorer/controller"
	"ethereum-explorer/db"
	"ethereum-explorer/repositories"
	"ethereum-explorer/usecases"
	"net/http"
	"time"
)

func NewBlockRouter(timeout time.Duration, db *db.DB, router *http.ServeMux) {
	br := repositories.NewBlockRepository(db)
	bc := &controller.BlockController{
		BlockUsecase: usecases.NewBlockUsecase(br, timeout),
	}
	router.HandleFunc("/blocks", bc.GetBlocks)
	router.HandleFunc("/block/:height", bc.GetBlockByHeight)
}
