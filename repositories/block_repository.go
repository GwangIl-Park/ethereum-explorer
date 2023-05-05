package repositories

import (
	"context"
	"ethereum-explorer/models"

	"ethereum-explorer/db"
)

type blockRepository struct {
	db *db.DB
}

func NewBlockRepository(db *db.DB) models.BlockRepository {
	return &blockRepository {
		db,
	}
}

func (br *blockRepository) GetBlocks(c context.Context) ([]models.Block, error) {
	blocks, err := db.DB.ReadDocument("block", "", "")
	if err != nil {
		return nil, err
	}
	return blocks, nil
}

func (br *blockRepository) GetBlockByHeight(c context.Context, height uint) (models.Block, error) {
	blocks, err := db.DB.ReadDocument("block", "height", height)
	if err != nil {
		return nil, err
	}
	return blocks, nil
}
