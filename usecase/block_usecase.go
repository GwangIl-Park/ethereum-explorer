package usecase

import (
	"context"
	"ethereum-explorer/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type blockUsecase struct {
	blockRepository model.BlockRepository
	contextTimeout  time.Duration
}

func NewBlockUsecase(blockRepository model.BlockRepository, timeout time.Duration) model.BlockUseCase {
	return &blockUsecase{
		blockRepository,
		timeout,
	}
}

func (bu *blockUsecase) GetBlocks(c *gin.Context) ([]model.Block, error) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		return nil, err
	}
	show, err := strconv.Atoi(c.DefaultQuery("show", "10"))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(c, bu.contextTimeout)
	defer cancel()
	return bu.blockRepository.GetBlocks(ctx, int64(page), int64(show))
}

func (bu *blockUsecase) GetBlockByHeight(c *gin.Context, height string) (model.Block, error) {
	ctx, cancel := context.WithTimeout(c, bu.contextTimeout)
	defer cancel()
	return bu.blockRepository.GetBlockByHeight(ctx, height)
}