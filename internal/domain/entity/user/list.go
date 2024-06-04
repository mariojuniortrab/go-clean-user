package user_entity

import (
	protocol_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/protocol"
	util_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/util"
)

type ListUserDto struct {
	protocol_entity.List
	ID                string
	Name              string
	WillFilterActives bool
	Active            bool
	Email             string
}
type UserListItemOutputDto struct {
	ID     string `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	Birth  string `json:"birth"`
	Active bool   `json:"active"`
}

type filter struct {
	ID     string
	Email  string
	Name   string
	Active string
}

type ListUserInputDto struct {
	protocol_entity.ListInputDto
	filter
}

type ListOutputDto struct {
	protocol_entity.ListOutputDto
	Items []*UserListItemOutputDto `json:"items"`
}

func NewListUserDto(input *ListUserInputDto) (*ListUserDto, error) {
	willFilterActives := false
	active := false

	if input.Active != "" {
		willFilterActives = true
	}

	if input.Active == "true" {
		active = true
	}

	listUserDto := &ListUserDto{
		Active:            active,
		WillFilterActives: willFilterActives,
		ID:                input.ID,
		Name:              input.Name,
		Email:             input.Email,
	}

	err := protocol_entity.FillListFromInput(&input.ListInputDto, &listUserDto.List)
	if err != nil {
		return nil, err
	}

	return listUserDto, nil
}

func NewUserListItemOutputDto(user *User) *UserListItemOutputDto {
	birth := util_entity.GetStringFromDate(user.Birth)

	return &UserListItemOutputDto{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Birth:  birth,
		Active: user.Active,
	}
}
