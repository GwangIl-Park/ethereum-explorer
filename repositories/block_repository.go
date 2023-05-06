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
	documents, err := br.db.ReadDocuments("blocks", "", "")
	if err != nil {
		return nil, err
	}

	var blocks []models.Block
	for _, document := range documents {
		blocks = append(blocks, document.(models.Block))
	}
	
	return blocks, nil
}

func (br *blockRepository) GetBlockByHeight(c context.Context, height string) (models.Block, error) {
	document := br.db.ReadDocument("blocks", "height", height)

	block := document.(models.Block)
	
	return block, nil
}
