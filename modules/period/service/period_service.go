package service

import (
	"backend-ekkn/modules/period/domain"
	"backend-ekkn/pkg/shareddomain"
	"github.com/google/uuid"
)

type PeriodService interface {
	CreatePeriod(request shareddomain.RequestPeriod) error
	FindAllPeriod() ([]domain.Period, error)
	FindPeriodById(id uuid.UUID) (domain.Period, error)
	UpdatePeriod(request shareddomain.RequestPeriod) error
}
