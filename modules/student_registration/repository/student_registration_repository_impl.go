package repository

import (
	"backend-ekkn/modules/student_registration/domain"
	"github.com/google/uuid"
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

func (repo *StudentRegistrationRepositoryImpl) FindByStudentId(id string) ([]domain.StudentRegistration, error) {
	var registeredUser []domain.StudentRegistration

	if err := repo.db.Preload("Period").
		Preload("Student").
		Where("student_id = ?", id).
		Find(&registeredUser).
		Error; err != nil {
		return registeredUser, err
	}

	return registeredUser, nil
}

func (repo *StudentRegistrationRepositoryImpl) FindByStudentIdPeriodId(studentId string, periodId uuid.UUID) (domain.StudentRegistration, error) {
	var registeredUser domain.StudentRegistration

	if err := repo.db.Preload("Period").
		Preload("Student").
		Where("student_id = ?", studentId).
		Where("period_id = ?", periodId).
		Find(&registeredUser).
		Error; err != nil {
		return registeredUser, err
	}

	return registeredUser, nil
}
