package domain

import (
	"github.com/google/uuid"
)

type Period struct {
	ID                    uuid.UUID `gorm:"primary_key; unique;type:uuid; column:id_period; default:uuid_generate_v4()"`
	Semester              string    `gorm:"type:varchar(6)"`
	TahunAjaran           string    `gorm:"type:varchar(10)"`
	StatusRegisterStudent bool
	StatusRegisterLecture bool
	StatusRegisterGroup   bool
	Status                bool
	Start                 int64
	End                   int64
	CreatedAt             int64 `gorm:"autoCreateTime"`
	UpdatedAt             int64 `gorm:"autoUpdateTime"`
}
