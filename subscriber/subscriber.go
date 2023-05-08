package subscriber

import (
	"context"
	dbPackage "ethereum-explorer/db"
	"ethereum-explorer/ethClient"
	"ethereum-explorer/logger"
	"ethereum-explorer/models"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
)

func makeBlockModel(block *types.Block) *models.Block {
	return &models.Block{
		BlockHeight:block.Number().String(),
		Receipient:block.Coinbase().String(),
		Reward:block.BaseFee().String(),
		Size:strconv.FormatUint(block.Size(), 10),
		GasUsed:strconv.FormatUint(block.GasUsed(), 10),
		Hash:block.Hash().String(),}
}

func makeTransactionModel(transaction *types.Transaction, height *big.Int) (*models.Transaction, error) {
	msg, err := core.TransactionToMessage(transaction, types.LatestSignerForChainID(transaction.ChainId()), nil)
	if err != nil {
		return nil, err
	}
	return &models.Transaction{
		Hash:transaction.Hash().String(),
		BlockHeight:height.String(),
		From:msg.From.String(),
		To:transaction.To().String(),
		Value:transaction.Value().String(),
		TxFee:transaction.Cost().String(),
	}, nil
}

func insertBlockDocument(block *types.Block, db *dbPackage.DB, errorChan chan error) {
	err := db.InsertOneDocument("blocks", makeBlockModel(block))
	if err != nil {
		errorChan <- err
		return
	}
	
	transactions := block.Transactions()
	if transactions.Len() > 0 {
		var documents dbPackage.Documents

		for _, transaction := range transactions {
			document, err := makeTransactionModel(transaction, block.Number())
			if err != nil {
				errorChan <- err
				return
			}
			documents = append(documents, document)
		}
		
		err = db.InsertManyDocument("transactions", documents)
		if err != nil {
			errorChan <- err
			return
		}
	}
}

func insertPreviousDocuments(blocks []*types.Block, db *dbPackage.DB) error {
	var blockDocuments dbPackage.Documents
	var transactionDocuments dbPackage.Documents

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
		err := db.InsertManyDocument("blocks", blockDocuments)
		if err != nil {
			return err
		}
		if len(transactionDocuments) > 0 {
			err = db.InsertManyDocument("transactions", transactionDocuments)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type Subscriber struct {
	sub ethereum.Subscription
	header chan *types.Header
	errorChan chan error
}

func NewSubscriber(ethClient *ethClient.EthClient, db *dbPackage.DB, errorChan chan error) (*Subscriber, *big.Int, error) {
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

	go insertBlockDocument(block, db, errorChan)

	return &Subscriber{
		sub,
		headers,
		errorChan,
	}, block.Number(), nil
}

func (sub *Subscriber) ProcessSubscribe(ethClient *ethClient.EthClient, db *dbPackage.DB) {
	for {
		select {
		case header:= <-sub.header:
			block, err := ethClient.Http.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				sub.errorChan <- err
				return
			}
			go insertBlockDocument(block, db, sub.errorChan)
		case err := <-sub.sub.Err():
			sub.errorChan <- err
			return
		}
	}
}

func (sub *Subscriber) ProcessPrevious(ethClient *ethClient.EthClient, db *dbPackage.DB, start *big.Int, initBlock *big.Int) {
	one := big.NewInt(1)
	var blocks []*types.Block

	for i:=start ; i.Cmp(initBlock) < 0; i.Add(i, one) {
		block, err := ethClient.Http.BlockByNumber(context.Background(), i)
		if err != nil {
			sub.errorChan <- err
			return
		}
		blocks = append(blocks, block)
	}
	err := insertPreviousDocuments(blocks, db)
	if err != nil {
		sub.errorChan <- err
	}
}