package models

import "context"

type Block struct {
	BlockHeight int `json:"blockHeight"`
	Receipient string `json:"receipient"`
	Reward float32 `json:"reward"`
	Size int `json:"size"`
	GasUsed int `json:"gasUsed"`
	Hash string `json:"hash"`
}

type BlockRepository interface{
	CreateBlock(c context.Context, block *Block) error
	GetBlock(c context.Context) ([]Block, error)
	GetBlockByHeight(c context.Context, height uint) (Block, error)
}