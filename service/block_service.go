package service

import (
	"ethereum-explorer/model"
	"ethereum-explorer/repository"
	"net/http"
	"time"
)

type BlockService interface {
	GetBlocks(r *http.Request) ([]model.Block, error)
	GetBlockByHeight(r *http.Request, height string) (model.Block, error)
}

type blockService struct {
	blockRepository repository.BlockRepository
	contextTimeout  time.Duration
}

func NewBlockService(blockRepository repository.BlockRepository, timeout time.Duration) BlockService {
	return &blockService{
		blockRepository,
		timeout,
	}
}

func (bu *blockService) GetBlocks(r *http.Request) ([]model.Block, error) {
	// page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	// if err != nil {
	// 	return nil, err
	// }
	// show, err := strconv.Atoi(c.DefaultQuery("show", "10"))
	// if err != nil {
	// 	return nil, err
	// }

	// ctx, cancel := context.WithTimeout(c, bu.contextTimeout)
	// defer cancel()
	// return bu.blockRepository.GetBlocks(ctx, int64(page), int64(show))
	return nil, nil
}

func (bu *blockService) GetBlockByHeight(r *http.Request, height string) (model.Block, error) {
	// ctx, cancel := context.WithTimeout(c, bu.contextTimeout)
	// defer cancel()
	// return bu.blockRepository.GetBlockByHeight(ctx, height)
	return model.Block{}, nil
}
