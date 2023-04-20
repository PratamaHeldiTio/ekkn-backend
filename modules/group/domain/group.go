package domain

import (
	"backend-ekkn/modules/period/domain"
	student "backend-ekkn/modules/student/domain"
)

type Group struct {
	ID        string            `gorm:"primary_key; type:uuid; column:group_id; default:uuid_generate_v4()"`
	Name      string            `gorm:"type:varchar(100)"`
	Students  []student.Student `gorm:"many2many:student_groups;"`
	PeriodID  string
	Period    domain.Period
	Leader    string `gorm:"type:varchar(12)"`
	Referral  string `gorm:"type:varchar(6); unique; not null"`
	Status    string `gorm:"type:varchar(5); default:false"`
	CreatedAt int64  `gorm:"autoCreateTime"`
	UpdateAt  int64  `gorm:"autoUpdateTime"`
}

type StudentGroup struct {
	GroupID    string
	StudentNim string
}