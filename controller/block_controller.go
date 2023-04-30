package controller

import (
	"ethereum-explorer/models"
	"net/http"
	"strconv"

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
	height, _ := strconv.ParseUint(c.Param("height"), 10, 64)
	block, err := bc.BlockUsecase.GetBlockByHeight(c, uint(height))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, block)
}