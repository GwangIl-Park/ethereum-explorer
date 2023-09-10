package model

import (
	"github.com/ethereum/go-ethereum/core/types"
)

type Block struct {
	BlockHeight string `json:"blockHeight"`
	//Status      bool   `json:"status"`
	Timestamp  uint64 `json:"timestamp"`
	Receipient string `json:"receipient"`
	//Reward
	Size       uint64 `json:"size"`
	GasUsed    uint64 `json:"gasUsed"`
	GasLimit   uint64 `json:"gasLimit"`
	BaseFee    string `json:"baseFee"`
	ExtraData  string `json:"extraData"`
	Hash       string `json:"hash"`
	ParentHash string `json:"parentHash"`
}

func MakeBlockModelFromTypes(block *types.Block) *Block {
	return &Block{
		BlockHeight: block.Number().String(),
		Timestamp:   block.Header().Time,
		Receipient:  block.Coinbase().String(),
		Size:        block.Size(),
		GasUsed:     block.GasUsed(),
		GasLimit:    block.GasLimit(),
		BaseFee:     block.BaseFee().String(),
		ExtraData:   string(block.Extra()),
		Hash:        block.Hash().String(),
		ParentHash:  string(block.ParentHash().String()),
	}
}
