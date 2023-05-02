package repository

import "backend-ekkn/modules/logbook/domain"

type LogbookRepository interface {
	Create(logbook domain.Logbook) error
	FindAllByStudentPeriod(studentID, periodID string) ([]domain.Logbook, error)
}
