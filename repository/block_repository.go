package repository

import (
	"ethereum-explorer/dto"
	"ethereum-explorer/model"
	"fmt"

	"ethereum-explorer/db"
)

type BlockRepository interface {
	GetBlocks() (*dto.GetBlocksDTO, error)
	GetBlockHeights() (*dto.GetBlockHeightsDTO, error)
	GetBlockByHeight(height string) (*dto.GetBlockByHeightDTO, error)
	CreateBlock(block *model.Block) error
	CreateBlocks(blocks []*model.Block) error
}

type blockRepository struct {
	db *db.DB
}

func NewBlockRepository(db *db.DB) BlockRepository {
	return &blockRepository{
		db,
	}
}

func (br *blockRepository) GetBlocks() (*dto.GetBlocksDTO, error) {
	rows, err := br.db.Client.Query(`SELECT * FROM "Block"`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var getBlocksDTO *dto.GetBlocksDTO
	for rows.Next() {
		var block dto.Block
		err = rows.Scan(&block)
		if err != nil {
			return nil, err
		}
		getBlocksDTO.Blocks = append(getBlocksDTO.Blocks, block)
	}

	return getBlocksDTO, nil
}

func (br *blockRepository) GetBlockHeights() (*dto.GetBlockHeightsDTO, error) {
	rows, err := br.db.Client.Query(`SELECT blockHeight FROM "Block"`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var getBlockHeightsDTO *dto.GetBlockHeightsDTO
	for rows.Next() {
		var blockHeight string
		err = rows.Scan(&blockHeight)
		if err != nil {
			return nil, err
		}
		getBlockHeightsDTO.Heights = append(getBlockHeightsDTO.Heights, blockHeight)
	}

	return getBlockHeightsDTO, nil
}

func (br *blockRepository) GetBlockByHeight(height string) (*dto.GetBlockByHeightDTO, error) {
	rows, err := br.db.Client.Query(`SELECT * FROM "Block" WHERE blockHeight = %s`, height)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var getBlockByHeightDTO *dto.GetBlockByHeightDTO
	for rows.Next() {
		err = rows.Scan(&getBlockByHeightDTO.Block)
		if err != nil {
			return nil, err
		}
	}

	return getBlockByHeightDTO, nil
}

func (br *blockRepository) CreateBlock(block *model.Block) error {

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

func (br *blockRepository) CreateBlocks(blocks []*model.Block) error {
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
