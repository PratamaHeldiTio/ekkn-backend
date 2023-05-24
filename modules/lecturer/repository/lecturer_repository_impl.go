package repository

import (
	"backend-ekkn/modules/lecturer/domain"
	"gorm.io/gorm"
)

type LecturerRepositoryImpl struct {
	db *gorm.DB
}

func NewLectureRepositoryImpl(db *gorm.DB) LecturerRepository {
	return &LecturerRepositoryImpl{db}
}

func (repo *LecturerRepositoryImpl) Create(lecturer domain.Lecturer) error {
	if err := repo.db.Create(&lecturer).Error; err != nil {
		return err
	}

	return nil
}

func (repo *LecturerRepositoryImpl) Update(lecturer domain.Lecturer) error {
	if err := repo.db.Model(&lecturer).Updates(lecturer).Error; err != nil {
		return err
	}
	return nil
}

func (repo *LecturerRepositoryImpl) FindByID(ID string) (domain.Lecturer, error) {
	var lecturer domain.Lecturer
	if err := repo.db.Where("id = ?", ID).Find(&lecturer).Error; err != nil {
		return lecturer, err
	}

	return lecturer, nil
}

func (repo *LecturerRepositoryImpl) Delete(lecturer domain.Lecturer) error {
	if err := repo.db.Delete(&lecturer).Error; err != nil {
		return err
	}
	return nil
}

func (repo *LecturerRepositoryImpl) FindAll(query string) ([]domain.Lecturer, error) {
	var lecturer []domain.Lecturer
	if err := repo.db.Order("created_at asc").
		Where("id LIKE ?", "%"+query+"%").
		Or("name ILIKE ?", "%"+query+"%").
		Or("prodi ILIKE ?", "%"+query+"%").
		Or("fakultas ILIKE ?", "%"+query+"%").
		Find(&lecturer).Error; err != nil {
		return lecturer, err
	}

	return lecturer, nil
}
