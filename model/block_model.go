package model

import "context"

type Block struct {
	
}

type BlockRepository interface{
	CreateBlock(c context.Context, block *Block) error
	GetBlock(c context.Context) ([]Block, error)
	GetBlockByHeight(c context.Context, height uint) (Block, error)
}