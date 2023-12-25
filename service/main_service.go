package service

import (
	"ethereum-explorer/dto"
	"ethereum-explorer/repository"
	"net/http"
)

type MainService interface {
	GetInformationForMain(r *http.Request) (dto.GetInformationForMainDTO, error)
}

type mainService struct {
	mainRepository repository.MainRepository
}

func NewMainService(mainRepository repository.MainRepository) MainService {
	return &mainService{
		mainRepository,
	}
}

func(ms *mainService) GetInformationForMain(r *http.Request) (dto.GetInformationForMainDTO, error) {
	return ms.mainRepository.GetInformationForMain()
}