package service

import (
	"backend-ekkn/modules/period/domain"
	"backend-ekkn/modules/period/repository"
	"backend-ekkn/pkg/shareddomain"
	"errors"
	"github.com/google/uuid"
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

func (service *PeriodServiceImpl) FindPeriodById(id uuid.UUID) (domain.Period, error) {
	// get data
	period, err := service.repo.FindById(id)
	if err != nil {
		return period, err
	}

	if period.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return period, errors.New("No student found on that nim")
	}

	return period, nil
}

// update period

func (service *PeriodServiceImpl) UpdatePeriod(request shareddomain.RequestPeriod) error {
	period := domain.Period{
		ID:                    request.ID,
		Semester:              request.Semester,
		TahunAjaran:           request.TahunAjaran,
		StatusRegisterStudent: request.StatusRegisterStudent,
		StatusRegisterLecture: request.StatusRegisterLecture,
		StatusRegisterGroup:   request.StatusRegisterGroup,
		Status:                request.Status,
		Start:                 request.Start,
		End:                   request.End,
	}

	if err := service.repo.Update(period); err != nil {
		return err
	}

	return nil
}

// service delete period

func (service *PeriodServiceImpl) DeletePeriodById(id uuid.UUID) error {
	// check data isExist
	period, err := service.repo.FindById(id)
	if err != nil {
		return err
	}

	if period.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return errors.New("No student found on that nim")
	}

	if err := service.repo.Delete(period); err != nil {
		return err
	}

	return nil
}
