package shareddomain

import "backend-ekkn/modules/student/domain"

type CreateStudent struct {
	Nim      string `json:"nim" binding:"required,max=13"`
	Name     string `json:"name" binding:"required,max=100"`
	Prodi    string `json:"prodi" binding:"max=50"`
	Fakultas string `json:"fakultas" binding:"max=50"`
}

type UpdateStudent struct {
	Nim        string `json:"nim"`
	Name       string `json:"name" binding:"required,max=100"`
	Prodi      string `json:"prodi" binding:"max=50"`
	Fakultas   string `json:"fakultas" binding:"max=50"`
	Gender     string `json:"gender" binding:"max=9"`
	MaduraLang bool   `json:"madura_lang" binding:"boolean"`
}

type LoginStudent struct {
	Nim      string `json:"nim" binding:"required,max=13"`
	Password string `json:"password" binding:"required"`
}

type FindStudentByNim struct {
	Nim        string `json:"nim"`
	Name       string `json:"name"`
	Prodi      string `json:"prodi"`
	Fakultas   string `json:"fakultas"`
	Gender     string `json:"gender"`
	Position   string `json:"position"`
	MaduraLang bool   `json:"madura_lang"`
	GroupKkn   string `json:"groupKkn"`
	Grade      string `json:"grade"`
	CreatedAt  int64  `json:"created_at"`
	UpdateAt   int64  `json:"update_at"`
}

func ToResponseStudent(student domain.Student) CreateStudent {
	studentResponse := CreateStudent{
		Nim:      student.Nim,
		Name:     student.Name,
		Prodi:    student.Prodi,
		Fakultas: student.Fakultas,
	}

	return studentResponse

}

func ToResponseFindStudentByNim(student domain.Student) FindStudentByNim {
	studentResponse := FindStudentByNim{
		Nim:        student.Nim,
		Name:       student.Name,
		Prodi:      student.Prodi,
		Fakultas:   student.Fakultas,
		Gender:     student.Gender,
		Position:   student.Position,
		MaduraLang: student.MaduraLang,
		GroupKkn:   student.GroupKkn,
		Grade:      student.Grade,
		CreatedAt:  student.CreatedAt,
		UpdateAt:   student.UpdateAt,
	}

	return studentResponse
}
