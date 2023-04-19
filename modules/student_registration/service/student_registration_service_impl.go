package service

import (
	"backend-ekkn/modules/student_registration/domain"
	"backend-ekkn/modules/student_registration/repository"
	"backend-ekkn/pkg/shareddomain"
	"errors"
	"fmt"
)

type StudentRegistrationServiceImpl struct {
	repo repository.StudentRegistrationRepository
}

// init repo
func NewStudentRegistrationService(repo repository.StudentRegistrationRepository) StudentRegistrationService {
	return &StudentRegistrationServiceImpl{repo}
}

func (service *StudentRegistrationServiceImpl) CreateStudentRegistration(request shareddomain.RequestStudentRegistration) error {
	// get register by student id and period id
	registeredStudent, err := service.FindStudentRegistrationByNimPeriodID(request.Nim, request.PeriodID)
	if err != nil {
		return err
	}
	// cek isExist
	if registeredStudent.ID != "" {
		return errors.New("Pendaftaran gagal anda telah terdaftar")
	}

	fmt.Println(registeredStudent)
	studentRegistration := domain.StudentRegistration{
		PeriodID:  request.PeriodID,
		StudentID: request.Nim,
	}

	if err := service.repo.Create(studentRegistration); err != nil {
		return err
	}

	return nil
}

func (service *StudentRegistrationServiceImpl) FindStudentRegistrationByStudentID(id string) ([]domain.StudentRegistration, error) {
	registeredUser, err := service.repo.FindByStudentId(id)
	if err != nil {
		return registeredUser, err
	}
	return registeredUser, err
}

func (service *StudentRegistrationServiceImpl) FindStudentRegistrationByNimPeriodID(nim, periodID string) (domain.StudentRegistration, error) {
	// get register by student id and period id
	registeredStudent, err := service.repo.FindByNimPeriodId(nim, periodID)
	if err != nil {
		return registeredStudent, err
	}

	return registeredStudent, nil
}
