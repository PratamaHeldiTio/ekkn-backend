package domain

type Admin struct {
	Username string `gorm:"type:varchar(50);primaryKey"`
	Password string `gorm:"type:varchar(255)"`
}
