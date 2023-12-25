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

type GetAccountByAddressResult struct {
	Address string
	Balance uint64
	TxList []TransactionOfAccount
}

type GetAccountByAddressDTO struct {
	GetAccountByAddressResult GetAccountByAddressResult
}

func (gabaDTO GetAccountByAddressDTO) GetDTO() interface{} {
	return gabaDTO
}