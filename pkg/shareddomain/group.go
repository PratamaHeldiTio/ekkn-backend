package shareddomain

import "backend-ekkn/modules/group/domain"

type RequestGroup struct {
	Name     string `json:"name" binding:"required,max=100"`
	PeriodID string `json:"period_id"`
	Leader   string `json:"leader"`
}

type Student struct {
	Nim        string `json:"nim"`
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Prodi      string `json:"prodi"`
	Fakultas   string `json:"fakultas"`
	MaduraLang string `json:"madura_lang"`
}

type ResponseGroupByStudentPeriodID struct {
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

func ToResponseGroupByStudentPeriodID(group domain.Group) ResponseGroupByStudentPeriodID {
	// maping students
	var students []Student
	for _, student := range group.Students {
		students = append(students, Student{
			Nim:        student.Nim,
			Name:       student.Name,
			Gender:     student.Gender,
			Prodi:      student.Prodi,
			Fakultas:   student.Fakultas,
			MaduraLang: student.MaduraLang,
		})
	}

	// result response
	responseGroup := ResponseGroupByStudentPeriodID{
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
