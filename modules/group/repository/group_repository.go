package repository

import (
	"backend-ekkn/modules/group/domain"
)

type GroupRepository interface {
	Create(group domain.Group) error
	FindByReferal(referral, periodID string) (domain.Group, error)
	Join(studentID, groupID, referral string) error
	FindByID(ID string) (domain.Group, error)
}
