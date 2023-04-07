package helper

type ResponseWithoutData struct {
	Code    int    `json:"code"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ResponseWithData struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseWithError struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

func APIResponseWithoutData(code int, success bool, message string) ResponseWithoutData {
	response := ResponseWithoutData{
		Code:    code,
		Success: success,
		Message: message,
	}
	return response
}

func APIResponseWithData(code int, success bool, message string, data interface{}) ResponseWithData {
	response := ResponseWithData{
		Code:    code,
		Success: success,
		Message: message,
		Data:    data,
	}
	return response
}

func APIResponseWithError(code int, success bool, message string, err interface{}) ResponseWithError {
	response := ResponseWithError{
		Code:    code,
		Success: success,
		Message: message,
		Error:   err,
	}
	return response
}
