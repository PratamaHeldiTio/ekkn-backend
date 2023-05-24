package service

import (
	"backend-ekkn/modules/output/domain"
	"backend-ekkn/pkg/shareddomain"
)

type OutputService interface {
	CreateOutput(request shareddomain.OutputRequest) error
	FindOutputByID(ID string) (domain.Output, error)
	UpdateOutput(request shareddomain.UpdateOutputRequest) error
	FindOutputByGroup(ID string) ([]domain.Output, error)
}
