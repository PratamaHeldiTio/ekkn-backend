package helper

import (
	"errors"
	"github.com/gin-gonic/gin"
	"math"
	"math/rand"
	"mime/multipart"
	"strconv"
	"strings"
	"time"
)

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

func SavePDF(c *gin.Context, name string) (string, error) {
	file, err := c.FormFile(name)
	if err != nil {
		return "", err
	}

	if file.Size > 10485760 {
		return "", errors.New("file terlalu besar")
	}

	if file.Header.Values("Content-Type")[0] != "application/pdf" {
		return "", errors.New("format file salah")
	}

	// save file to directory
	filename := strconv.FormatInt(time.Now().UnixMilli(), 10) + "_" + file.Filename
	path := "public/" + name + "/" + filename
	if err := c.SaveUploadedFile(file, path); err != nil {
		return "", err
	}

	return filename, nil
}

func degreToRadian(degre float64) float64 {
	return degre * math.Pi / 180
}

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

func DistanceHarversine(origin, destination Coordinate) int {
	// convert degre to radian for trigono
	origin.Latitude = degreToRadian(origin.Latitude)
	origin.Longitude = degreToRadian(origin.Longitude)
	destination.Latitude = degreToRadian(destination.Latitude)
	destination.Longitude = degreToRadian(destination.Longitude)

	// create different 2 coordinate
	differentLatitude := origin.Latitude - destination.Latitude
	differentLongitude := origin.Longitude - destination.Longitude

	innerBlock := math.Pow(math.Sin(differentLatitude/2), 2) + math.Cos(origin.Latitude)*math.Cos(destination.Latitude)*math.Pow(math.Sin(differentLongitude/2), 2)

	result := 2 * 6371 * math.Asin(math.Sqrt(innerBlock)) * 1000
	return int(result) // its meter value integer
}

func SaveImage(c *gin.Context, image *multipart.FileHeader, destination string) (string, error) {
	if image.Size > 1048576 {
		return "", errors.New("file terlalu besar")
	}

	if !strings.Contains(image.Header.Values("Content-Type")[0], "image") {
		return "", errors.New("format file salah")
	}

	// save file to directory
	filename := strconv.FormatInt(time.Now().UnixMilli(), 10) + "_" + image.Filename
	path := "public/" + destination + "/" + filename
	if err := c.SaveUploadedFile(image, path); err != nil {
		return "", err
	}

	return filename, nil
}

func StringDateToArray(date string, day string) (int64, error) {
	var arrayDate []int
	var unixDate int64
	for _, date := range strings.Split(date, "-") {
		value, err := strconv.Atoi(date)
		if err != nil {
			return unixDate, err
		}
		arrayDate = append(arrayDate, value)
	}

	if day == "start" {
		unixDate = time.Date(arrayDate[0], time.Month(arrayDate[1]), arrayDate[2], 0, 0, 0, 0, time.Local).Unix()
	} else if day == "end" {
		unixDate = time.Date(arrayDate[0], time.Month(arrayDate[1]), arrayDate[2], 23, 59, 59, 0, time.Local).Unix()
	}

	return unixDate, nil
}
