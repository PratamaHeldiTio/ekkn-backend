package domain

import (
	"backend-ekkn/modules/period/domain"
	student "backend-ekkn/modules/student/domain"
)

type Logbook struct {
	ID        string `gorm:"primary_key; type:uuid; default:uuid_generate_v4()"`
	PeriodID  string
	Period    domain.Period
	StudentID string
	Student   student.Student
	Activity  string
	Image     string `gorm:"type:varchar(255)"`
	Radius    int    // in database save with scale meter from save to center village
	Date      int64
	Submitted int64
}
