package repository

import (
	"backend-ekkn/modules/student_registration/domain"
)

type StudentRegistrationRepository interface {
	Create(registration domain.StudentRegistration) error
	FindByStudentId(id string) ([]domain.StudentRegistration, error)
	FindByNimPeriodId(nim, periodId string) (domain.StudentRegistration, error)
}
