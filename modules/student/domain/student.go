package domain

type Student struct {
	Nim        string `gorm:"primaryKey"`
	Name       string
	Gender     string
	Password   string
	Position   string
	Prodi      string
	Fakultas   string
	MaduraLang bool
	GroupKkn   string
	Grade      string
	CreatedAt  int64 `gorm:"autoCreateTime"`
	UpdateAt   int64 `gorm:"autoUpdateTime"`
}
