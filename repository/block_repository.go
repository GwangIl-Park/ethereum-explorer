package repository

import (
	"context"
	"ethereum-explorer/model"
	"fmt"

	"ethereum-explorer/db"
)

type BlockRepository interface {
	GetBlocks(c context.Context, page int64, show int64) ([]model.Block, error)
	GetBlockHeights(c context.Context) ([]string, error)
	GetBlockByHeight(c context.Context, height string) (model.Block, error)
	CreateBlock(c context.Context, block *model.Block) error
	CreateBlocks(c context.Context, blocks []*model.Block) error
}

type blockRepository struct {
	db *db.DB
}

func NewBlockRepository(db *db.DB) BlockRepository {
	return &blockRepository{
		db,
	}
}

func (br *blockRepository) GetBlocks(c context.Context, page int64, show int64) ([]model.Block, error) {
	rows, err := br.db.Client.Query(`SELECT * FROM "Block"`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var blocks []model.Block
	for rows.Next() {
		var block model.Block
		err = rows.Scan(&block)
		if err != nil {
			panic(err)
		}
		blocks = append(blocks, block)
	}

	return blocks, nil
}

func (br *blockRepository) GetBlockHeights(c context.Context) ([]string, error) {
	rows, err := br.db.Client.Query(`SELECT blockHeight FROM "Block"`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var blockHeights []string
	for rows.Next() {
		var blockHeight string
		err = rows.Scan(&blockHeight)
		if err != nil {
			panic(err)
		}
		blockHeights = append(blockHeights, blockHeight)
	}

	return blockHeights, nil
}

func (br *blockRepository) GetBlockByHeight(c context.Context, height string) (model.Block, error) {
	rows, err := br.db.Client.Query(`SELECT * FROM "Block" WHERE blockHeight = %s`, height)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var block model.Block
	err = rows.Scan(&block)
	if err != nil {
		panic(err)
	}

	return block, nil
}

func (br *blockRepository) CreateBlock(c context.Context, block *model.Block) error {

	valuesStr := fmt.Sprintf("(%s,%v,%s,%v,%v,%v,%s,%s,%s,%s)",
		block.BlockHeight,
		block.Timestamp,
		block.Receipient,
		block.Size,
		block.GasUsed,
		block.GasLimit,
		block.BaseFee,
		block.ExtraData,
		block.Hash,
		block.ParentHash,
	)

	_, err := br.db.Client.Exec(`INSERT INTO Block VALUES %s`, valuesStr)
	if err != nil {
		return err
	}
	return nil
}

func (br *blockRepository) CreateBlocks(c context.Context, blocks []*model.Block) error {
	var valuesStr string

	for _, block := range blocks {
		valueStr := fmt.Sprintf("(%s,%v,%s,%v,%v,%v,%s,%s,%s,%s)",
			block.BlockHeight,
			block.Timestamp,
			block.Receipient,
			block.Size,
			block.GasUsed,
			block.GasLimit,
			block.BaseFee,
			block.ExtraData,
			block.Hash,
			block.ParentHash,
		)
		valuesStr = fmt.Sprintf("%s%s", valuesStr, valueStr)
	}

	_, err := br.db.Client.Exec(`INSERT INTO Block VALUES %s`, valuesStr[:len(valuesStr)-1])
	if err != nil {
		return err
	}
	return nil
}
