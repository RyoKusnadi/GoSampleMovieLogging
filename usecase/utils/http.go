package utils

import (
	"encoding/json"
	"net/http"
	"reflect"
	"scalable-go-movie/domain"
	"strings"
)

func CreateSuccessApiResponse(result interface{}) domain.ApiResponse {
	return domain.ApiResponse{
		Result: result,
	}
}

func CreateDefaultErrorApiResponse(code int, message string) domain.ApiResponse {
	return domain.ApiResponse{
		Error: &domain.ErrorResponse{
			Code:    code,
			Message: message,
		},
	}
}

func DecodeResponse(statusCode int, response []byte, successStruct interface{}, successExcludeFields map[string]bool, errorStruct interface{}, errorField string) domain.ApiResponse {
	if statusCode == http.StatusOK {
		if err := json.Unmarshal(response, &successStruct); err != nil {
			return CreateDefaultErrorApiResponse(http.StatusBadRequest, err.Error())
		}

		excludeFieldsByReflection(successStruct, successExcludeFields)
		return CreateSuccessApiResponse(successStruct)
	}

	if err := json.Unmarshal(response, &errorStruct); err != nil {
		return CreateDefaultErrorApiResponse(http.StatusInternalServerError, err.Error())
	}

	if errorMsg := getErrorMsg(errorStruct, errorField); errorMsg != "" {
		return CreateDefaultErrorApiResponse(statusCode, errorMsg)
	}
	return CreateDefaultErrorApiResponse(http.StatusInternalServerError, "unknown error when decode http msg")
}

func excludeFieldsByReflection(structPtr interface{}, excludeFields map[string]bool) {
	value := reflect.ValueOf(structPtr).Elem()
	typ := value.Type()

	for i := 0; i < value.NumField(); i++ {
		field := typ.Field(i)
		fieldName := strings.Replace(field.Tag.Get("json"), ",omitempty", "", -1)

		if excludeFields[fieldName] {
			zeroValue := reflect.Zero(field.Type)
			value.Field(i).Set(zeroValue)
		}
	}
}

func getErrorMsg(errorStruct interface{}, errorField string) string {
	value := reflect.ValueOf(errorStruct).Elem()
	typ := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := typ.Field(i)
		fieldName := strings.Replace(field.Tag.Get("json"), ",omitempty", "", -1)

		if fieldName == errorField {
			return value.Field(i).String()
		}
	}
	return ""
}
