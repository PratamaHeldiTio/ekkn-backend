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

func (repo *StudentRepositoryImpl) Create(student domain.Student) (domain.Student, error) {
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

func (repo *StudentRepositoryImpl) FindAll() ([]domain.Student, error) {
	var students []domain.Student
	if err := repo.db.Find(&students).Error; err != nil {
		return students, err
	}

	return students, nil

}

// update data to db from data service
func (repo *StudentRepositoryImpl) Update(student domain.Student) (domain.Student, error) {
	if err := repo.db.Model(&student).Updates(student).Error; err != nil {
		return student, err
	} else {
		return student, nil
	}
}

func (repo *StudentRepositoryImpl) Delete(student domain.Student) error {
	if err := repo.db.Delete(&student).Error; err != nil {
		return err
	}

	return nil
}

