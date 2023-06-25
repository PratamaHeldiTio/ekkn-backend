package service

import (
	"backend-ekkn/modules/student/domain"
	"backend-ekkn/pkg/shareddomain"
)

type StudentService interface {
	CreateStudent(request shareddomain.CreateStudent) error
	FindStudentByNim(nim string) (domain.Student, error)
	LoginStudent(request shareddomain.LoginStudent) (domain.Student, error)
	FindAllStudent(query string) ([]domain.Student, error)
	UpdateStudent(request shareddomain.UpdateStudent) error
	DeleteStudent(nim string) error
	ChangePassword(request shareddomain.ChangePasswordRequest) error
	ResetPassword(studentID string) error
	UploadProfile(ID, filename string) error
	ImportStudents(request shareddomain.ImportStudent) error
}
