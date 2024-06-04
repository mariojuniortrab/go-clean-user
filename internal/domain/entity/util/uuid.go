package util_entity

import "github.com/google/uuid"

func IsUIID(stringToValidate string) bool {
	_, err := uuid.Parse(stringToValidate)
	return err == nil
}
