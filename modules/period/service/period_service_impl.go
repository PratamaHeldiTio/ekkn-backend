package service

import (
	"backend-ekkn/modules/period/domain"
	"backend-ekkn/modules/period/repository"
	"backend-ekkn/pkg/shareddomain"
)

type PeriodServiceImpl struct {
	repo repository.PeriodRepository
}

func NewPeriodService(repo repository.PeriodRepository) PeriodService {
	return &PeriodServiceImpl{repo}
}

func (service *PeriodServiceImpl) CreatePeriod(request shareddomain.RequestPeriod) error {
	period := domain.Period{
		Semester:              request.Semester,
		TahunAjaran:           request.TahunAjaran,
		StatusRegisterStudent: request.StatusRegisterStudent,
		StatusRegisterLecture: request.StatusRegisterLecture,
		StatusRegisterGroup:   request.StatusRegisterGroup,
		Status:                request.Status,
		Start:                 request.Start,
		End:                   request.End,
	}

	err := service.repo.Create(period)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (service *PeriodServiceImpl) FindAllPeriod() ([]domain.Period, error) {
	periods, err := service.repo.FindAll()
	if err != nil {
		return periods, err
	}

	return periods, nil
}
