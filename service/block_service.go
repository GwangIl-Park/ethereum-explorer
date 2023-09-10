package service

import (
	"ethereum-explorer/model"
	"ethereum-explorer/repository"
	"net/http"
)

type BlockService interface {
	GetBlocks(r *http.Request) ([]model.Block, error)
	GetBlockByHeight(r *http.Request) (model.Block, error)
}

type blockService struct {
	blockRepository repository.BlockRepository
}

func NewBlockService(blockRepository repository.BlockRepository) BlockService {
	return &blockService{
		blockRepository,
	}
}

func (bs *blockService) GetBlocks(r *http.Request) ([]model.Block, error) {
	return bs.blockRepository.GetBlocks()
}

func (bs *blockService) GetBlockByHeight(r *http.Request) (model.Block, error) {
	height := r.RequestURI[len("/block/"):]
	return bs.blockRepository.GetBlockByHeight(height)
}
