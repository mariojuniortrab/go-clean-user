package user_validation

import (
	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	util_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/util"
	protocol_application "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol/application"
	errors_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/errors"
)

type UpdateValidation interface {
	Validate(input *user_entity.UserUpdateInputDto) []*errors_validation.CustomFieldErrorMessage
}
type updateValidation struct {
	validator protocol_application.Validator
}

func NewUpdateValidation(validator protocol_application.Validator) *updateValidation {
	return &updateValidation{
		validator,
	}
}

func (v *updateValidation) Validate(input *user_entity.UserUpdateInputDto) []*errors_validation.CustomFieldErrorMessage {
	if input.Password != nil {
		v.validatePassword(*input.Password)
	}

	if input.Name != nil {
		v.validateName(*input.Name)
	}

	if input.Birth != nil {
		v.validateBirth(*input.Birth)
	}

	if input.PasswordConfirmation != nil && input.Password != nil {
		v.validatePasswordConfirmation(*input.PasswordConfirmation, *input.Password)
	}

	if input.PasswordConfirmation != nil && input.Password == nil {
		v.validator.AddError(errors_validation.IsRequired("password"), "password")
	}

	if input.PasswordConfirmation == nil && input.Password != nil {
		v.validator.AddError(errors_validation.IsRequired("passwordConfirmation"), "passwordConfirmation")
	}

	if v.validator.HasErrors() {
		return v.validator.GetErrorsAndClean()
	}

	return nil
}

func (v *updateValidation) validatePassword(input string) {
	const fieldName = "password"

	v.validator.
		ValidateStringField(input, fieldName).
		ValidateFieldMaxLength(input, fieldName, 50).
		ValidateFieldMinLength(input, fieldName, 8)
}

func (v *updateValidation) validateName(input string) {
	const fieldName = "name"

	v.validator.
		ValidateFieldMaxLength(input, fieldName, 255)
}

func (v *updateValidation) validateBirth(input string) {
	const fieldName = "birth"

	_, err := util_entity.GetDateFromString(input)
	if err != nil {
		v.validator.AddError(errors_validation.MustBeDateFormat(fieldName), fieldName)
	}
}

func (v *updateValidation) validatePasswordConfirmation(input, password string) {
	const fieldName = "passwordConfirmation"

	if input == "" && password == "" {
		return
	}

	v.validator.
		ValidateStringField(input, fieldName).
		ValidateFieldMaxLength(input, fieldName, 50).
		ValidatePasswordConfirmationEquals(input, password)
}
