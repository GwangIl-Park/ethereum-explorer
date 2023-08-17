package subscriber

import (
	"context"
	dbPackage "ethereum-explorer/db"
	"ethereum-explorer/ethClient"
	"ethereum-explorer/models"
	"ethereum-explorer/repositories"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

type Subscriber struct {
	headerChan chan *types.Header
	ethClient  *ethClient.EthClient
	br         models.BlockRepository
	tr         models.TransactionRepository
	errorChan  chan error
}

func NewSubscriber(ethClient *ethClient.EthClient, db *dbPackage.DB, errorChan chan error) (*Subscriber, error) {
	br := repositories.NewBlockRepository(db)
	tr := repositories.NewTransactionRepository(db)

	headerChan := make(chan *types.Header)

	return &Subscriber{
		headerChan,
		ethClient,
		br,
		tr,
		errorChan,
	}, nil
}

func (sub *Subscriber) insertNewBlocks(blocks []*types.Block) {
	ctx := context.Background()

	var blockModels []*models.Block
	var transactionModels []*models.Transaction

	for _, block := range blocks {
		transactions := block.Transactions()
		if transactions.Len() > 0 {
			for _, transaction := range transactions {
				receipt, err := sub.ethClient.Http.TransactionReceipt(ctx, transaction.Hash())
				if err != nil {
					sub.errorChan <- err
					return
				}
				transactionModels = append(transactionModels, models.MakeTransactionModelFromTypes(receipt, *transaction, *block))
			}
		}
		blockModels = append(blockModels, models.MakeBlockModelFromTypes(block))
	}

	if len(transactionModels) > 0 {
		err := sub.tr.CreateTransactions(ctx, transactionModels)
		if err != nil {
			sub.errorChan <- err
			return
		}
	}

	err := sub.br.CreateBlocks(ctx, blockModels)
	if err != nil {
		sub.errorChan <- err
		return
	}
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
	ctx := context.Background()
	bigZero := big.NewInt(0)
	bigOne := big.NewInt(1)
	var blocks []*types.Block

	blockHeights, err := sub.br.GetBlockHeights(ctx)
	if err != nil {
		sub.errorChan <- err
	}

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
