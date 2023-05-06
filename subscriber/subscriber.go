package subscriber

import (
	"context"
	dbPackage "ethereum-explorer/db"
	"ethereum-explorer/ethClient"
	"ethereum-explorer/models"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
)

func insertDocument(block *types.Block, db *dbPackage.DB) {
		err := db.InsertOneDocument("blocks", 
			models.Block{
				BlockHeight:block.Number().String(),
				Receipient:block.Coinbase().String(),
				Reward:block.BaseFee().String(),
				Size:strconv.FormatUint(block.Size(), 10),
				GasUsed:strconv.FormatUint(block.GasUsed(), 10),
				Hash:block.Hash().String(),
			})
		if err != nil{}
		
		transactions := block.Transactions()

		var documents dbPackage.Documents

		for _, transaction := range transactions {
			msg, err := core.TransactionToMessage(transaction, types.LatestSignerForChainID(transaction.ChainId()), nil)
			if err != nil {}
			documents = append(documents,
				models.Transaction{
					Hash:transaction.Hash().String(),
					BlockHeight:block.Number().String(),
					From:msg.From.String(),
					To:transaction.To().String(),
					Value:transaction.Value().String(),
					TxFee:transaction.Cost().String(),
			})
		}
		
		err = db.InsertManyDocument("transactions", documents)
		if err != nil {}
}

type Subscriber struct {
	sub ethereum.Subscription
	header chan *types.Header
}

func NewSubscriber(ethClient *ethClient.EthClient, db *dbPackage.DB) (*Subscriber, *big.Int) {
	headers := make(chan *types.Header)
	sub, err := ethClient.Eth.SubscribeNewHead(context.Background(), headers)
	if err != nil {

	}

	header := <-headers
	block, err := ethClient.Eth.BlockByHash(context.Background(), header.Hash())
	if err != nil {
	}
	go insertDocument(block, db)
	return &Subscriber{
		sub,
		headers,
	}, block.Number()
}

func (sub *Subscriber) ProcessSubscribe(ethClient *ethClient.EthClient, db *dbPackage.DB) {
	for {
		select {
		case header:= <-sub.header:
			block, err := ethClient.Eth.BlockByHash(context.Background(), header.Hash())
			if err != nil {
			}
			go insertDocument(block, db)
		case err := <-sub.sub.Err():
			fmt.Println(err)
			return
		}
	}
}

func (sub *Subscriber) ProcessPrevious(ethClient *ethClient.EthClient, db *dbPackage.DB, start *big.Int, initBlock *big.Int) {
	one := big.NewInt(1)
	for i := start; i.Cmp(initBlock) > 0; i.Add(i, one) {
		block, err := ethClient.Eth.BlockByNumber(context.Background(), i)
		if err != nil {}
		insertDocument(block, db)
	}
}