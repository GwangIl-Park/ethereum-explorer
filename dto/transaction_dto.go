package dto

type GetTransactionByBlockNumberDTO struct {
	TxHash string
	Timestamp string
	From string
	To string
	Value string
	TxFee string
}