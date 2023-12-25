package service

import (
	"ethereum-explorer/dto"
	"ethereum-explorer/repository"
)

type AccountService interface {
	GetAccountByAddress(address string) (*dto.GetAccountByAddressDTO, error)
}

type accountService struct {
	accountRepository repository.AccountRepository
}

func NewAccountService(accountRepository repository.AccountRepository) AccountService {
	return &accountService{
		accountRepository,
	}
}

func(as *accountService) GetAccountByAddress(address string) (*dto.GetAccountByAddressDTO, error) {
	return as.accountRepository.GetAccountByAddress(address)
}