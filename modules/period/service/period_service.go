package service

import (
	"backend-ekkn/modules/period/domain"
	"backend-ekkn/pkg/shareddomain"
)

type PeriodService interface {
	CreatePeriod(request shareddomain.RequestPeriod) error
	FindAllPeriod() ([]domain.Period, error)
	FindPeriodById(id string) (domain.Period, error)
	UpdatePeriod(request shareddomain.RequestPeriod) error
	DeletePeriodById(id string) error
}
