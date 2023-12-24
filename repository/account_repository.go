package repository

import (
	"context"
	"ethereum-explorer/db"
	"ethereum-explorer/dto"
)

type AccountRepository interface {
	GetAccountByAddress(c context.Context, address string) dto.GetAccountByAddressDTO
}

type accountRepository struct {
	db *db.DB
}

func NewAccountRepository(db *db.DB) AccountRepository {
	return &accountRepository{
		db,
	}
}

func (ar *accountRepository) GetAccountByAddress(c context.Context, address string) dto.GetAccountByAddressDTO {
	accountRow, err := ar.db.Client.Query(`SELECT address, balance FROM account WHERE address = %s`, address)
	if err != nil {
		panic(err)
	}

	defer accountRow.Close()

	txRows, err := ar.db.Client.Query(`SELECT txHash, blockNumber, timestamp, from, to, value, fee FROM transaction WHERE from = %s OR to = %s`, address, address)
	if err != nil {
		panic(err)
	}

	defer txRows.Close()

	var getAccountByAddressDTO dto.GetAccountByAddressDTO
	accountRow.Scan(&getAccountByAddressDTO.Address, &getAccountByAddressDTO.Balance)

	for txRows.Next() {
		var transactionOfAccount dto.TransactionOfAccount
		err = txRows.Scan(&transactionOfAccount)
		if err != nil {
			panic(err)
		}
		getAccountByAddressDTO.TxList = append(getAccountByAddressDTO.TxList, transactionOfAccount)
	}

	return getAccountByAddressDTO
}