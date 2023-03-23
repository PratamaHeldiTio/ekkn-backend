package service

import (
	"backend-ekkn/modules/student/domain"
	"backend-ekkn/pkg/shareddomain"
)

type StudentService interface {
	CreateStudent(request shareddomain.CreateStudentRequest) (domain.Student, error)
	FindStudentByNim(nim string) (domain.Student, error)
	LoginStudent(request shareddomain.LoginStudentRequest) (domain.Student, error)
	FindAllStudent() ([]domain.Student, error)
	UpdateStudent(request shareddomain.UpdateStudentRequest) (domain.Student, error)
}
