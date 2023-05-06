package routes

import (
	"ethereum-explorer/controller"
	"ethereum-explorer/db"
	"ethereum-explorer/repositories"
	"ethereum-explorer/usecases"
	"time"

	"github.com/gin-gonic/gin"
)
func NewBlockRouter(timeout time.Duration, db *db.DB, group *gin.RouterGroup) {
	br := repositories.NewBlockRepository(db)
	bc := &controller.BlockController{
		BlockUsecase: usecases.NewBlockUsecase(br, timeout),
	}
	group.GET("/blocks", bc.GetBlocks)
	group.GET("/block/:height", bc.GetBlockByHeight)
}