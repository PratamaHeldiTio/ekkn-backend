package repository

import (
	"backend-ekkn/modules/group/domain"
)

type GroupRepository interface {
	Create(group domain.Group) error
	FindByReferal(referral, periodID string) (domain.Group, error)
	Join(studentID, groupID, referral string) error
	FindByID(ID string) (domain.Group, error)
	FindByPeriodLeader(periodID, leader string) (domain.Group, error)
	Update(group domain.Group) error
	FindByPeriod(ID string) ([]domain.Group, error)
	FindByPeriodLecturer(periodID, lecturerID string) ([]domain.Group, error)
}
