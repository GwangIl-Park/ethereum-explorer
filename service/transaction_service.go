package service

import (
	"context"
	"ethereum-explorer/dto"
	"ethereum-explorer/ethClient"
	"ethereum-explorer/model"
	"ethereum-explorer/repository"

	"github.com/ethereum/go-ethereum/core/types"
)

type TransactionService interface {
	GetTransactions() (*dto.GetTransactionsDTO, error)
	GetTransactionByHash(txHash string) (*dto.GetTransactionsByHashDTO, error)
	GetTransactionsByBlockNumber(blockNumber string) (*dto.GetTransactionsByBlockNumberDTO, error)
	CreateTransaction(transaction *model.Transaction) error
	CreateTransactions(transactions []*model.Transaction) error
	CreateTransactionsFromCoreBlocks(ethClient *ethClient.EthClient, blocks []*types.Block) error
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionService(transactionRepository repository.TransactionRepository) TransactionService {
	return &transactionService{
		transactionRepository,
	}
}

func (ts *transactionService) GetTransactions() (*dto.GetTransactionsDTO, error) {
	return ts.transactionRepository.GetTransactions()
}

func (ts *transactionService) GetTransactionByHash(txHash string) (*dto.GetTransactionsByHashDTO, error) {
	return ts.transactionRepository.GetTransactionByHash(txHash)
}

func (ts *transactionService) GetTransactionsByBlockNumber(blockNumber string) (*dto.GetTransactionsByBlockNumberDTO, error) {
	return ts.transactionRepository.GetTransactionsByBlockNumber(blockNumber)
}

func (ts *transactionService) CreateTransaction(transaction *model.Transaction) error {
	return ts.transactionRepository.CreateTransaction(transaction)
}

func (ts *transactionService) CreateTransactions(transactions []*model.Transaction) error {
	return ts.transactionRepository.CreateTransactions(transactions)
}

func (ts *transactionService) CreateTransactionsFromCoreBlocks(ethClient *ethClient.EthClient, blocks []*types.Block) error {
	ctx := context.Background()
	var transactionmodel []*model.Transaction
	for _, block := range blocks {
		transactions := block.Transactions()
		if transactions.Len() > 0 {
			for _, transaction := range transactions {
				receipt, err := ethClient.Http.TransactionReceipt(ctx, transaction.Hash())
				if err != nil {
					return err
				}
				transactionModel, err := model.MakeTransactionModelFromTypes(receipt, transaction, *block)
				if err != nil {
					return err
				}
				transactionmodel = append(transactionmodel, transactionModel)
			}
		}
	}
	return ts.transactionRepository.CreateTransactions(transactionmodel)
}
