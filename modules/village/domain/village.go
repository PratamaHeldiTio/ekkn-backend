package domain

import period "backend-ekkn/modules/period/domain"

type Village struct {
	ID           string `gorm:"primary_key; type:uuid; default:uuid_generate_v4()"`
	Name         string `gorm:"type:varchar(100)"`
	Kecamatan    string `gorm:"type:varchar(50)"`
	Kabupaten    string `gorm:"type:varchar(50)"`
	Period       period.Period
	PeriodID     string
	Latitude     float64
	Longitude    float64
	Strength     string
	Weakness     string
	Oportunities string
	Threats      string
	Status       string `gorm:"type:varchar(6); default:false"`
	CreatedAt    int64  `gorm:"autoCreateTime"`
	UpdatedAt    int64  `gorm:"autoUpdateTime"`
}
