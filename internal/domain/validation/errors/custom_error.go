package errors_validation

type CustomErrorMessage struct {
	Message string
}

type CustomFieldErrorMessage struct {
	CustomErrorMessage
	Field string
}

func NewCustomFieldErrorMessage(err error, field string) *CustomFieldErrorMessage {
	errorMessage := &CustomFieldErrorMessage{Field: field}
	errorMessage.Message = err.Error()
	return errorMessage
}

func NewCustomErrorMessage(err error) *CustomErrorMessage {
	errorMessage := &CustomErrorMessage{}
	errorMessage.Message = err.Error()
	return errorMessage
}
