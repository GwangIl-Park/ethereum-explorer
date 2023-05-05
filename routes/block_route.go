package routes

import (
	"ethereum-explorer/controller"
	"ethereum-explorer/repositories"
	"ethereum-explorer/usecases"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)
func NewBlockRouter(timeout time.Duration, db *mongo.Client, group *gin.RouterGroup) {
	br := repositories.NewBlockRepository(db)
	bc := &controller.BlockController{
		BlockUsecase: usecases.NewBlockUsecase(br, timeout),
	}
	group.GET("/blocks", bc.GetBlocks)
	group.GET("/block/:height", bc.GetBlockByHeight)
}