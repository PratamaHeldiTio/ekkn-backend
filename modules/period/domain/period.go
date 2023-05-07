package domain

type Period struct {
	ID                    string `gorm:"primary_key; type:uuid; column:id; default:uuid_generate_v4()"`
	Semester              string `gorm:"type:varchar(6)"`
	TahunAjaran           string `gorm:"type:varchar(10)"`
	StatusRegisterStudent string `gorm:"type:varchar(5); default:false"`
	StatusRegisterLecture string `gorm:"type:varchar(5); default:false"`
	StatusRegisterGroup   string `gorm:"type:varchar(5); default:false"`
	Status                string `gorm:"type:varchar(5); default:false"`
	Start                 int64
	End                   int64
	CreatedAt             int64 `gorm:"autoCreateTime"`
	UpdatedAt             int64 `gorm:"autoUpdateTime"`
}
