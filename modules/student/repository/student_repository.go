package repository

import "backend-ekkn/modules/student/domain"

type StudentRepository interface {
	Save(student domain.Student) (domain.Student, error)
	FindByNim(nim string) (domain.Student, error)
	FindAll() ([]domain.Student, error)
}
