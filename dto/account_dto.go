package dto

type TransactionOfAccount struct {
	TxHash string
	BlockNumber string
	Timestamp string
	From string
	To string
	Value string
	Fee string
}

type GetAccountByAddressDTO struct {
	Address string
	Balance uint64
	TxList []TransactionOfAccount
}