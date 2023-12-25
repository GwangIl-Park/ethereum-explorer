package service

import (
	"ethereum-explorer/dto"
	"ethereum-explorer/repository"
	"net/http"
)

type AccountService interface {
	GetAccountByAddress(r *http.Request) (dto.GetAccountByAddressDTO, error)
}

type accountService struct {
	accountRepository repository.AccountRepository
}

func NewAccountService(accountRepository repository.AccountRepository) AccountService {
	return &accountService{
		accountRepository,
	}
}

func(as *accountService) GetAccountByAddress(r *http.Request) (dto.GetAccountByAddressDTO, error) {
	address := r.RequestURI[len("/address/"):]
	return as.accountRepository.GetAccountByAddress(address)
}