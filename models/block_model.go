package models

import (
	"context"

	"github.com/gin-gonic/gin"
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
	GetBlocks(c context.Context, page int64, show int64) ([]Block, error)
	GetBlockByHeight(c context.Context, height string) (Block, error)
}

type BlockUseCase interface {
	GetBlocks(c *gin.Context) ([]Block, error)
	GetBlockByHeight(c *gin.Context, height string) (Block, error)
}