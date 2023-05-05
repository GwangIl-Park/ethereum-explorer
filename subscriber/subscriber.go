package subscriber

import (
	"context"
	"ethereum-explorer/ethClient"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
)

type Subscriber struct {
	sub ethereum.Subscription
	header chan *types.Header
}

func NewSubscriber(ethClient *ethClient.EthClient) *Subscriber {
	headers := make(chan *types.Header)
	sub, err := ethClient.Eth.SubscribeNewHead(context.Background(), headers)
	if err != nil {

	}
	return &Subscriber{
		sub,
		headers,
	}
}

func (sub *Subscriber) ProcessSubscribe(ethClient *ethClient.EthClient) {
	for {
		select {
		case header:= <-sub.header:
			block, err := ethClient.Eth.BlockByHash(context.Background(), header.Hash())
			if err != nil {
			}
			go func() {
				db.QueryRow(
					"Insert into blocks (blockHeight, receipient, reward, size, gasUsed, hash) values ( ? ? ? ? ? ? )", 
					block.Number(),
					block.Coinbase(),
					block.Coinbase(),
					block.Size(),
					block.GasUsed(),
					block.Hash(),
				)
			}
			
			tr.db.QueryRow("Insert into transactions (id, hash, blockHeight, from, to, value, txFee) values (?, ?, ?, ?, ?, ?, ?)", transaction.Id, transaction.Hash, transaction.BlockHeight, transaction.From, transaction.To, transaction.Value, transaction.TxFee)
			return
		case err := <-sub.sub.Err():
			return
		}
	}
}