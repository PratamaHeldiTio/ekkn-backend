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

//func (repo *PeriodRepositoryImpl) FindAll() ([]domain.Period, error) {
//	// temp data
//	var periods []domain.Period
//	if err := repo.db.Find(&periods).Error; err != nil {
//		return periods, err
//	} else {
//
//	}
//}
