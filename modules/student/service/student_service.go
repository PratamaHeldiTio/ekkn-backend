package service

import (
	"backend-ekkn/modules/student/domain"
	"backend-ekkn/pkg/shareddomain"
)

type StudentService interface {
	CreateStudent(request shareddomain.CreateStudent) (domain.Student, error)
	FindStudentByNim(nim string) (domain.Student, error)
	LoginStudent(request shareddomain.LoginStudent) (domain.Student, error)
	FindAllStudent() ([]domain.Student, error)
	UpdateStudent(request shareddomain.UpdateStudent) (domain.Student, error)
}
