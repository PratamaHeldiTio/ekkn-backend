package shareddomain

import "backend-ekkn/modules/lecturer/domain"

type LecturerRequest struct {
	ID         string `json:"id" binding:"required,max=18"`
	Name       string `json:"name" binding:"required,max=100"`
	Gender     string `json:"gender" binding:"max=9"`
	Prodi      string `json:"prodi" binding:"max=50"`
	Fakultas   string `json:"fakultas" binding:"max=50"`
	MaduraLang string `json:"madura_lang" binding:"max=5"`
	Contact    string `json:"contact" binding:"max=15"`
}

type LecturerResponse struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Gender     string `json:"gender"`
	Prodi      string `json:"prodi"`
	Fakultas   string `json:"fakultas"`
	MaduraLang string `json:"madura_lang"`
	Contact    string `json:"contact"`
}

type LecturerLogin struct {
	ID       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ChangePasswordLecturerRequest struct {
	ID                string
	OldPassword       string `json:"old_password" binding:"required"`
	NewPassword       string `json:"new_password" binding:"required"`
	RepeatNewPassword string `json:"repeat_new_password" binding:"required"`
}

func ToLecturerResponse(lecturer domain.Lecturer) LecturerResponse {
	lecturerResponse := LecturerResponse{
		ID:         lecturer.ID,
		Name:       lecturer.Name,
		Gender:     lecturer.Gender,
		Prodi:      lecturer.Prodi,
		Fakultas:   lecturer.Fakultas,
		MaduraLang: lecturer.MaduraLang,
		Contact:    lecturer.Contact,
	}

	return lecturerResponse
}
