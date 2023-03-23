package repository

import "backend-ekkn/modules/student/domain"

type StudentRepository interface {
	Create(student domain.Student) (domain.Student, error)
	FindByNim(nim string) (domain.Student, error)
	FindAll() ([]domain.Student, error)
	Update(student domain.Student) (domain.Student, error)
	Delete(student domain.Student) error
}
