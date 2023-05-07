package subscriber

import (
	"context"
	"errors"
	dbPackage "ethereum-explorer/db"
	"ethereum-explorer/ethClient"
	"ethereum-explorer/logger"
	"ethereum-explorer/models"
	"fmt"
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

func insertDocument(block *types.Block, db *dbPackage.DB) error {
	document, err := db.ReadDocument("blocks", "height", block.Number().String())
	if err != nil {
		return err
	}

	if document == nil {
		return errors.New(fmt.Errorf("block %s already exists", block.Number().String()).Error())
	}
	
	err = db.InsertOneDocument("blocks", makeBlockModel(block))
	if err != nil{
		return err
	}
	
	transactions := block.Transactions()

	var documents dbPackage.Documents

	for _, transaction := range transactions {
		document, err := makeTransactionModel(transaction, block.Number())
		if err != nil {
			return err
		}
		documents = append(documents, document)
	}
	
	err = db.InsertManyDocument("transactions", documents)
	if err != nil {
		return err
	}

	return nil
}

func insertPreviousDocuments(blocks []*types.Block, db *dbPackage.DB) error {
	var blockDocuments dbPackage.Documents
	var transactionDocuments dbPackage.Documents

	for _, block := range blocks {
		blockDocuments = append(blockDocuments, makeBlockModel(block))

		transactions := block.Transactions()

		for _, transaction := range transactions {
			document, err := makeTransactionModel(transaction, block.Number())
			if err != nil {
				return err
			}
			transactionDocuments = append(transactionDocuments, document)
		}
	}

	err := db.InsertManyDocument("blocks", blockDocuments)
	if err != nil {
		return err
	}
	
	err = db.InsertManyDocument("transactions", transactionDocuments)
	if err != nil {
		return err
	}

	return nil
}

type Subscriber struct {
	sub ethereum.Subscription
	header chan *types.Header
}

func NewSubscriber(ethClient *ethClient.EthClient, db *dbPackage.DB) (*Subscriber, *big.Int, error) {
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
		fmt.Println("bb")
		return nil, nil, err
	}

	go insertDocument(block, db)

	return &Subscriber{
		sub,
		headers,
	}, block.Number(), nil
}

func (sub *Subscriber) ProcessSubscribe(ethClient *ethClient.EthClient, db *dbPackage.DB, errorChan chan error) {
	for {
		select {
		case header:= <-sub.header:
			block, err := ethClient.Http.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				errorChan <- err
				return
			}
			go insertDocument(block, db)
		case err := <-sub.sub.Err():
			errorChan <- err
			return
		}
	}
}

func (sub *Subscriber) ProcessPrevious(ethClient *ethClient.EthClient, db *dbPackage.DB, start *big.Int, initBlock *big.Int, errorChan chan error) {
	one := big.NewInt(1)
	var blocks []*types.Block
	for i := start; i.Cmp(initBlock) > 0; i.Add(i, one) {
		block, err := ethClient.Http.BlockByNumber(context.Background(), i)
		if err != nil {
			errorChan <- err
			return
		}
		blocks = append(blocks, block)
	}
	insertPreviousDocuments(blocks, db)
}