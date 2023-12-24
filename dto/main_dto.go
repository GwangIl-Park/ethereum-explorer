package dto

type BlockForMain struct {
	BlockNumber string
	Timestamp string
	FeeRecipient string
	TxCount uint64
	BlockReward string
}

type TransactionForMain struct {
	TxHash string
	From string
	To string
	Value string
}

type GetInformationForMainDTO struct {
	LatestBlockList []BlockForMain
	LatestTransactionList []TransactionForMain
}