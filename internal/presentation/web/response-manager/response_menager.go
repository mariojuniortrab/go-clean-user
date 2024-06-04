package web_response_manager

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	errors_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/errors"
)

type messageText struct {
	Message string `json:"message"`
}

type messageSucessful struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type messageFieldErrorArray struct {
	Errors []*errors_validation.CustomFieldErrorMessage
}

type messageFieldError struct {
	Errors *errors_validation.CustomFieldErrorMessage
}
type messageError struct {
	Error *errors_validation.CustomErrorMessage
}

func setJsonContentTypeResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func RespondOk(w http.ResponseWriter, message string, data any) {
	setJsonContentTypeResponse(w)
	w.WriteHeader(http.StatusOK)
	if data != nil {
		json.NewEncoder(w).Encode(&messageSucessful{Message: message, Data: data})
	} else {
		json.NewEncoder(w).Encode(&messageText{Message: message})
	}
}

func RespondUiidIsRequired(w http.ResponseWriter) {
	setJsonContentTypeResponse(w)
	errorMessage := errors_validation.NewCustomErrorMessage(errors_validation.UiidFromPathIsRequired())

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(&messageError{Error: errorMessage})
}

func RespondInternalServerError(w http.ResponseWriter, err error) {
	setJsonContentTypeResponse(w)
	log.Print(err)
	errorMessage := errors_validation.NewCustomErrorMessage(errors_validation.InternalServerError())
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(&messageError{Error: errorMessage})
}

func RespondNotFound(w http.ResponseWriter, resource string) {
	setJsonContentTypeResponse(w)
	errorMessage := errors_validation.NewCustomErrorMessage(errors_validation.NotFound(resource))
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(&messageError{Error: errorMessage})
}

func RespondFieldErrorValidation(w http.ResponseWriter, errs []*errors_validation.CustomFieldErrorMessage) {
	setJsonContentTypeResponse(w)
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(&messageFieldErrorArray{Errors: errs})
}

func RespondLoginInvalid(w http.ResponseWriter) {
	setJsonContentTypeResponse(w)
	errorMessage := errors_validation.NewCustomErrorMessage(errors_validation.UserNotFound())
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(&messageError{Error: errorMessage})
}

func RespondConflictError(w http.ResponseWriter, errs *errors_validation.CustomFieldErrorMessage) {
	setJsonContentTypeResponse(w)
	w.WriteHeader(http.StatusConflict)
	json.NewEncoder(w).Encode(&messageFieldError{Errors: errs})
}

func RespondCreated(w http.ResponseWriter, message string, data any) {
	setJsonContentTypeResponse(w)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&messageSucessful{Message: message, Data: data})
}

func RespondGenericError(w http.ResponseWriter, statusCode int, message string) {
	setJsonContentTypeResponse(w)
	w.WriteHeader(statusCode)
	errorMessage := errors_validation.NewCustomErrorMessage(errors.New(message))
	json.NewEncoder(w).Encode(&messageError{Error: errorMessage})
}

func RespondUnsupportedMediaType(w http.ResponseWriter) {
	setJsonContentTypeResponse(w)
	w.WriteHeader(http.StatusCreated)
	errorMessage := errors_validation.NewCustomErrorMessage(errors_validation.ContentTypeIsNotJSON())
	json.NewEncoder(w).Encode(&messageError{Error: errorMessage})
}

func RespondUnauthorized(w http.ResponseWriter) {
	setJsonContentTypeResponse(w)
	w.WriteHeader(http.StatusCreated)
	errorMessage := errors_validation.NewCustomErrorMessage(errors_validation.Unauthorized())
	json.NewEncoder(w).Encode(&messageError{Error: errorMessage})
}

func RespondUiidInvalid(w http.ResponseWriter) {
	setJsonContentTypeResponse(w)
	w.WriteHeader(http.StatusCreated)
	errorMessage := errors_validation.NewCustomErrorMessage(errors_validation.Unauthorized())
	json.NewEncoder(w).Encode(&messageError{Error: errorMessage})
}
