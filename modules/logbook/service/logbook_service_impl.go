package service

import (
	service2 "backend-ekkn/modules/group/service"
	"backend-ekkn/modules/logbook/domain"
	"backend-ekkn/modules/logbook/repository"
	"backend-ekkn/modules/period/service"
	"backend-ekkn/pkg/helper"
	"backend-ekkn/pkg/shareddomain"
	"errors"
	"github.com/google/uuid"
	"strconv"
	"strings"
	"time"
)

type LogbookServiceImpl struct {
	repo          repository.LogbookRepository
	servicePeriod service.PeriodService
	groupService  service2.GroupService
}

func NewLogbookService(repo repository.LogbookRepository, servicePeriod service.PeriodService, groupService service2.GroupService) LogbookService {
	return &LogbookServiceImpl{repo, servicePeriod, groupService}
}

func (service *LogbookServiceImpl) CreateLogbook(request shareddomain.LogbookRequest) error {
	// find period
	//parse uuid string to uuid
	periodID, err := uuid.Parse(request.PeriodID)
	if err != nil {
		return err
	}

	period, err := service.servicePeriod.FindPeriodById(periodID)
	if err != nil {
		return err
	}

	// submitted
	submitted := time.Now().Unix()

	var arrayDate []int
	for _, date := range strings.Split(request.Date, "-") {
		value, err := strconv.Atoi(date)
		if err != nil {
			return err
		}
		arrayDate = append(arrayDate, value)
	}
	date := time.Date(arrayDate[0], time.Month(arrayDate[1]), arrayDate[2], 0, 0, 0, 0, time.Local).Unix()

	// check if student submit logbook on same date
	logbookByDate, err := service.repo.FindAllByStudentDate(request.StudentID, date)
	if err != nil {
		return err
	}

	if logbookByDate.ID != "" {
		return errors.New("anda telah mengisi logbook pada tanggal ini")
	}

	// check submitted > period start and submitted < period end  date < time.now
	if submitted < period.Start || submitted > period.End || submitted < date || date < period.Start {
		return errors.New("tanggal tidak sesuai")
	}

	// get coordinate village from group if
	group, err := service.groupService.FindGroupID(request.GroupID)
	if err != nil {
		return err
	}

	// definition struct coordinat
	studentLocation := helper.Coordinate{
		Latitude:  request.Latitude,
		Longitude: request.Longitude,
	}

	if group.Village.Latitude == 0 || group.Village.Longitude == 0 {
		return errors.New("Koordinat desa belum ditentukan")
	}

	villageLocation := helper.Coordinate{
		Latitude:  group.Village.Latitude,
		Longitude: group.Village.Longitude,
	}

	radius := helper.DistanceHarversine(studentLocation, villageLocation)

	logbook := domain.Logbook{
		PeriodID:  request.PeriodID,
		StudentID: request.StudentID,
		Activity:  request.Activity,
		Image:     request.Image,
		Radius:    radius,
		Date:      date,
		Submitted: submitted,
	}

	if err := service.repo.Create(logbook); err != nil {
		return err
	}

	return nil
}

func (service *LogbookServiceImpl) FindLogbookByStudentPeriod(studentID, periodID string) ([]domain.Logbook, error) {
	logbooks, err := service.repo.FindAllByStudentPeriod(studentID, periodID)
	if err != nil {
		return logbooks, err
	}

	return logbooks, nil
}
