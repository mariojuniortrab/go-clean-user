package user_entity

import (
	protocol_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/protocol"
	util_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/util"
)

type UserUpdateInputDto struct {
	Name                 *string `json:"name"`
	Birth                *string `json:"birth"`
	Active               *bool   `json:"active,omitempty"`
	Password             *string `json:"password"`
	PasswordConfirmation *string `json:"passwordConfirmation"`
}

func (u UserUpdateInputDto) IsEmpty() bool {
	return u == UserUpdateInputDto{}
}

func (u *UserUpdateInputDto) New() protocol_entity.Emptyable {
	return &UserUpdateInputDto{}
}

func NewUserUpdateInputDto(user *User) *UserUpdateInputDto {
	birth := util_entity.GetStringFromDate(user.Birth)

	return &UserUpdateInputDto{
		Name:     &user.Name,
		Birth:    &birth,
		Active:   &user.Active,
		Password: &user.Password,
	}
}

func NewUserFromUpdateInputDto(input *UserUpdateInputDto) (*User, error) {
	formattedDate, err := util_entity.GetDateFromString(*input.Birth)
	if err != nil {
		return nil, err
	}

	return &User{
		Name:     *input.Name,
		Password: *input.Password,
		Active:   *input.Active,
		Birth:    formattedDate,
	}, nil
}

func GetEditedUser(user *User, payload *UserUpdateInputDto) (*User, error) {
	if payload.Birth != nil {
		formattedDate, err := util_entity.GetDateFromString(*payload.Birth)
		if err != nil {
			return nil, err
		}

		user.Birth = formattedDate
	}

	if payload.Active != nil {
		user.Active = *payload.Active
	}

	if payload.Password != nil {
		user.Password = *payload.Password
	}

	if payload.Name != nil {
		user.Name = *payload.Name
	}

	return user, nil
}
