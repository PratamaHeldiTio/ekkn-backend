package repository

import (
	"backend-ekkn/modules/period/domain"
)

type PeriodRepository interface {
	Create(period domain.Period) error
	//FindAll() ([]domain.Period, error)
	//FindById(id uuid.UUID) (domain.Period, error)
	//Update(period domain.Period) (domain.Period, error)
	//Delete(id uuid.UUID)
}
