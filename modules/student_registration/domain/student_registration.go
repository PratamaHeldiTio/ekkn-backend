package domain

import (
	periodDomain "backend-ekkn/modules/period/domain"
	studentDomain "backend-ekkn/modules/student/domain"
	"github.com/google/uuid"
)

type StudentRegistration struct {
	ID        uuid.UUID `gorm:"primary_key; unique;type:uuid; column:student_registration_id; default:uuid_generate_v4()"`
	PeriodID  uuid.UUID
	Period    periodDomain.Period
	StudentID string
	Student   studentDomain.Student
	Status    string `gorm:"type:varchar(5); default:false"`
	CreatedAt int64  `gorm:"autoCreateTime"`
}
