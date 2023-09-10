package controller

import (
	"encoding/json"
	"ethereum-explorer/logger"
	"ethereum-explorer/service"
	"net/http"

	"ethereum-explorer/model"

	"github.com/gin-gonic/gin"
)

type BlockController struct {
	BlockService service.BlockService
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

	blocks, err := bc.BlockService.GetBlocks(r)
	if err != nil {
		logger.LogInternalServerError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(blocks)

	w.Write(jsonData)
}

func (bc *BlockController) GetBlockByHeight(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		logger.LogMethodNotAllowed(r)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	block, err := bc.BlockService.GetBlockByHeight(r)
	if err != nil {
		logger.LogInternalServerError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonData, _ := json.Marshal(block)

	w.Write(jsonData)
}
