package shareddomain

import (
	"backend-ekkn/modules/group/domain"
)

type RequestGroup struct {
	Name     string `json:"name" binding:"required,max=50"`
	PeriodID string `json:"period_id"`
	Leader   string `json:"leader"`
}

type GroupUpdateRequest struct {
	ID        string
	Village   string `json:"village"`
	Nim       string
	Proposal  string
	Report    string
	Potential string
}

type Student struct {
	Nim        string `json:"nim"`
	Name       string `json:"name"`
	Prodi      string `json:"prodi"`
	MaduraLang string `json:"madura_lang"`
}

type ResponseGroupByID struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Students  []Student `json:"students"`
	Leader    string    `json:"leader"`
	Referral  string    `json:"referral"`
	Village   Village   `json:"village"`
	Status    string    `json:"status"`
	CreatedAt int64     `json:"created_at"`
	UpdatedAt int64     `json:"updated_at"`
}

type RequestJoin struct {
	Referral string `json:"referral" binding:"required,max=6"`
}

type AddVillage struct {
	ID      string
	Nim     string
	Village string `json:"village"`
}

type Village struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Kecamatan string `json:"kecamatan"`
	Kabupaten string `json:"kabupaten"`
}

func ToResponseGroupByID(group domain.Group) ResponseGroupByID {
	// maping students
	var students []Student
	for _, student := range group.Students {
		students = append(students, Student{
			Nim:        student.Nim,
			Name:       student.Name,
			Prodi:      student.Prodi,
			MaduraLang: student.MaduraLang,
		})
	}

	// village
	village := Village{
		Name:      group.Village.Name,
		Kecamatan: group.Village.Kecamatan,
		Kabupaten: group.Village.Kabupaten,
	}

	// result response
	responseGroup := ResponseGroupByID{
		ID:        group.ID,
		Name:      group.Name,
		Students:  students,
		Leader:    group.Leader,
		Referral:  group.Referral,
		Status:    group.Status,
		Village:   village,
		CreatedAt: group.CreatedAt,
		UpdatedAt: group.UpdateAt,
	}

	return responseGroup
}

type GroupByPeriodLeaderResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Students  []Student `json:"students"`
	Leader    string    `json:"leader"`
	Referral  string    `json:"referral"`
	Status    string    `json:"status"`
	Proposal  string    `json:"proposal"`
	Report    string    `json:"report"`
	Potential string    `json:"potential"`
	Village   Village   `json:"village"`
	CreatedAt int64     `json:"created_at"`
	UpdatedAt int64     `json:"updated_at"`
}

func ToGroupByPeriodLeaderResponse(group domain.Group) GroupByPeriodLeaderResponse {
	// village
	village := Village{
		ID:        group.Village.ID,
		Name:      group.Village.Name,
		Kecamatan: group.Village.Kecamatan,
		Kabupaten: group.Village.Kabupaten,
	}

	// result response
	responseGroup := GroupByPeriodLeaderResponse{
		ID:        group.ID,
		Proposal:  group.Proposal,
		Report:    group.Report,
		Potential: group.Potential,
		Village:   village,
	}

	return responseGroup
}
