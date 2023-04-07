package service

import (
	"backend-ekkn/pkg/shareddomain"
)

type PeriodService interface {
	CreatePeriod(request shareddomain.RequestPeriod) error
	//Update
}
