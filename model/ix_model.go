package model

type InternalTransaction struct {
	Id uint64 `json:"timestamp"`
	BlockHeight string `json:"blockHeight"`
	TransactionHash string `json:"transactionHash"`
	From string `json:"from"`
	To string `json:"to"`
	Data string `json:"balance"`
}