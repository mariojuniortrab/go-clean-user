package auth_entity

import (
	protocol_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/protocol"
	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
)

type LoginInputDto struct {
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

func (u LoginInputDto) IsEmpty() bool {
	return u == LoginInputDto{}
}

func (u *LoginInputDto) New() protocol_entity.Emptyable {
	return &LoginInputDto{}
}

type LoginDto struct {
	ID       string
	Name     string
	Active   bool
	Email    string
	Password string
}

type LoginOutputDto struct {
	Token  string `json:"token"`
	ID     string `json:"id"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
	Email  string `json:"email"`
}

func NewLoginOutputDto(input *LoginDto, token string) *LoginOutputDto {
	return &LoginOutputDto{
		Token:  token,
		ID:     input.ID,
		Name:   input.Name,
		Active: input.Active,
		Email:  input.Email,
	}
}

func NewLoginDto(user *user_entity.User) *LoginDto {
	return &LoginDto{
		ID:       user.ID,
		Name:     user.Name,
		Active:   user.Active,
		Email:    user.Email,
		Password: user.Password,
	}
}
