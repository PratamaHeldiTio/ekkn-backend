package repository

import (
	"backend-ekkn/modules/logbook/domain"
	"gorm.io/gorm"
)

type LogbookRespositoryImpl struct {
	db *gorm.DB
}

func NewLogbookRepository(db *gorm.DB) LogbookRepository {
	return &LogbookRespositoryImpl{db}
}

func (repo *LogbookRespositoryImpl) Create(logbook domain.Logbook) error {
	if err := repo.db.Create(&logbook).Error; err != nil {
		return err
	}

	return nil
}

func (repo *LogbookRespositoryImpl) FindAllByStudentPeriod(studentID, periodID string) ([]domain.Logbook, error) {
	var logbooks []domain.Logbook
	if err := repo.db.Where("student_id = ? and period_id = ?", studentID, periodID).
		Find(&logbooks).Error; err != nil {
		return logbooks, err
	}

	return logbooks, nil
}
