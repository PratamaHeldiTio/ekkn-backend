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

func (repo *StudentRegistrationRepositoryImpl) FindByStudentId(id string) ([]domain.StudentRegistration, error) {
	var registeredUser []domain.StudentRegistration

	if err := repo.db.Preload("Period").
		Order("created_at desc").
		Preload("Student").
		Where("student_id = ?", id).
		Find(&registeredUser).
		Error; err != nil {
		return registeredUser, err
	}

	return registeredUser, nil
}

func (repo *StudentRegistrationRepositoryImpl) FindByNimPeriodId(nim string, periodId string) (domain.StudentRegistration, error) {
	var registeredUser domain.StudentRegistration

	if err := repo.db.Preload("Period").
		Preload("Student").
		Where("student_id = ? and period_id = ?", nim, periodId).
		Find(&registeredUser).
		Error; err != nil {
		return registeredUser, err
	}

	return registeredUser, nil
}

func (repo *StudentRegistrationRepositoryImpl) FindByPeriod(periodID string) ([]domain.StudentRegistration, error) {
	var studentRegistration []domain.StudentRegistration

	if err := repo.db.Preload("Student").
		Order("student_id").
		Where("period_id = ?", periodID).
		Find(&studentRegistration).
		Error; err != nil {
		return studentRegistration, err
	}

	return studentRegistration, nil
}

func (repo *StudentRegistrationRepositoryImpl) Update(registration domain.StudentRegistration) error {
	if err := repo.db.Model(&registration).Updates(registration).Error; err != nil {
		return err
	}
	return nil
}

func (repo *StudentRegistrationRepositoryImpl) FindByID(ID string) (domain.StudentRegistration, error) {
	var registration domain.StudentRegistration
	if err := repo.db.Where("student_registration_id = ?", ID).Find(&registration).Error; err != nil {
		return registration, err
	}

	return registration, nil
}
