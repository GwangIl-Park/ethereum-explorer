package service

import (
	"ethereum-explorer/dto"
	"ethereum-explorer/model"
	"ethereum-explorer/repository"

	"github.com/ethereum/go-ethereum/core/types"
)

type BlockService interface {
	GetBlocks() (*dto.GetBlocksDTO, error)
	GetBlockHeights() (*dto.GetBlockHeightsDTO, error)
	GetBlockByHeight(height string) (*dto.GetBlockByHeightDTO, error)
	CreateBlock(block *model.Block) error
	CreateBlocks(blocks []*model.Block) error
	CreateBlocksFromCoreBlocks(blocks []*types.Block) error
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

func (bs *blockService) CreateBlock(block *model.Block) error {
	return bs.blockRepository.CreateBlock(block)
}

func (bs *blockService) CreateBlocks(blocks []*model.Block) error {
	return bs.blockRepository.CreateBlocks(blocks)
}

func (bs *blockService) CreateBlocksFromCoreBlocks(blocks []*types.Block) error {
	var blockmodel []*model.Block
	for _, block := range blocks {
		blockmodel = append(blockmodel, model.MakeBlockModelFromTypes(block))
	}
	return bs.blockRepository.CreateBlocks(blockmodel)
}
