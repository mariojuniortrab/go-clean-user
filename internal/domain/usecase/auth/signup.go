package auth_usecase

import (
	auth_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/auth"
	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	util_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/util"
	protocol_application "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol/application"
	protocol_data "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol/data"
)

type Signup struct {
	repository protocol_data.SignupRepository
	encrypter  protocol_application.Encrypter
}

func NewSignupUseCase(repository protocol_data.SignupRepository,
	encrypter protocol_application.Encrypter) *Signup {
	return &Signup{repository, encrypter}
}

func (u *Signup) Execute(input *auth_entity.SignupInputDto) (*auth_entity.SignupOutputDto, error) {

	formattedDate, err := util_entity.GetDateFromString(*input.Birth)
	if err != nil {
		return nil, err
	}

	hashPassword, err := u.encrypter.Hash(*input.Password)
	if err != nil {
		return nil, err
	}

	user := user_entity.NewUser(*input.Name, hashPassword, *input.Email, formattedDate)
	user.FillUpdatableFieldsForCreate(user.ID)

	err = u.repository.Create(user)
	if err != nil {
		return nil, err
	}

	return auth_entity.NewSignupOutputDto(user), nil
}
