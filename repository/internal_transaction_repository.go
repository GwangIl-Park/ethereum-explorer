package repository

import (
	"context"
	"ethereum-explorer/db"
)

type InternalTransactionRepository interface {
	GetInternalTransactionsByBlockHeight(c context.Context, blockHeight string)
}

type internalTransactionRepository struct {
	db *db.DB
}

func NewInternalTransactionRepository(db *db.DB) InternalTransactionRepository {
	return &internalTransactionRepository{
		db,
	}
}

func (itr *internalTransactionRepository) GetInternalTransactionsByBlockHeight(c context.Context, blockHeight string) {

}
