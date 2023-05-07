package service

import (
	"backend-ekkn/modules/period/domain"
	"backend-ekkn/modules/period/repository"
	"backend-ekkn/pkg/helper"
	"backend-ekkn/pkg/shareddomain"
	"errors"
)

type PeriodServiceImpl struct {
	repo repository.PeriodRepository
}

func NewPeriodService(repo repository.PeriodRepository) PeriodService {
	return &PeriodServiceImpl{repo}
}

func (service *PeriodServiceImpl) CreatePeriod(request shareddomain.RequestPeriod) error {

	start, err := helper.StringDateToArray(request.Start, "start")
	if err != nil {
		return err
	}
	end, err := helper.StringDateToArray(request.End, "end")
	if err != nil {
		return err
	}

	period := domain.Period{
		Semester:    request.Semester,
		TahunAjaran: request.TahunAjaran,
		Start:       start,
		End:         end,
	}

	err = service.repo.Create(period)
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

func (service *PeriodServiceImpl) FindPeriodById(id string) (domain.Period, error) {
	// get data
	period, err := service.repo.FindById(id)
	if err != nil {
		return period, err
	}

	if period.ID == "" {
		return period, errors.New("Data tidak ditemukan")
	}

	return period, nil
}

// update period

func (service *PeriodServiceImpl) UpdatePeriod(request shareddomain.RequestPeriod) error {
	period, err := service.FindPeriodById(request.ID)
	if err != nil {
		return err
	}

	start, err := helper.StringDateToArray(request.Start, "start")
	if err != nil {
		return err
	}
	end, err := helper.StringDateToArray(request.End, "end")
	if err != nil {
		return err
	}

	period.Semester = request.Semester
	period.TahunAjaran = request.TahunAjaran
	period.Start = start
	period.End = end
	period.Status = request.Status
	period.StatusRegisterStudent = request.StudentRegistrationStatus
	period.StatusRegisterLecture = request.LectureRegistrationStatus
	period.StatusRegisterGroup = request.GroupRegistrationStatus

	if err := service.repo.Update(period); err != nil {
		return err
	}

	return nil
}

// service delete period

func (service *PeriodServiceImpl) DeletePeriodById(id string) error {
	// check data isExist
	period, err := service.repo.FindById(id)
	if err != nil {
		return err
	}

	if period.ID == "" {
		return errors.New("No student found on that nim")
	}

	if err := service.repo.Delete(period); err != nil {
		return err
	}

	return nil
}
