package dto


type GetTransactionsResult struct {
	TransactionHash string
	Status          bool
	BlockHeight     string
	From            string
	To              string
	Value           string
	TransactionFee  string
	GasPrice        string
	GasLimit        uint64
	GasUsed         uint64
	Input           string
}

type GetTransactionsDTO struct {
	GetTransactionsResult []GetTransactionsResult
}

type GetTransactionsByHashDTO struct {
	GetTransactionsResult GetTransactionsResult
}

type GetTransactionsByBlockNumberResult struct {
	TxHash string
	Timestamp string
	From string
	To string
	Value string
	TxFee string
}

type GetTransactionsByBlockNumberDTO struct {
	GetTransactionsByBlockNumberResult []GetTransactionsByBlockNumberResult
}

func (gtDTO GetTransactionsDTO) GetDTO() interface{} {
	return gtDTO
}

func (gtbhDTO GetTransactionsByHashDTO) GetDTO() interface{} {
	return gtbhDTO
}

func (gtbnDTO GetTransactionsByBlockNumberDTO) GetDTO() interface{} {
	return gtbnDTO
}