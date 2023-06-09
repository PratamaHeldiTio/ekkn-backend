package repository

import (
	"backend-ekkn/modules/student_registration/domain"
)

type StudentRegistrationRepository interface {
	Create(registration domain.StudentRegistration) error
	FindByStudentId(id string) ([]domain.StudentRegistration, error)
	FindByNimPeriodId(nim, periodId string) (domain.StudentRegistration, error)
	FindByPeriod(periodID, query string) ([]domain.StudentRegistration, error)
	Update(registration domain.StudentRegistration) error
	FindByID(ID string) (domain.StudentRegistration, error)
	FindByGroup(ID string) ([]domain.StudentRegistration, error)
}
