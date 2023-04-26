package repositories

import (
	"context"
	"database/sql"
	"ethereum-explorer/models"
)

type blockRepository struct {
	db *sql.DB
}

func NewBlockRepository(db *sql.DB) models.BlockRepository {
	return &blockRepository {
		db,
	}
}

func (br *blockRepository) CreateBlock(c context.Context, block *models.Block) error{
	br.db.QueryRow("Insert into blocks (blockHeight, receipient, reward, size, gasUsed, hash) values ( ? ? ? ? ? ? )", block.BlockHeight, block.Receipient, block.Reward, block.Size, block.GasUsed, block.Hash)
	return nil
}

func (br *blockRepository) GetBlock(c context.Context) ([]models.Block, error) {
	rows, err := br.db.Query("SELECT * from blocks")
	if err != nil {

	}
	var blockHeight, size, gasUsed int
	var receipient, hash string
	var reward float32
	var blocks []models.Block
	for rows.Next() {
		err = rows.Scan(&blockHeight, &receipient, &reward, &size, &gasUsed, &hash)
		if err != nil {
			blocks = append(blocks, models.Block{BlockHeight: blockHeight, Receipient: receipient, Reward: reward, Size: size, GasUsed: gasUsed, Hash: hash})
		}
	}
	return blocks, nil
}

func (br *blockRepository) GetBlockByHeight(c context.Context, height uint) (models.Block, error) {
	var blockHeight, size, gasUsed int
	var receipient, hash string
	var reward float32
	br.db.QueryRow("SELECT * FROM blocks WHERE blockHeight=?", height).Scan(&blockHeight, &receipient, &reward, &size, &gasUsed, &hash)
	return models.Block{BlockHeight: blockHeight, Receipient: receipient, Reward: reward, Size: size, GasUsed: gasUsed, Hash: hash}, nil
}
