package repository

import (
	"backend-ekkn/modules/student_registration/domain"
	"github.com/google/uuid"
)

type StudentRegistrationRepository interface {
	Create(registration domain.StudentRegistration) error
	FindByStudentId(id string) ([]domain.StudentRegistration, error)
	FindByStudentIdPeriodId(studentId string, periodId uuid.UUID) (domain.StudentRegistration, error)
}
