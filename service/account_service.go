package service

import "ethereum-explorer/repository"

type AccountService interface {

}

type accountService struct {
	accountRepository repository.AccountRepository
}

func NewAccountService(accountRepository repository.AccountRepository) AccountService {
	return &accountService{
		accountRepository,
	}
}