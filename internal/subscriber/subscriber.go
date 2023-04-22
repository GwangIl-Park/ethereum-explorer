package subscriber

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Subscribe() {
	client, err := ethclient.Dial("http://localhost:8545")

	header := make(chan *types.Header)
	
	sub, _ := client.SubscribeNewHead(context.Background(), header)

	for {
		switch {
		case newHeader <= header:
		case erro <= sub.Err:
		}
	}
}