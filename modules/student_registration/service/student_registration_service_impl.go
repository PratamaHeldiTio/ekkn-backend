package service

import (
	"backend-ekkn/modules/student_registration/domain"
	"backend-ekkn/modules/student_registration/repository"
	"backend-ekkn/pkg/shareddomain"
)

type StudentRegistrationServiceImpl struct {
	repo repository.StudentRegistrationRepository
}

// init repo
func NewStudentRegistrationService(repo repository.StudentRegistrationRepository) StudentRegistrationService {
	return &StudentRegistrationServiceImpl{repo}
}

func (service *StudentRegistrationServiceImpl) CreateStudentRegistration(request shareddomain.RequestStudentRegistration) error {
	studentRegistration := domain.StudentRegistration{
		PeriodID:  request.PeriodID,
		StudentID: request.Nim,
		Status:    "Belum validasi",
	}

	if err := service.repo.Create(studentRegistration); err != nil {
		return err
	}

	return nil
}
