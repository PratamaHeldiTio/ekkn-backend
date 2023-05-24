package service

import (
	"backend-ekkn/modules/output/domain"
	"backend-ekkn/modules/output/repository"
	"backend-ekkn/pkg/shareddomain"
	"errors"
)

type OutputServiceImpl struct {
	repo repository.OutputRepository
}

func NewOutputService(repo repository.OutputRepository) OutputService {
	return &OutputServiceImpl{repo}
}

func (service *OutputServiceImpl) CreateOutput(request shareddomain.OutputRequest) error {
	output := domain.Output{
		GroupID:      request.GroupID,
		Type:         request.Type,
		File:         request.File,
		Description:  request.Description,
		Contribution: request.Contribution,
	}

	if err := service.repo.Create(output); err != nil {
		return err
	}

	return nil
}

func (service *OutputServiceImpl) FindOutputByID(ID string) (domain.Output, error) {
	output, err := service.repo.FindByID(ID)
	if err != nil {
		return output, err
	}

	if output.ID == "" {
		return output, errors.New("data tidak ditemukan")
	}

	return output, nil
}

func (service *OutputServiceImpl) UpdateOutput(request shareddomain.UpdateOutputRequest) error {
	output, err := service.FindOutputByID(request.ID)
	if err != nil {
		return err
	}

	output.File = request.File
	output.Type = request.Type
	output.Contribution = request.Contribution
	output.Description = request.Description
	if err := service.repo.Update(output); err != nil {
		return err
	}

	return nil
}

func (service *OutputServiceImpl) FindOutputByGroup(ID string) ([]domain.Output, error) {
	outputs, err := service.repo.FindByGroup(ID)
	if err != nil {
		return outputs, err
	}

	return outputs, nil
}
