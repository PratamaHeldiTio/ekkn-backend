package shareddomain

import "backend-ekkn/modules/output/domain"

type OutputRequest struct {
	ID           string
	GroupID      string `json:"group_id" binding:"required"`
	Description  string `json:"description" binding:"required"`
	Type         string `json:"type" binding:"required"`
	File         string `json:"file" binding:"required"`
	Contribution string `json:"contribution" binding:"required"`
}

type UpdateOutputRequest struct {
	ID           string
	Type         string `json:"type"`
	File         string `json:"file"`
	Description  string `json:"description"`
	Contribution string `json:"contribution"`
}

type OutputResponse struct {
	ID           string `json:"id"`
	Type         string `json:"type"`
	File         string `json:"file"`
	Description  string `json:"description"`
	Contribution string `json:"contribution"`
}

func ToOutputResponse(output domain.Output) OutputResponse {
	return OutputResponse{
		ID:           output.ID,
		Type:         output.Type,
		File:         output.File,
		Description:  output.Description,
		Contribution: output.Contribution,
	}
}
