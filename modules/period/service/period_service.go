package service

import (
	"backend-ekkn/modules/period/domain"
	"backend-ekkn/pkg/shareddomain"
)

type PeriodService interface {
	CreatePeriod(request shareddomain.RequestPeriod) error
	FindAllPeriod() ([]domain.Period, error)
	//Update
}
