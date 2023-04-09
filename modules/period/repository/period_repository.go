package repository

import (
	"backend-ekkn/modules/period/domain"
	"github.com/google/uuid"
)

type PeriodRepository interface {
	Create(period domain.Period) error
	FindAll() ([]domain.Period, error)
	FindById(id uuid.UUID) (domain.Period, error)
	Update(period domain.Period) error
	Delete(period domain.Period) error
}
