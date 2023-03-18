package service

import (
	"backend-ekkn/modules/student/domain"
	"backend-ekkn/pkg/shareddomain"
)

type StudentService interface {
	CreateStudent(request shareddomain.CreateStudentRequest) (domain.Student, error)
}
