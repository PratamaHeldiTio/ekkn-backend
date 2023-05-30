package domain

import (
	periodDomain "backend-ekkn/modules/period/domain"
	studentDomain "backend-ekkn/modules/student/domain"
)

type StudentRegistration struct {
	ID        string `gorm:"primary_key; unique;type:uuid; column:student_registration_id; default:uuid_generate_v4()"`
	PeriodID  string
	Period    periodDomain.Period
	StudentID string
	Student   studentDomain.Student
	GroupID   string
	Status    string `gorm:"type:varchar(5)"`
	Proker    string
	Grade     uint8
	CreatedAt int64 `gorm:"autoCreateTime"`
}
