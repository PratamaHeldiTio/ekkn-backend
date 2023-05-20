package service

import (
	"backend-ekkn/modules/student_registration/domain"
	"backend-ekkn/pkg/shareddomain"
)

type StudentRegistrationService interface {
	CreateStudentRegistration(request shareddomain.RequestStudentRegistration) error
	FindStudentRegistrationByStudentID(id string) ([]domain.StudentRegistration, error)
	FindStudentRegistrationByNimPeriodID(nim, periodID string) (domain.StudentRegistration, error)
	FindStudentRegistrationByPeriod(periodID string) ([]domain.StudentRegistration, error)
	UpdateStudentRegistration(shareddomain.UpdateStudentRegistrationRequest) error
	AddProkerStudent(request shareddomain.AddProkerStudent) error
	FindStudentRegistrationByID(ID string) (domain.StudentRegistration, error)
	FindStudentRegistrationByGroup(ID string) ([]domain.StudentRegistration, error)
	SaveGradeStudent(request shareddomain.SaveGradeRequest) error
}
