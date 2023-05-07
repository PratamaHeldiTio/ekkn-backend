package repository

import (
	"backend-ekkn/modules/period/domain"
)

type PeriodRepository interface {
	Create(period domain.Period) error
	FindAll() ([]domain.Period, error)
	FindById(id string) (domain.Period, error)
	Update(period domain.Period) error
	Delete(period domain.Period) error
}
