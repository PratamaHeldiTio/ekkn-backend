package repository

import (
	"backend-ekkn/modules/output/domain"
	"gorm.io/gorm"
)

type OutputRepositoryImpl struct {
	db *gorm.DB
}

func NewOutputRepository(db *gorm.DB) OutputRepository {
	return &OutputRepositoryImpl{db}
}

func (repo *OutputRepositoryImpl) Create(output domain.Output) error {
	if err := repo.db.Create(&output).Error; err != nil {
		return err
	}

	return nil
}

func (repo *OutputRepositoryImpl) FindByID(ID string) (domain.Output, error) {
	var output domain.Output
	if err := repo.db.Where("id = ?", ID).Find(&output).Error; err != nil {
		return output, err
	}

	return output, nil
}

func (repo *OutputRepositoryImpl) Update(output domain.Output) error {
	if err := repo.db.Model(&output).Updates(output).Error; err != nil {
		return err
	}
	return nil
}

func (repo *OutputRepositoryImpl) FindByGroup(ID string) ([]domain.Output, error) {
	var outputs []domain.Output

	if err := repo.db.Where("group_id = ?", ID).Find(&outputs).Error; err != nil {
		return outputs, err
	}

	return outputs, nil
}
