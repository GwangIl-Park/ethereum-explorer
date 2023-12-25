package controller

import (
	"ethereum-explorer/httpResponse"
	"ethereum-explorer/logger"
	"ethereum-explorer/service"
	"net/http"

	"ethereum-explorer/model"

	"github.com/gin-gonic/gin"
)

type BlockController struct {
	BlockService service.BlockService
}

func NewBlockController(blockService service.BlockService) BlockController {
	return BlockController{
		blockService,
	}
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
	if r.Method != "GET" {
		logger.LogMethodNotAllowed(r)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	responseData, err := bc.BlockService.GetBlocks()
	if err != nil {
		httpResponse.ErrorResponse(w, r, err)
		return
	}

	httpResponse.SendResponse(w, responseData)
}

func (bc *BlockController) GetBlockByHeight(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		logger.LogMethodNotAllowed(r)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	height := r.RequestURI[len("/block/"):]

	responseData, err := bc.BlockService.GetBlockByHeight(height)
	if err != nil {
		httpResponse.ErrorResponse(w, r, err)
		return
	}

	httpResponse.SendResponse(w, responseData)
}
