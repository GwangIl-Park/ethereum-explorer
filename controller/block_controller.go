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
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message:err.Error()})
		return
	}
}

func (bc *BlockController) GetBlocks(c *gin.Context) {
	blocks, err := bc.BlockUsecase.GetBlocks(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, blocks)
}

func (bc *BlockController) GetBlockByHeight(c *gin.Context) {
	block, err := bc.BlockUsecase.GetBlockByHeight(c, c.Param("height"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, block)
}