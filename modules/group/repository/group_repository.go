package repository

import (
	"backend-ekkn/modules/group/domain"
)

type GroupRepository interface {
	Create(group domain.Group) error
	FindByStudentPeriodID(studentID, periodID string) (domain.Group, error)
	FindByReferal(referral string) (domain.Group, error)
	Join(studentID, groupID, referral string) error
}
