package repository

import (
	"ethereum-explorer/db"
	"ethereum-explorer/dto"
)

type MainRepository interface {
	GetInformationForMain() (dto.GetInformationForMainDTO)
}

type mainRepository struct {
	db *db.DB
}

func NewMainRepository(db *db.DB) MainRepository {
	return &mainRepository{
		db,
	}
}

func (mr *mainRepository) GetInformationForMain() (dto.GetInformationForMainDTO) {
	blockRows, err := mr.db.Client.Query(`SELECT blocknumber, timestamp, feeRecipient, blockReward FROM block LIMIT 6`)
	if err != nil {
		panic(err)
	}
	defer blockRows.Close()

	transactionRows, err := mr.db.Client.Query(`SELECT txHash, from, to, value FROM transaction LIMIT 6`)
	if err != nil {
		panic(err)
	}
	defer transactionRows.Close()

	var getInformationForMain dto.GetInformationForMainDTO


	for blockRows.Next() {
		var blockForMain dto.BlockForMain
		err = blockRows.Scan(&blockForMain)
		if err != nil {
			panic(err)
		}
		getInformationForMain.LatestBlockList = append(getInformationForMain.LatestBlockList, blockForMain)
	}

	for transactionRows.Next() {
		var transactionForMain dto.TransactionForMain
		err = transactionRows.Scan(&transactionForMain)
		if err != nil {
			panic(err)
		}
		getInformationForMain.LatestTransactionList = append(getInformationForMain.LatestTransactionList, transactionForMain)
	}

	return getInformationForMain
}