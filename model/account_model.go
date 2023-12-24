package model

type Account struct {
	Id uint64 `json:"timestamp"`
	Address string `json:"address"`
	Balance uint64 `json:"balance"`
}