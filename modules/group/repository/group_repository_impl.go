package repository

import (
	"backend-ekkn/modules/group/domain"
	studentRegistration "backend-ekkn/modules/student_registration/domain"
	"gorm.io/gorm"
)

type GroupRepositoryImpl struct {
	db *gorm.DB
}

func NewGroupRepository(db *gorm.DB) GroupRepository {
	return &GroupRepositoryImpl{db}
}

func (repo *GroupRepositoryImpl) Create(group domain.Group) error {
	//var registration studentRegistration.StudentRegistration
	result := repo.db.Create(&group)
	if result.Error != nil {
		return result.Error
	}

	var registration studentRegistration.StudentRegistration
	if err := repo.db.Model(&registration).
		Where("student_id = ? AND period_id = ?", group.Leader, group.PeriodID).
		Update("group", group.ID).Error; err != nil {
		return err
	}

	studentGroup := domain.StudentGroup{
		GroupID:    group.ID,
		StudentNim: group.Leader,
	}
	if err := repo.db.Create(&studentGroup).Error; err != nil {
		return err
	}

	return nil
}

func (repo *GroupRepositoryImpl) FindByStudentPeriodID(studentID, periodID string) (domain.Group, error) {
	var group domain.Group

	if err := repo.db.Preload("Students").
		Preload("Periods").
		Where("Periods.student_id = ? and period_id = ?", studentID, periodID).
		Find(&group).
		Error; err != nil {
		return group, err
	}

	return group, nil
}

func (repo *GroupRepositoryImpl) FindByReferal(referral string) (domain.Group, error) {
	var group domain.Group
	if err := repo.db.Preload("Students").
		Preload("Period").
		Where("referral = ?", referral).
		Find(&group).Error; err != nil {
		return group, err
	}

	return group, nil
}

func (repo *GroupRepositoryImpl) Join(studentID, PeriodID, groupID string) error {
	// update field group in registration student
	var registration studentRegistration.StudentRegistration
	if err := repo.db.Model(&registration).Where("student_id = ? and period_id = ?", studentID, PeriodID).
		Update("group", groupID).Error; err != nil {
		return err
	}

	studentGroup := domain.StudentGroup{
		GroupID:    groupID,
		StudentNim: studentID,
	}

	if err := repo.db.Create(&studentGroup).Error; err != nil {
		return err
	}

	return nil
}
