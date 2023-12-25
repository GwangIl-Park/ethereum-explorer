package service

import (
	"ethereum-explorer/dto"
	"ethereum-explorer/repository"
)

type MainService interface {
	GetInformationForMain() (*dto.GetInformationForMainDTO, error)
}

type mainService struct {
	mainRepository repository.MainRepository
}

func NewMainService(mainRepository repository.MainRepository) MainService {
	return &mainService{
		mainRepository,
	}
}

func(ms *mainService) GetInformationForMain() (*dto.GetInformationForMainDTO, error) {
	return ms.mainRepository.GetInformationForMain()
}