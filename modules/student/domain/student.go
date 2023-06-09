package domain

type Student struct {
	Nim        string `gorm:"type:varchar(12);primaryKey"`
	Name       string `gorm:"type:varchar(100)"`
	Gender     string `gorm:"type:varchar(9)"`
	Password   string `gorm:"type:varchar(255)"`
	Prodi      string `gorm:"type:varchar(50)"`
	Fakultas   string `gorm:"type:varchar(50)"`
	MaduraLang string `gorm:"type:varchar(5)"`
	Profile    string `gorm:"type:varchar(255)"`
	CreatedAt  int64  `gorm:"autoCreateTime"`
	UpdateAt   int64  `gorm:"autoUpdateTime"`
}
