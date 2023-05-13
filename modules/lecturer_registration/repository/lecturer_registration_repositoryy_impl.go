package repository

import (
	"backend-ekkn/modules/lecturer_registration/domain"
	"gorm.io/gorm"
)

type lecturerRegistrationRepositoryImpl struct {
	db *gorm.DB
}

func NewLecturerRegistrationRepository(db *gorm.DB) LecturerRegistrationRepository {
	return &lecturerRegistrationRepositoryImpl{db}
}

func (repo *lecturerRegistrationRepositoryImpl) Create(registration domain.LecturerRegistration) error {
	if err := repo.db.Create(&registration).Error; err != nil {
		return err
	}

	return nil
}

func (repo *lecturerRegistrationRepositoryImpl) FindByPeriodLectureID(periodID, lectureID string) (domain.LecturerRegistration, error) {
	var lecturerRegistration domain.LecturerRegistration

	if err := repo.db.Where("lecturer_id = ? and period_id = ?", lectureID, periodID).
		Find(&lecturerRegistration).Error; err != nil {
		return lecturerRegistration, err
	}

	return lecturerRegistration, nil
}

func (repo *lecturerRegistrationRepositoryImpl) FindByLectureID(lectureID string) ([]domain.LecturerRegistration, error) {
	var lecturerRegistration []domain.LecturerRegistration

	if err := repo.db.Preload("Period").Where("lecturer_id = ?", lectureID).
		Find(&lecturerRegistration).Error; err != nil {
		return lecturerRegistration, err
	}

	return lecturerRegistration, nil
}

func (repo *lecturerRegistrationRepositoryImpl) Update(registration domain.LecturerRegistration) error {
	if err := repo.db.Model(&registration).Updates(registration).Error; err != nil {
		return err
	}
	return nil
}

func (repo *lecturerRegistrationRepositoryImpl) FindByID(ID string) (domain.LecturerRegistration, error) {
	var lecturerRegistration domain.LecturerRegistration

	if err := repo.db.Where("id = ?", ID).
		Find(&lecturerRegistration).Error; err != nil {
		return lecturerRegistration, err
	}

	return lecturerRegistration, nil
}

func (repo *lecturerRegistrationRepositoryImpl) FindByPeriod(ID string) ([]domain.LecturerRegistration, error) {
	var lecturerRegistration []domain.LecturerRegistration

	if err := repo.db.Preload("Lecturer").Where("period_id = ?", ID).
		Find(&lecturerRegistration).Error; err != nil {
		return lecturerRegistration, err
	}

	return lecturerRegistration, nil
}
