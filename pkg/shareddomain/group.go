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
	Report    string
	Potential string
}

type Student struct {
	Nim        string `json:"nim"`
	Name       string `json:"name"`
	Prodi      string `json:"prodi"`
	Fakultas   string `json:"fakultas"`
	MaduraLang string `json:"madura_lang"`
}

type Village struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Kecamatan    string `json:"kecamatan"`
	Kabupaten    string `json:"kabupaten"`
	Strength     string `json:"strength"`
	Weakness     string `json:"weakness"`
	Oportunities string `json:"oportunities"`
	Threats      string `json:"threats"`
}

type Lecturer struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ResponseGroupByID struct {
	ID        string    `json:"id"`
	PeriodID  string    `json:"period_id"`
	Lecturer  Lecturer  `json:"lecturer"`
	Name      string    `json:"name"`
	Students  []Student `json:"students"`
	Leader    string    `json:"leader"`
	Referral  string    `json:"referral"`
	Village   Village   `json:"village"`
	Status    string    `json:"status"`
	Potential string    `json:"potential"`
	Report    string    `json:"report"`
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
			Fakultas:   student.Fakultas,
			MaduraLang: student.MaduraLang,
		})
	}

	// village
	village := Village{
		ID:           group.Village.ID,
		Name:         group.Village.Name,
		Kecamatan:    group.Village.Kecamatan,
		Kabupaten:    group.Village.Kabupaten,
		Strength:     group.Village.Strength,
		Weakness:     group.Village.Weakness,
		Oportunities: group.Village.Oportunities,
		Threats:      group.Village.Threats,
	}

	// lecturer
	lecturer := Lecturer{
		ID:   group.Lecturer.ID,
		Name: group.Lecturer.Name,
	}

	// result response
	responseGroup := ResponseGroupByID{
		ID:        group.ID,
		PeriodID:  group.PeriodID,
		Name:      group.Name,
		Students:  students,
		Leader:    group.Leader,
		Referral:  group.Referral,
		Status:    group.Status,
		Village:   village,
		Potential: group.Potential,
		Report:    group.Report,
		Lecturer:  lecturer,
		CreatedAt: group.CreatedAt,
		UpdatedAt: group.UpdateAt,
	}

	return responseGroup
}

type GroupByPeriodLeaderResponse struct {
	ID        string  `json:"id,omitempty"`
	Report    string  `json:"report"`
	Potential string  `json:"potential"`
	Village   Village `json:"village"`
}

func ToGroupByPeriodLeaderResponse(group domain.Group) GroupByPeriodLeaderResponse {
	// village
	village := Village{
		ID:           group.Village.ID,
		Name:         group.Village.Name,
		Kecamatan:    group.Village.Kecamatan,
		Kabupaten:    group.Village.Kabupaten,
		Strength:     group.Village.Strength,
		Weakness:     group.Village.Weakness,
		Oportunities: group.Village.Oportunities,
		Threats:      group.Village.Threats,
	}

	// result response
	responseGroup := GroupByPeriodLeaderResponse{
		ID:        group.ID,
		Report:    group.Report,
		Potential: group.Potential,
		Village:   village,
	}

	return responseGroup
}

type GroupRegisteredByPeriodResponse struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Lecturer string  `json:"lecturer"`
	Village  Village `json:"village"`
}

func ToGroupRegisteredByPeriod(group domain.Group) GroupRegisteredByPeriodResponse {
	village := Village{
		Name:      group.Village.Name,
		Kecamatan: group.Village.Kecamatan,
		Kabupaten: group.Village.Kabupaten,
	}

	return GroupRegisteredByPeriodResponse{
		ID:       group.ID,
		Name:     group.Name,
		Village:  village,
		Lecturer: group.Lecturer.Name,
	}
}

type AddLecturerRequest struct {
	ID         string
	LecturerID string `json:"lecturer_id"`
}
