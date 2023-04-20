package service

import (
	"backend-ekkn/modules/group/domain"
	"backend-ekkn/modules/group/repository"
	"backend-ekkn/modules/student_registration/service"
	"backend-ekkn/pkg/helper"
	"backend-ekkn/pkg/shareddomain"
	"errors"
)

type GroupServiceImpl struct {
	repo                       repository.GroupRepository
	studentRegistrationService service.StudentRegistrationService
}

func NewGroupServiceImpl(repo repository.GroupRepository, studentRegistrationService service.StudentRegistrationService) GroupService {
	return &GroupServiceImpl{repo, studentRegistrationService}
}

func (service *GroupServiceImpl) CreateGroup(request shareddomain.RequestGroup) error {
	// get student registration for validation status true and haven't group
	studentRegistration, err := service.studentRegistrationService.FindStudentRegistrationByNimPeriodID(request.Leader, request.PeriodID)
	if err != nil {
		return err
	}

	// cek student registration is valid and status true and have group
	if studentRegistration.ID == "" || studentRegistration.Status == "false" || studentRegistration.Group != "" {
		return errors.New("gagal membuat kelompok")
	}

	group := domain.Group{
		Name:     request.Name,
		PeriodID: request.PeriodID,
		Leader:   request.Leader,
		Referral: helper.RandomString(6),
	}
	if err := service.repo.Create(group); err != nil {
		return err
	}
	return nil
}

func (service *GroupServiceImpl) FindGroupID(ID string) (domain.Group, error) {
	// get group by student id and period id
	group, err := service.repo.FindByID(ID)
	if err != nil {
		return group, nil
	}

	// cek isExist
	if group.ID == "" {
		return group, errors.New("kelompok tidak dapat ditemukan")
	}

	return group, nil
}

func (service *GroupServiceImpl) JoinGroup(studentID, periodID, referral string) error {
	// get student registration for validation status true and haven't group
	studentRegistration, err := service.studentRegistrationService.FindStudentRegistrationByNimPeriodID(studentID, periodID)
	if err != nil {
		return err
	}

	// cek student registration is valid and status true and have group
	if studentRegistration.ID == "" || studentRegistration.Status == "false" || studentRegistration.Group != "" {
		return errors.New("gagal bergabung dengan kelompok")
	}

	// get group by referral
	group, err := service.repo.FindByReferal(referral, periodID)
	if err != nil {
		return err
	}

	// check referral valid
	if group.ID == "" {
		errors.New("referral anda salah")
	}

	if len(group.Students) >= 15 {
		return errors.New("kelompok sudah penuh")
	}

	// call join repo
	if err := service.repo.Join(studentID, periodID, group.ID); err != nil {
		return err
	}

	return nil
}
