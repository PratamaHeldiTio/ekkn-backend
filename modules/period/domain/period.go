package domain

import (
	"github.com/google/uuid"
)

type Period struct {
	ID                    uuid.UUID `gorm:"primary_key; unique;type:uuid; column:period_id; default:uuid_generate_v4()"`
	Semester              string    `gorm:"type:varchar(6)"`
	TahunAjaran           string    `gorm:"type:varchar(10)"`
	StatusRegisterStudent string    `gorm:"type:varchar(5)"`
	StatusRegisterLecture string    `gorm:"type:varchar(5)"`
	StatusRegisterGroup   string    `gorm:"type:varchar(5)"`
	Status                string    `gorm:"type:varchar(5)"`
	Start                 int64
	End                   int64
	CreatedAt             int64 `gorm:"autoCreateTime"`
	UpdatedAt             int64 `gorm:"autoUpdateTime"`
}
