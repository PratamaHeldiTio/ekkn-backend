package domain

import (
	periodDomain "backend-ekkn/modules/period/domain"
	studentDomain "backend-ekkn/modules/student/domain"
	"github.com/google/uuid"
)

type StudentRegistration struct {
	ID        uuid.UUID `gorm:"primary_key; unique;type:uuid; column:id_student_registration; default:uuid_generate_v4()"`
	PeriodID  uuid.UUID `gorm:"unique"`
	Period    periodDomain.Period
	StudentID string
	Student   studentDomain.Student
	Status    string `gorm:"type:varchar(14)"`
	CreatedAt int64  `gorm:"autoCreateTime"`
}
