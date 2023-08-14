package subscriber

import (
	"context"
	dbPackage "ethereum-explorer/db"
	"ethereum-explorer/ethClient"
	"ethereum-explorer/logger"
	"ethereum-explorer/models"
	"ethereum-explorer/repositories"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
)

type Subscriber struct {
	sub       ethereum.Subscription
	header    chan *types.Header
	br        models.BlockRepository
	tr        models.TransactionRepository
	errorChan chan error
}

func makeBlockModel(block *types.Block) *models.Block {
	return &models.Block{
		BlockHeight: block.Number().String(),
		Receipient:  block.Coinbase().String(),
		Reward:      block.BaseFee().String(),
		Size:        strconv.FormatUint(block.Size(), 10),
		GasUsed:     strconv.FormatUint(block.GasUsed(), 10),
		Hash:        block.Hash().String()}
}

func makeTransactionModel(transaction *types.Transaction, height *big.Int) (*models.Transaction, error) {
	msg, err := core.TransactionToMessage(transaction, types.LatestSignerForChainID(transaction.ChainId()), nil)
	if err != nil {
		return nil, err
	}
	return &models.Transaction{
		Hash:        transaction.Hash().String(),
		BlockHeight: height.String(),
		From:        msg.From.String(),
		To:          transaction.To().String(),
		Value:       transaction.Value().String(),
		TxFee:       transaction.Cost().String(),
	}, nil
}

func (sub *Subscriber) insertBlockDocument(block *types.Block, db *dbPackage.DB, errorChan chan error) {
	ctx := context.Background()

	err := sub.br.CreateBlock(ctx, makeBlockModel(block))
	if err != nil {
		errorChan <- err
		return
	}

	transactions := block.Transactions()
	if transactions.Len() > 0 {
		var documents []*models.Transaction

		for _, transaction := range transactions {
			document, err := makeTransactionModel(transaction, block.Number())
			if err != nil {
				errorChan <- err
				return
			}
			documents = append(documents, document)
		}

		err = sub.tr.CreateTransactions(ctx, documents)
		if err != nil {
			errorChan <- err
			return
		}
	}
}

func (sub *Subscriber) insertPreviousDocuments(blocks []*types.Block, db *dbPackage.DB) error {
	var blockDocuments []*models.Block
	var transactionDocuments []*models.Transaction
	ctx := context.Background()

	if len(blocks) > 0 {
		for _, block := range blocks {
			blockDocuments = append(blockDocuments, makeBlockModel(block))
			transactions := block.Transactions()

			if transactions.Len() > 0 {
				for _, transaction := range transactions {
					document, err := makeTransactionModel(transaction, block.Number())
					if err != nil {
						return err
					}
					transactionDocuments = append(transactionDocuments, document)
				}
			}
		}
		err := sub.br.CreateBlocks(ctx, blockDocuments)
		if err != nil {
			return err
		}
		if len(transactionDocuments) > 0 {
			err = sub.tr.CreateTransactions(ctx, transactionDocuments)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func NewSubscriber(ethClient *ethClient.EthClient, db *dbPackage.DB, errorChan chan error) (*Subscriber, *big.Int, error) {
	br := repositories.NewBlockRepository(db)
	tr := repositories.NewTransactionRepository(db)

	headers := make(chan *types.Header)

	logger.Logger.Info("Subscribe New Head")

	sub, err := ethClient.Ws.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		return nil, nil, err
	}

	logger.Logger.Info("Waiting New Block")

	header := <-headers
	block, err := ethClient.Http.BlockByHash(context.Background(), header.Hash())
	if err != nil {
		return nil, nil, err
	}

	return &Subscriber{
		sub,
		headers,
		br,
		tr,
		errorChan,
	}, block.Number(), nil
}

func (sub *Subscriber) ProcessSubscribe(ethClient *ethClient.EthClient, db *dbPackage.DB) {
	for {
		select {
		case header := <-sub.header:
			block, err := ethClient.Http.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				sub.errorChan <- err
				return
			}
			go sub.insertBlockDocument(block, db, sub.errorChan)
		case err := <-sub.sub.Err():
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
	err = sub.insertPreviousDocuments(blocks, db)
	if err != nil {
		sub.errorChan <- err
	}
}
