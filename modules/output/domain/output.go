package domain

import "backend-ekkn/modules/group/domain"

type Output struct {
	ID           string `gorm:"primary_key; type:uuid; default:uuid_generate_v4()"`
	Group        domain.Group
	GroupID      string
	Type         string `gorm:"type:varchar(100)"`
	File         string
	Contribution string
	Description  string
	CreatedAt    int64 `gorm:"autoCreateTime"`
	UpdatedAt    int64 `gorm:"autoUpdateTime"`
}
