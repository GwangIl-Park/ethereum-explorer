package service

import "ethereum-explorer/repository"

type MainService interface {

}

type mainService struct {
	mainRepository repository.MainRepository
}

func NewMainService(mainRepository repository.MainRepository) MainService {
	return &mainService{
		mainRepository,
	}
}