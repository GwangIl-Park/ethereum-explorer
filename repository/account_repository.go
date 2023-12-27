package repository

import (
	"ethereum-explorer/db"
	"ethereum-explorer/dto"
	"strconv"
)

type AccountRepository interface {
	GetAccountByAddress(address string) (*dto.GetAccountByAddressDTO, error)
}

type accountRepository struct {
	db *db.DB
}

func NewAccountRepository(db *db.DB) AccountRepository {
	return &accountRepository{
		db,
	}
}

func (ar *accountRepository) GetAccountByAddress(address string) (*dto.GetAccountByAddressDTO, error) {
	var getAccountByAddressDTO *dto.GetAccountByAddressDTO

	var balance string
	err := ar.db.Client.QueryRow(`SELECT balance FROM account WHERE address = $1`, address).Scan(&balance)
	if err != nil {
		return nil, err
	}

	getAccountByAddressDTO.GetAccountByAddressResult.Balance, err = strconv.ParseUint(balance, 10, 64)
	if err != nil {
		return nil, err
	}

	txRows, err := ar.db.Client.Query(`SELECT txhash, blockheight, timestamp, fromaddress, toaddress, value, fee FROM transaction WHERE fromaddress = $1 OR toaddress = $2`, address, address)
	if err != nil {
		return nil, err
	}
	defer txRows.Close()

	for txRows.Next() {
		var transactionOfAccount dto.TransactionOfAccount
		err = txRows.Scan(&transactionOfAccount)
		if err != nil {
			return nil, err
		}
		getAccountByAddressDTO.GetAccountByAddressResult.TxList = append(getAccountByAddressDTO.GetAccountByAddressResult.TxList, transactionOfAccount)
	}
	return getAccountByAddressDTO, nil
}