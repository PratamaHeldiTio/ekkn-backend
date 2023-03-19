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
	// insert data to database
	if err := repo.db.Create(&student).Error; err != nil {
		return student, err
	} else {
		return student, nil
	}
}

func (repo *StudentRepositoryImpl) FindByNim(nim string) (domain.Student, error) {
	var student domain.Student

	// select data from database
	if err := repo.db.Where("nim = ?", nim).Find(&student).Error; err != nil {
		return student, err
	}

	return student, nil
}
