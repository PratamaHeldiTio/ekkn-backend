package shareddomain

import (
	"backend-ekkn/modules/village/domain"
)

type RequestVillage struct {
	ID           string
	GroupID      string
	Name         string  `json:"name" binding:"required,max=50"`
	Kecamatan    string  `json:"kecamatan" binding:"required,max=50"`
	Kabupaten    string  `json:"kabupaten" binding:"required,max=50"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Strength     string  `json:"strength"`
	Weakness     string  `json:"weakness"`
	Oportunities string  `json:"oportunities"`
	Threats      string  `json:"threats"`
	Status       string  `json:"status"`
}

type UpdateVillageRequest struct {
	ID           string
	Strength     string `json:"strength"`
	Weakness     string `json:"weakness"`
	Oportunities string `json:"oportunities"`
	Threats      string `json:"threats"`
	Status       string `json:"status"`
}

type ResponseVillage struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Kecamatan string  `json:"kecamatan"`
	Kabupaten string  `json:"kabupaten"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Status    string  `json:"status"`
	CreatedAt int64   `json:"created_at"`
	UpdatedAt int64   `json:"updated_at"`
}

func ToVillageResponse(village domain.Village) ResponseVillage {
	return ResponseVillage{
		ID:        village.ID,
		Name:      village.Name,
		Kecamatan: village.Kecamatan,
		Kabupaten: village.Kabupaten,
		Latitude:  village.Latitude,
		Longitude: village.Longitude,
		Status:    village.Status,
		CreatedAt: village.CreatedAt,
		UpdatedAt: village.UpdatedAt,
	}
}
