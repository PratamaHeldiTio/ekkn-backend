package service

import (
	"backend-ekkn/modules/group/domain"
	"backend-ekkn/modules/group/repository"
	"backend-ekkn/modules/student_registration/service"
	service2 "backend-ekkn/modules/village/service"
	"backend-ekkn/pkg/helper"
	"backend-ekkn/pkg/shareddomain"
	"errors"
)

type GroupServiceImpl struct {
	repo                       repository.GroupRepository
	studentRegistrationService service.StudentRegistrationService
	villageService             service2.VillageService
}

func NewGroupServiceImpl(
	repo repository.GroupRepository,
	studentRegistrationService service.StudentRegistrationService,
	villageService service2.VillageService) GroupService {
	return &GroupServiceImpl{
		repo,
		studentRegistrationService,
		villageService}
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

func (service *GroupServiceImpl) RegisterGroup(ID, Nim string) error {
	//find group
	group, err := service.FindGroupID(ID)
	if err != nil {
		return err
	}

	//status group registration must be true register only leader and min member 12
	// minimal can madura lang 1 and min 3 prodi
	maduraMember := 0
	var prodiStudents []string
	for _, student := range group.Students {
		if student.MaduraLang == "true" {
			maduraMember += 1
		}
		prodiStudents = append(prodiStudents, student.Prodi)
	}
	IdenticProdi := len(helper.UniqueSlice(prodiStudents))

	if group.Period.StatusRegisterGroup == "false" || group.Leader != Nim ||
		len(group.Students) < 12 || maduraMember < 1 || IdenticProdi < 3 {
		return errors.New("gagal mendaftarkan kelompok")
	}

	updateGroup := domain.Group{
		ID:     ID,
		Status: "true",
	}

	if err := service.repo.Update(updateGroup); err != nil {
		return err
	}

	return nil

}

func (service *GroupServiceImpl) UpdateGroup(request shareddomain.GroupUpdateRequest) error {
	//find group
	group, err := service.FindGroupID(request.ID)
	if err != nil {
		return err
	}

	if group.Leader != request.Nim {
		return errors.New("gagal memperbaharui kelompok")
	}

	// change new value update
	group.VillageID = request.Village
	group.Report = request.Report
	group.Potential = request.Potential

	if err := service.repo.Update(group); err != nil {
		return err
	}

	return nil
}

func (service *GroupServiceImpl) AddVillage(request shareddomain.AddVillage) error {
	//find group
	group, err := service.FindGroupID(request.ID)
	if err != nil {
		return err
	}

	if group.Leader != request.Nim || group.Status != "true" || group.VillageID != "" {
		return errors.New("gagal menambahkan desa")
	}

	// get villa check status
	village, err := service.villageService.FindVillageById(request.Village)
	if err != nil {
		return err
	}

	if village.Status == "true" {
		return errors.New("gagal menambahkan desa")
	}

	// update group add village id
	group.VillageID = request.Village
	if err := service.repo.Update(group); err != nil {
		return err
	}

	// update status village
	requestVillage := shareddomain.UpdateVillageRequest{
		ID:     request.Village,
		Status: "true",
	}
	if err := service.villageService.UpdateVillage(requestVillage); err != nil {
		return err
	}

	return nil
}

func (service *GroupServiceImpl) FindGroupByPeriodLeader(periodID, leader string) (domain.Group, error) {
	// get group by student id and period id
	group, err := service.repo.FindByPeriodLeader(periodID, leader)
	if err != nil {
		return group, nil
	}

	// cek isExist
	if group.ID == "" {
		return group, errors.New("kelompok tidak dapat ditemukan")
	}

	return group, nil
}

func (service *GroupServiceImpl) FindRegisteredGroupByPeriod(ID string) ([]domain.Group, error) {
	groups, err := service.repo.FindByPeriod(ID)
	if err != nil {
		return groups, err
	}

	return groups, nil
}

func (service *GroupServiceImpl) AddLecturer(request shareddomain.AddLecturerRequest) error {
	//find group
	group, err := service.FindGroupID(request.ID)
	if err != nil {
		return err
	}

	// change new value update
	group.LecturerID = request.LecturerID
	if err := service.repo.Update(group); err != nil {
		return err
	}

	return nil
}

func (service *GroupServiceImpl) FindGroupByPeriodLecturer(periodID, lecturerID string) ([]domain.Group, error) {
	groups, err := service.repo.FindByPeriodLecturer(periodID, lecturerID)
	if err != nil {
		return groups, err
	}

	return groups, nil
}
