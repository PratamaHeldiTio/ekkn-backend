package shareddomain

import "backend-ekkn/modules/student/domain"

type CreateStudentRequest struct {
	Nim      string `json:"nim" binding:"required,min=1,max=13"`
	Name     string `json:"name" binding:"required,min=1,max=100"`
	Prodi    string `json:"prodi"`
	Fakultas string `json:"fakultas"`
}

type StudentResponse struct {
	Nim      string `json:"nim"`
	Name     string `json:"name"`
	Prodi    string `json:"prodi"`
	Fakultas string `json:"fakultas"`
	Token    string `json:"token"`
}

func ToResponseStudent(student domain.Student, token string) StudentResponse {
	studentResponse := StudentResponse{
		Nim:      student.Nim,
		Name:     student.Name,
		Prodi:    student.Prodi,
		Fakultas: student.Fakultas,
		Token:    token,
	}

	return studentResponse
}
