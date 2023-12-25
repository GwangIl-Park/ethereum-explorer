package repository

import (
	"ethereum-explorer/db"
	"ethereum-explorer/dto"
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
	accountRow, err := ar.db.Client.Query(`SELECT address, balance FROM account WHERE address = $1`, address)
	if err != nil {
		return nil, err
	}
	defer accountRow.Close()
	
	txRows, err := ar.db.Client.Query(`SELECT txhash, blockheight, timestamp, fromaddress, toaddress, value, fee FROM transaction WHERE fromaddress = $1 OR toaddress = $2`, address, address)
	if err != nil {
		return nil, err
	}
	defer txRows.Close()

	var getAccountByAddressDTO *dto.GetAccountByAddressDTO
	for accountRow.Next() {
		err = accountRow.Scan(&getAccountByAddressDTO.GetAccountByAddressResult.Address, &getAccountByAddressDTO.GetAccountByAddressResult.Balance)
		if err != nil {
			return nil, err
		}
	}

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