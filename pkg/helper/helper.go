package helper

type ResponseSuccess struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type ResponseFail struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Error   interface{} `json:"error"`
}

func APIResponseSuccess(code int, success bool, data interface{}) ResponseSuccess {
	response := ResponseSuccess{
		Code:    code,
		Success: success,
		Data:    data,
	}
	return response
}

func APIResponseFail(code int, success bool, err interface{}) ResponseFail {
	response := ResponseFail{
		Code:    code,
		Success: success,
		Error:   err,
	}
	return response
}
