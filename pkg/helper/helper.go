package helper

import "math/rand"

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

// define the given charset, char only
var charset = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

// n is the length of random string we want to generate
func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		// randomly select 1 character from given charset
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// func to get unique slice

func UniqueSlice(input []string) []string {
	unique := make([]string, 0, len(input))
	mapBool := make(map[string]bool)

	for _, val := range input {
		if _, ok := mapBool[val]; !ok {
			mapBool[val] = true
			unique = append(unique, val)
		}
	}

	return unique
}
