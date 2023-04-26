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
	ID       string
	Village  string `json:"village"`
	Nim      string
	Proposal string
	Report   string
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

	// result response
	responseGroup := ResponseGroupByID{
		ID:        group.ID,
		Name:      group.Name,
		Students:  students,
		Leader:    group.Leader,
		Referral:  group.Referral,
		Status:    group.Status,
		CreatedAt: group.CreatedAt,
		UpdatedAt: group.UpdateAt,
	}

	return responseGroup
}
