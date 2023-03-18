package repository

import (
	"backend-ekkn/modules/student/domain"
	"gorm.io/gorm"
)

type StudentRepositoryImpl struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &StudentRepositoryImpl{db}
}

func (repo *StudentRepositoryImpl) Save(student domain.Student) (domain.Student, error) {
	if err := repo.db.Create(&student).Error; err != nil {
		return student, err
	} else {
		return student, nil
	}
}
