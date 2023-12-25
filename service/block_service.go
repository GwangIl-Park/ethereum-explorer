package service

import (
	"ethereum-explorer/dto"
	"ethereum-explorer/repository"
)

type BlockService interface {
	GetBlocks() (*dto.GetBlocksDTO, error)
	GetBlockHeights() (*dto.GetBlockHeightsDTO, error)
	GetBlockByHeight(height string) (*dto.GetBlockByHeightDTO, error)
}

type blockService struct {
	blockRepository repository.BlockRepository
}

func NewBlockService(blockRepository repository.BlockRepository) BlockService {
	return &blockService{
		blockRepository,
	}
}

func (bs *blockService) GetBlocks() (*dto.GetBlocksDTO, error) {
	return bs.blockRepository.GetBlocks()
}

func (bs *blockService) GetBlockHeights() (*dto.GetBlockHeightsDTO, error) {
	return bs.blockRepository.GetBlockHeights()
}

func (bs *blockService) GetBlockByHeight(height string) (*dto.GetBlockByHeightDTO, error) {
	return bs.blockRepository.GetBlockByHeight(height)
}
