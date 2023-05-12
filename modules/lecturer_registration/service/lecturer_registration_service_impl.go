package service

import (
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
