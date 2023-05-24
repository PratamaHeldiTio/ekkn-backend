package repository

import "backend-ekkn/modules/output/domain"

type OutputRepository interface {
	Create(output domain.Output) error
	FindByID(ID string) (domain.Output, error)
	Update(output domain.Output) error
	FindByGroup(ID string) ([]domain.Output, error)
}
