package repository

import (
	"backend-ekkn/modules/student_registration/domain"
	"gorm.io/gorm"
)

type StudentRegistrationRepositoryImpl struct {
	db *gorm.DB
}

// init repo
func NewStudentRegistrationRepository(db *gorm.DB) StudentRegistrationRepository {
	return &StudentRegistrationRepositoryImpl{db}
}

// repo for create student registration
func (repo *StudentRegistrationRepositoryImpl) Create(registration domain.StudentRegistration) error {
	if err := repo.db.Create(&registration).Error; err != nil {
		return err
	}

	return nil
}
