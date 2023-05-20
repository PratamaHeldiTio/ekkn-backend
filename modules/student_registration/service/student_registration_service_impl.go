package service

import (
	"backend-ekkn/modules/student_registration/domain"
	"backend-ekkn/modules/student_registration/repository"
	"backend-ekkn/pkg/shareddomain"
	"errors"
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
	return registeredUser, nil
}

func (service *StudentRegistrationServiceImpl) FindStudentRegistrationByNimPeriodID(nim, periodID string) (domain.StudentRegistration, error) {
	// get register by student id and period id
	registeredStudent, err := service.repo.FindByNimPeriodId(nim, periodID)
	if err != nil {
		return registeredStudent, err
	}

	return registeredStudent, nil
}

func (service *StudentRegistrationServiceImpl) FindStudentRegistrationByPeriod(periodID string) ([]domain.StudentRegistration, error) {
	studentRegistation, err := service.repo.FindByPeriod(periodID)
	if err != nil {
		return studentRegistation, err
	}
	return studentRegistation, nil
}

func (service *StudentRegistrationServiceImpl) UpdateStudentRegistration(request shareddomain.UpdateStudentRegistrationRequest) error {
	// cek registration isExist
	registration, err := service.repo.FindByID(request.ID)
	if err != nil {
		return err
	}

	if registration.ID == "" {
		return errors.New("data tidak ditemukan")
	}

	registration.Status = request.Status
	if err := service.repo.Update(registration); err != nil {
		return err
	}
	return nil
}

func (service *StudentRegistrationServiceImpl) AddProkerStudent(request shareddomain.AddProkerStudent) error {
	// cek registration isExist
	registration, err := service.FindStudentRegistrationByID(request.ID)
	if err != nil {
		return err
	}

	registration.Proker = request.Proker
	if err := service.repo.Update(registration); err != nil {
		return err
	}
	return nil
}

func (service *StudentRegistrationServiceImpl) FindStudentRegistrationByID(ID string) (domain.StudentRegistration, error) {
	// cek registration isExist
	registration, err := service.repo.FindByID(ID)
	if err != nil {
		return registration, err
	}

	if registration.ID == "" {
		return registration, errors.New("data tidak ditemukan")
	}

	return registration, nil
}

func (service *StudentRegistrationServiceImpl) FindStudentRegistrationByGroup(ID string) ([]domain.StudentRegistration, error) {
	studentRegistation, err := service.repo.FindByGroup(ID)
	if err != nil {
		return studentRegistation, err
	}
	return studentRegistation, nil
}
