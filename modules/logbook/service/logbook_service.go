package service

import (
	"backend-ekkn/modules/logbook/domain"
	"backend-ekkn/pkg/shareddomain"
)

type LogbookService interface {
	CreateLogbook(request shareddomain.LogbookRequest) error
	FindLogbookByStudentPeriod(studentID, periodID string) ([]domain.Logbook, error)
}
