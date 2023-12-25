package repository

import (
	"ethereum-explorer/db"
	"ethereum-explorer/dto"
)

type AccountRepository interface {
	GetAccountByAddress(address string) (dto.GetAccountByAddressDTO, error)
}

type accountRepository struct {
	db *db.DB
}

func NewAccountRepository(db *db.DB) AccountRepository {
	return &accountRepository{
		db,
	}
}

func (ar *accountRepository) GetAccountByAddress(address string) (dto.GetAccountByAddressDTO, error) {
	accountRow, err := ar.db.Client.Query(`SELECT address, balance FROM account WHERE address = %s`, address)
	if err != nil {
		return dto.GetAccountByAddressDTO{}, err
	}

	defer accountRow.Close()

	txRows, err := ar.db.Client.Query(`SELECT txHash, blockNumber, timestamp, from, to, value, fee FROM transaction WHERE from = %s OR to = %s`, address, address)
	if err != nil {
		return dto.GetAccountByAddressDTO{}, err
	}

	defer txRows.Close()

	var getAccountByAddressDTO dto.GetAccountByAddressDTO
	err = accountRow.Scan(getAccountByAddressDTO.Address, getAccountByAddressDTO.Balance)
	if err != nil {
		return dto.GetAccountByAddressDTO{}, err
	}

	for txRows.Next() {
		var transactionOfAccount dto.TransactionOfAccount
		err = txRows.Scan(&transactionOfAccount)
		if err != nil {
			return dto.GetAccountByAddressDTO{}, err
		}
		getAccountByAddressDTO.TxList = append(getAccountByAddressDTO.TxList, transactionOfAccount)
	}

	return getAccountByAddressDTO, nil
}