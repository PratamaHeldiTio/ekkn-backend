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
	UpdateGroup(request shareddomain.GroupUpdateRequest) error
	AddVillage(request shareddomain.AddVillage) error
	FindGroupByPeriodLeader(periodID, leader string) (domain.Group, error)
	FindRegisteredGroupByPeriod(ID, query string) ([]domain.Group, error)
	AddLecturer(request shareddomain.AddLecturerRequest) error
	FindGroupByPeriodLecturer(periodID, lecturerID string) ([]domain.Group, error)
}
