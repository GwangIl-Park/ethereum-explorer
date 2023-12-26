package subscriber

import (
	"context"
	dbPackage "ethereum-explorer/db"
	"ethereum-explorer/ethClient"
	"ethereum-explorer/service"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

type Subscriber struct {
	headerChan chan *types.Header
	ethClient  *ethClient.EthClient
	bs         service.BlockService
	ts         service.TransactionService
	errorChan  chan error
}

func NewSubscriber(ethClient *ethClient.EthClient, blockService service.BlockService, transactionService service.TransactionService, errorChan chan error) (*Subscriber, error) {
	headerChan := make(chan *types.Header)

	return &Subscriber{
		headerChan,
		ethClient,
		blockService,
		transactionService,
		errorChan,
	}, nil
}

func (sub *Subscriber) insertNewBlocks(blocks []*types.Block) {
	sub.bs.CreateBlocksFromCoreBlocks(blocks)
	sub.ts.CreateTransactionsFromCoreBlocks(sub.ethClient, blocks)
}

func (sub *Subscriber) ProcessSubscribe(ethClient *ethClient.EthClient, initBlockNumberChan chan *big.Int) {
	subscription, err := ethClient.Ws.SubscribeNewHead(context.Background(), sub.headerChan)
	if err != nil {
		sub.errorChan <- err
	}

	header := <-sub.headerChan
	initBlock, err := ethClient.Http.BlockByHash(context.Background(), header.Hash())
	if err != nil {
		sub.errorChan <- err
	}

	initBlockNumberChan <- initBlock.Number()

	for {
		select {
		case header := <-sub.headerChan:
			block, err := ethClient.Http.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				sub.errorChan <- err
				return
			}
			go sub.insertNewBlocks([]*types.Block{block})
		case err := <-subscription.Err():
			sub.errorChan <- err
			return
		}
	}
}

func (sub *Subscriber) ProcessPrevious(ethClient *ethClient.EthClient, db *dbPackage.DB, initBlock *big.Int) {
	bigZero := big.NewInt(0)
	bigOne := big.NewInt(1)
	var blocks []*types.Block

	blockHeightsDTO, err := sub.bs.GetBlockHeights()
	if err != nil {
		sub.errorChan <- err
	}

	blockHeights := blockHeightsDTO.Heights
	blockHeightsMap := make(map[*big.Int]bool)
	for _, blockHeight := range blockHeights {
		num := new(big.Int)
		num.SetString(blockHeight, 10)
		blockHeightsMap[num] = true
	}

	for i := initBlock; blockHeightsMap[i] && i.Cmp(bigZero) < 0; i.Sub(i, bigOne) {
		block, err := ethClient.Http.BlockByNumber(context.Background(), i)
		if err != nil {
			sub.errorChan <- err
			return
		}
		blocks = append(blocks, block)
	}

	if len(blocks) > 0 {
		sub.insertNewBlocks(blocks)
	}
}
