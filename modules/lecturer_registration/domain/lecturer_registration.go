package domain

import (
	lecture "backend-ekkn/modules/lecturer/domain"
	periodDomain "backend-ekkn/modules/period/domain"
)

type LecturerRegistration struct {
	ID         string `gorm:"primary_key;type:uuid; column:id; default:uuid_generate_v4()"`
	PeriodID   string
	Period     periodDomain.Period
	LecturerID string
	Lecturer   lecture.Lecturer
	Status     string `gorm:"type:varchar(5)"`
	CreatedAt  int64  `gorm:"autoCreateTime"`
}
