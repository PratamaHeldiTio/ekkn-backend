package service

import (
	"backend-ekkn/modules/logbook/domain"
	"backend-ekkn/modules/logbook/repository"
	"backend-ekkn/modules/period/service"
	"backend-ekkn/pkg/helper"
	"backend-ekkn/pkg/shareddomain"
)

type LogbookServiceImpl struct {
	repo          repository.LogbookRepository
	servicePeriod service.PeriodService
}

func NewLogbookService(repo repository.LogbookRepository, servicePeriod service.PeriodService) LogbookService {
	return &LogbookServiceImpl{repo, servicePeriod}
}

func (service *LogbookServiceImpl) CreateLogbook(request shareddomain.LogbookRequest) error {
	// find period
	//parse uuid string to uuid
	//periodID, err := uuid.Parse(request.PeriodID)
	//if err != nil {
	//	return err
	//}
	//period, err := service.servicePeriod.FindPeriodById(periodID)

	// submitted
	//submitted := time.Now().Unix()
	//date := time.Date(2023, 5, 2, 0, 0, 0, 0, time.Local).Unix()
	//
	//// check submitted > period start and submitted < period end  date < time.now
	//if submitted < period.Start || submitted > period.End || submitted < date {
	//	return errors.New("logbook gagal ditambahkan")
	//}

	// definition struct coordinat
	origin := helper.Coordinate{
		Latitude:  2.990353,
		Longitude: 101.533913,
	}
	destination := helper.Coordinate{
		Latitude:  2.960148,
		Longitude: 101.577888,
	}

	radius := helper.DistanceHarversine(origin, destination)

	logbook := domain.Logbook{
		PeriodID:  request.PeriodID,
		StudentID: request.StudentID,
		Activity:  request.Activity,
		Image:     request.Image,
		Radius:    radius,
		//Date:      date,
		//Submitted: submitted,
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
