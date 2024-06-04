package errors_validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_If_IsRequired_Returns_Correct_Error_Message(t *testing.T) {
	field := "any_field"
	message := "any_field is required"
	error := IsRequired(field)

	assert.Equal(t, error.Error(), message)
}

func Test_If_AlreadyExists_return_Correct_Error_Message(t *testing.T) {
	entity := "any_entity"
	error := AlreadyExists(entity)
	message := "any_entity already exists"

	assert.Equal(t, error.Error(), message)
}
