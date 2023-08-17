package models

import (
	"context"
	"strconv"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gin-gonic/gin"
)

type Block struct {
	BlockHeight string `json:"blockHeight"`
	//Status      bool   `json:"status"`
	Timestamp  string `json:"timestamp"`
	Receipient string `json:"receipient"`
	Reward     string `json:"reward"`
	Size       string `json:"size"`
	GasUsed    string `json:"gasUsed"`
	GasLimit   string `json:"gasLimit"`
	BaseFee    string `json:"baseFee"`
	ExtraData  string `json:"extraData"`
	Hash       string `json:"hash"`
	ParentHash string `json:"parentHash"`
}

type BlockRepository interface {
	GetBlocks(c context.Context, page int64, show int64) ([]Block, error)
	GetBlockHeights(c context.Context) ([]string, error)
	GetBlockByHeight(c context.Context, height string) (Block, error)
	CreateBlock(c context.Context, block *Block) error
	CreateBlocks(c context.Context, blocks []*Block) error
}

type BlockUseCase interface {
	GetBlocks(c *gin.Context) ([]Block, error)
	GetBlockByHeight(c *gin.Context, height string) (Block, error)
}

func MakeBlockModelFromTypes(block *types.Block) *Block {
	return &Block{
		BlockHeight: block.Number().String(),
		Timestamp:   string(block.Header().Time),
		Receipient:  block.Coinbase().String(),
		Reward:      block.BaseFee().String(),
		Size:        strconv.FormatUint(block.Size(), 10),
		GasUsed:     strconv.FormatUint(block.GasUsed(), 10),
		GasLimit:    block.GasLimit(),
		BaseFee:     block.BaseFee().String(),
		ExtraData:   string(block.Extra()),
		Hash:        block.Hash().String(),
		ParentHash:  string(block.ParentHash().String()),
	}
}
