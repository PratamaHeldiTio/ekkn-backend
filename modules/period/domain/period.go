package domain

import (
	"github.com/google/uuid"
)

type Period struct {
	ID                    uuid.UUID `gorm:"primary_key; unique;type:uuid; column:id_period; default:uuid_generate_v4()" json:"id" json:"semester"`
	Semester              string    `gorm:"type:varchar(6)" json:"semester"`
	TahunAjaran           string    `gorm:"type:varchar(10)" json:"tahun_ajaran"`
	StatusRegisterStudent bool      `json:"status_register_student"`
	StatusRegisterLecture bool      `json:"status_register_lecture"`
	StatusRegisterGroup   bool      `json:"status_register_group"`
	Status                bool      `json:"status"`
	Start                 int64     `json:"start"`
	End                   int64     `json:"end"`
	CreatedAt             int64     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt             int64     `gorm:"autoUpdateTime" json:"updated_at"`
}
