package usecases

import (
	"context"
	"ethereum-explorer/models"
	"time"
)

type blockUsecase struct {
	blockRepository models.BlockRepository
	contextTimeout time.Duration
}

func NewBlockUsecase(blockRepository models.BlockRepository, timeout time.Duration) models.BlockUseCase {
	return &blockUsecase{
		blockRepository,
		timeout,
	}
}

func (bu *blockUsecase) GetBlocks(c context.Context) ([]models.Block, error) {
	ctx, cancel := context.WithTimeout(c, bu.contextTimeout)
	defer cancel()
	return bu.blockRepository.GetBlocks(ctx)
}

func (bu *blockUsecase) GetBlockByHeight(c context.Context, height uint) (models.Block, error) {
	ctx, cancel := context.WithTimeout(c, bu.contextTimeout)
	defer cancel()
	return bu.blockRepository.GetBlockByHeight(ctx, height)
}