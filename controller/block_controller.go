package controller

import (
	"ethereum-explorer/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlockController struct {
	BlockUsecase models.BlockUseCase
}

func (bc *BlockController) CreateBlock(c *gin.Context) {
	var block models.Block

	err := c.ShouldBind(&block)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{Message:err.Error()})
		return
	}
}