package auth_usecase

import (
	auth_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/auth"
	protocol_application "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol/application"
	protocol_data "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol/data"
)

const timeToExpire = 4

type Login struct {
	getUserByEmailRepository protocol_data.GetUserByEmailRepository
	tokenizer                protocol_application.Tokenizer
}

func NewLoginUseCase(getUserByEmailRepository protocol_data.GetUserByEmailRepository,
	tokenizer protocol_application.Tokenizer) *Login {
	return &Login{getUserByEmailRepository, tokenizer}
}

func (u *Login) Execute(input *auth_entity.LoginDto) (*auth_entity.LoginOutputDto, error) {
	token, err := u.tokenizer.GenerateToken(input.ID, input.Email, timeToExpire)
	if err != nil {
		return nil, err
	}

	return auth_entity.NewLoginOutputDto(input, token), nil
}

func (u *Login) GetUserDataForLogin(input *auth_entity.LoginInputDto) (*auth_entity.LoginDto, error) {
	user, err := u.getUserByEmailRepository.GetByEmail(*input.Email, "")
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	return auth_entity.NewLoginDto(user), nil
}
