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

func (br *blockRepository) GetBlocks(c context.Context, page int64, show int64) ([]models.Block, error) {
	rows, err := br.db.Client.Query(`SELECT * FROM "Block"`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var blocks []models.Block
	for (rows.Next()) {
		var block models.Block
		err = rows.Scan(&block)
		if err != nil {
			panic(err)
		}
		blocks = append(blocks, block)
	}

	return blocks, nil
}

func (br *blockRepository) GetBlockByHeight(c context.Context, height string) (models.Block, error) {
	rows, err := br.db.Client.Query(`SELECT * FROM "Block" WHERE blockHeight = %s`, height)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var block models.Block
	err = rows.Scan(&block)
	if err != nil {
		panic(err)
	}

	return block, nil
}