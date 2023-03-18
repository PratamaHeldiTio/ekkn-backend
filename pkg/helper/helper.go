package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func APIResponse(code int, success bool, message string, data interface{}) Response {
	meta := Meta{
		Code:    code,
		Success: success,
		Message: message,
	}

	response := Response{
		Meta: meta,
		Data: data,
	}
	return response
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, err.Error())
	}

	return errors
}
