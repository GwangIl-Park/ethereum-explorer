package controller

import (
	"net/http"

	"ethereum-explorer/model"

	"github.com/gin-gonic/gin"
)

type BlockController struct {
	BlockService model.BlockUseCase
}

func (bc *BlockController) CreateBlock(c *gin.Context) {
	var block model.Block

	err := c.ShouldBind(&block)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResponse{Message: err.Error()})
		return
	}
}

func (bc *BlockController) GetBlocks(w http.ResponseWriter, r *http.Request) {
	blocks, err := bc.BlockService.GetBlocks(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, blocks)
}

func (bc *BlockController) GetBlockByHeight(w http.ResponseWriter, r *http.Request) {
	block, err := bc.BlockService.GetBlockByHeight(c, c.Param("height"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, block)
}
