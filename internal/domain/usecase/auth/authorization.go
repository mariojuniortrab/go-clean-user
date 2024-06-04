package auth_usecase

import (
	"errors"

	auth_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/auth"
	protocol_application "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol/application"
)

type AuthInputDto struct {
	Token string `json:"token"`
}

type Authorization struct {
	tokenizer protocol_application.Tokenizer
}

func NewAuthorization(tokenizer protocol_application.Tokenizer) *Authorization {
	return &Authorization{
		tokenizer,
	}
}

func (u *Authorization) Execute(input *AuthInputDto) (*auth_entity.TokenOutputDto, error) {
	if input.Token == "" {
		return nil, errors.New("token is empty")
	}

	output, err := u.tokenizer.ParseToken(input.Token)
	if err != nil {
		return nil, err
	}

	return output, nil
}
