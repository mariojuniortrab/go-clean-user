package user_entity

import (
	protocol_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/protocol"
	util_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/util"
)

type CreateUserInputDto struct {
	Password             *string `json:"password"`
	Name                 *string `json:"name"`
	PasswordConfirmation *string `json:"passwordConfirmation"`
	Email                *string `json:"email"`
	Birth                *string `json:"birth"`
}

func (u CreateUserInputDto) IsEmpty() bool {
	return u == CreateUserInputDto{}
}

func (u *CreateUserInputDto) New() protocol_entity.Emptyable {
	return &CreateUserInputDto{}
}

type CreateUserOutputDto struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Birth string `json:"birth"`
}

func NewCreateUserOutputDto(user *User) *CreateUserOutputDto {
	return &CreateUserOutputDto{
		Email: user.Email,
		Name:  user.Name,
		ID:    user.ID,
		Birth: util_entity.GetStringFromDate(user.Birth),
	}
}
