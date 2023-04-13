package domain

type Student struct {
	Nim        string `gorm:"type:varchar(13);primaryKey"`
	Name       string `gorm:"type:varchar(100)"`
	Gender     string `gorm:"type:varchar(9)"`
	Password   string `gorm:"type:varchar(255)"`
	Position   string `gorm:"type:varchar(6)"`
	Prodi      string `gorm:"type:varchar(50)"`
	Fakultas   string `gorm:"type:varchar(50)"`
	MaduraLang string `gorm:"type:varchar(5)"`
	GroupKkn   string
	Grade      string `gorm:"type:varchar(2);"`
	CreatedAt  int64  `gorm:"autoCreateTime"`
	UpdateAt   int64  `gorm:"autoUpdateTime"`
}
