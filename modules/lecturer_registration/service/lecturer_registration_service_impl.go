package service

import (
	"backend-ekkn/modules/lecturer_registration/domain"
	"backend-ekkn/modules/lecturer_registration/repository"
	"backend-ekkn/pkg/shareddomain"
	"errors"
)

type lecturerRegistrationServiceImpl struct {
	repo repository.LecturerRegistrationRepository
}

func NewLecturerRegistrationService(repo repository.LecturerRegistrationRepository) LecturerRegistrationService {
	return &lecturerRegistrationServiceImpl{repo}
}

func (service *lecturerRegistrationServiceImpl) LecturerRegistration(request shareddomain.LecturerRegistrationRequest) error {
	// get lecturer registration by lecturer id and period id
	lecturerRegistration, err := service.repo.FindByPeriodLectureID(request.PeriodID, request.LecturerID)
	if err != nil {
		return err
	}
	// cek isExist
	if lecturerRegistration.ID != "" {
		return errors.New("Pendaftaran gagal anda telah terdaftar")
	}

	lecturerRegistration.LecturerID = request.LecturerID
	lecturerRegistration.PeriodID = request.PeriodID

	if err := service.repo.Create(lecturerRegistration); err != nil {
		return err
	}

	return nil
}

func (service *lecturerRegistrationServiceImpl) FindLecturerRegistrationByLectureID(lectureID string) ([]domain.LecturerRegistration, error) {
	lecturerRegistrations, err := service.repo.FindByLectureID(lectureID)
	if err != nil {
		return lecturerRegistrations, err
	}

	return lecturerRegistrations, nil
}

func (service *lecturerRegistrationServiceImpl) FindLecturerRegistrationByID(ID string) (domain.LecturerRegistration, error) {
	lecturerRegistration, err := service.repo.FindByID(ID)
	if err != nil {
		return lecturerRegistration, err
	}

	if lecturerRegistration.ID == "" {
		return lecturerRegistration, errors.New("data tidak ditemukan")
	}

	return lecturerRegistration, nil
}

func (service *lecturerRegistrationServiceImpl) ValidationLecturerRegistration(request shareddomain.ValidationLectureRegistrationRequest) error {
	lecturerRegistration, err := service.FindLecturerRegistrationByID(request.ID)
	if err != nil {
		return err
	}

	// update
	lecturerRegistration.Status = request.Status
	if err := service.repo.Update(lecturerRegistration); err != nil {
		return err
	}

	return nil
}

func (service *lecturerRegistrationServiceImpl) FindLecturerRegistrationByPeriod(ID, query string) ([]domain.LecturerRegistration, error) {
	lecturerRegistrations, err := service.repo.FindByPeriod(ID, query)
	if err != nil {
		return lecturerRegistrations, err
	}

	return lecturerRegistrations, nil
}
