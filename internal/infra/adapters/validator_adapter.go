package infra_adapters

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	protocol_application "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol/application"
	errors_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/errors"
)

type validatorAdapter struct {
	validator *validator.Validate
	errors    []*errors_validation.CustomFieldErrorMessage
}

func NewValidator() *validatorAdapter {
	return &validatorAdapter{
		validator: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (v *validatorAdapter) HasErrors() bool {
	return len(v.errors) > 0
}

func (v *validatorAdapter) AddError(err error, fieldName string) protocol_application.Validator {
	v.errors = append(v.errors, errors_validation.NewCustomFieldErrorMessage(err, fieldName))
	return v
}

func (v *validatorAdapter) GetErrors() []*errors_validation.CustomFieldErrorMessage {
	return v.errors
}

func (v *validatorAdapter) GetErrorsAndClean() []*errors_validation.CustomFieldErrorMessage {
	errors := v.errors
	v.errors = []*errors_validation.CustomFieldErrorMessage{}

	return errors
}

func (v *validatorAdapter) ValidateRequiredField(f interface{}, fieldName string) protocol_application.Validator {
	return v.defaultValidation(f, fieldName, "required", errors_validation.IsRequired)
}

func (v *validatorAdapter) ValidateEmailField(f interface{}, fieldName string) protocol_application.Validator {
	return v.defaultValidation(f, fieldName, "omitempty,email", errors_validation.IsInvalid)
}

func (v *validatorAdapter) ValidatePasswordConfirmationEquals(password, passwordConfirmation string) protocol_application.Validator {
	fn := errors_validation.IsPasswordConfirmationInvalid
	return v.defaultFieldCompareValidation(password, passwordConfirmation, "passwordConfirmation", fn)
}

func (v *validatorAdapter) ValidateStringField(f interface{}, fieldName string) protocol_application.Validator {
	return v.defaultValidation(f, fieldName, "omitempty,alphanumunicode", errors_validation.MustBeString)
}

func (v *validatorAdapter) ValidateNumberField(f interface{}, fieldName string) protocol_application.Validator {
	return v.defaultValidation(f, fieldName, "omitempty,numeric", errors_validation.MustBeNumeric)
}

func (v *validatorAdapter) ValidateStringBooleanField(f interface{}, fieldName string) protocol_application.Validator {
	return v.defaultValidation(f, fieldName, "omitempty,oneof=true false", errors_validation.MustBeString)
}

func (v *validatorAdapter) ValidateFieldLength(f interface{}, fieldName string, length int) protocol_application.Validator {
	return v.defaultLenghValidation(f, fieldName, fmt.Sprintf("omitempty,len=%d", length), errors_validation.LengthMustBe(fieldName, length))
}

func (v *validatorAdapter) ValidateFieldMaxLength(f interface{}, fieldName string, length int) protocol_application.Validator {
	return v.defaultLenghValidation(f, fieldName, fmt.Sprintf("omitempty,max=%d", length), errors_validation.LengthMustBeOrLess(fieldName, length))
}

func (v *validatorAdapter) ValidateFieldMinLength(f interface{}, fieldName string, length int) protocol_application.Validator {
	return v.defaultLenghValidation(f, fieldName, fmt.Sprintf("omitempty,min=%d", length), errors_validation.LengthMustBeAtLeast(fieldName, length))
}

func (v *validatorAdapter) defaultLenghValidation(f interface{}, fieldName, flag string, errMessage error) protocol_application.Validator {
	err := v.validator.Var(f, flag)

	if err != nil {
		v.errors = append(v.errors, errors_validation.NewCustomFieldErrorMessage(errMessage, fieldName))
	}

	return v
}

func (v *validatorAdapter) defaultValidation(f interface{}, fieldName, flag string, fn func(string) error) protocol_application.Validator {
	err := v.validator.Var(f, flag)

	if err != nil {
		v.errors = append(v.errors, errors_validation.NewCustomFieldErrorMessage(fn(fieldName), fieldName))
	}

	return v
}

func (v *validatorAdapter) defaultFieldCompareValidation(f1, f2, fieldName string, fn func() error) protocol_application.Validator {
	err := v.validator.VarWithValue(f1, f2, "eqfield")

	if err != nil {
		v.errors = append(v.errors, errors_validation.NewCustomFieldErrorMessage(fn(), fieldName))
	}

	return v
}
