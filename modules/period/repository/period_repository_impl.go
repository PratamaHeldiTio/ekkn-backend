package repository

import (
	"backend-ekkn/modules/period/domain"
	"gorm.io/gorm"
)

type PeriodRepositoryImpl struct {
	db *gorm.DB
}

func NewPeriodRepository(db *gorm.DB) PeriodRepository {
	return &PeriodRepositoryImpl{db}
}

func (repo *PeriodRepositoryImpl) Create(period domain.Period) error {
	if err := repo.db.Create(&period).Error; err != nil {
		return err
	}

	return nil
}

func (repo *PeriodRepositoryImpl) FindAll() ([]domain.Period, error) {
	// temp data
	var periods []domain.Period
	if err := repo.db.Order("created_at desc").Find(&periods).Error; err != nil {
		return periods, err
	}

	return periods, nil
}

func (repo *PeriodRepositoryImpl) FindById(id string) (domain.Period, error) {
	var period domain.Period
	if err := repo.db.Where("id = ?", id).Find(&period).Error; err != nil {
		return period, err
	}

	return period, nil
}

func (repo *PeriodRepositoryImpl) Update(period domain.Period) error {
	if err := repo.db.Model(&period).Updates(period).Error; err != nil {
		return err
	}

	return nil
}

func (repo *PeriodRepositoryImpl) Delete(period domain.Period) error {
	if err := repo.db.Delete(&period).Error; err != nil {
		return err
	}

	return nil
}
