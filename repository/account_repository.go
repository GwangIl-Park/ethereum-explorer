package repository

import (
	"context"
	"ethereum-explorer/db"
)

type AccountRepository interface {
	GetAccountByAddress(c context.Context, address string)
}

type accountRepository struct {
	db *db.DB
}

func NewAccountRepository(db *db.DB) AccountRepository {
	return &accountRepository{
		db,
	}
}

func (ar *accountRepository) GetAccountByAddress(c context.Context, address string) {

}
