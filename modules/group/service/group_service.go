package service

import (
	"backend-ekkn/modules/group/domain"
	"backend-ekkn/pkg/shareddomain"
)

type GroupService interface {
	CreateGroup(request shareddomain.RequestGroup) error
	FindGroupID(ID string) (domain.Group, error)
	JoinGroup(studentID, periodID, refferal string) error
}