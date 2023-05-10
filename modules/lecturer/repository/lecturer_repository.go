package repository

import "backend-ekkn/modules/lecturer/domain"

type LecturerRepository interface {
	Create(lecturer domain.Lecturer) error
	FindByID(ID string) (domain.Lecturer, error)
	FindAll() ([]domain.Lecturer, error)
	Update(lecturer domain.Lecturer) error
	Delete(lecturer domain.Lecturer) error
}
