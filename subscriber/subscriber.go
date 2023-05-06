package subscriber

import (
	"context"
	"ethereum-explorer/db"
	"ethereum-explorer/ethClient"
	"ethereum-explorer/models"
	"fmt"
	"log"
	"strconv"

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

func (sub *Subscriber) ProcessSubscribe(ethClient *ethClient.EthClient, database *db.DB) {
	for {
		select {
		case header:= <-sub.header:
			block, err := ethClient.Eth.BlockByHash(context.Background(), header.Hash())
			if err != nil {
			}
			go func() {
				err := database.InsertOneDocument("blocks", 
					models.Block{
						block.Number().String(),
						block.Coinbase().String(),
						block.BaseFee().String(),
						strconv.FormatUint(block.Size(), 10),
						strconv.FormatUint(block.GasUsed(), 10),
						block.Hash().String(),
					})
				if err != nil{}
				
				transactions := block.Transactions()

				var documents db.Documents

				chainID, err := ethClient.NetworkID(context.Background())
        if err != nil {
            log.Fatal(err)
        }

        if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID)); err == nil {
            fmt.Println(msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
        }

				for _, transaction := range transactions {
					documents = append(documents,
						models.Transaction{
							transaction.Hash().String(),
							block.Number().String(),
							transaction.
							transaction.To().String(),
							transaction.Value().String(),
							transaction.Value(),
					})
				}
				
				err = database.InsertManyDocument("transactions", documents)
				if err != nil {}

			}()
			return
		case err := <-sub.sub.Err():
			return
		}
	}
}