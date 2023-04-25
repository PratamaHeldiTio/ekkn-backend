package service

import (
	"backend-ekkn/modules/group/domain"
	"backend-ekkn/pkg/shareddomain"
)

type GroupService interface {
	CreateGroup(request shareddomain.RequestGroup) error
	FindGroupID(ID string) (domain.Group, error)
	JoinGroup(studentID, periodID, refferal string) error
	RegisterGroup(ID, Nim string) error
	UpdateGroup(request shareddomain.RequestGroupUpdate) error
	AddVillage(request shareddomain.AddVillage) error
}
