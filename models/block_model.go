package models

import (
	"context"
)

type Block struct {
	BlockHeight string `json:"blockHeight"`
	Receipient string `json:"receipient"`
	Reward string `json:"reward"`
	Size string `json:"size"`
	GasUsed string `json:"gasUsed"`
	Hash string `json:"hash"`
}

type BlockRepository interface {
	GetBlocks(c context.Context) ([]Block, error)
	GetBlockByHeight(c context.Context, height string) (Block, error)
}

type BlockUseCase interface {
	GetBlocks(c context.Context) ([]Block, error)
	GetBlockByHeight(c context.Context, height string) (Block, error)
}