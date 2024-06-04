package user_usecase

import (
	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	util_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/util"
	protocol_application "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol/application"
	protocol_data "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol/data"
)

type Create struct {
	repository protocol_data.CreateUserRepository
	encrypter  protocol_application.Encrypter
}

func NewCreateUserUseCase(repository protocol_data.CreateUserRepository,
	encrypter protocol_application.Encrypter) *Create {
	return &Create{repository, encrypter}
}

func (u *Create) Execute(input *user_entity.CreateUserInputDto, loggedUserId string) (*user_entity.CreateUserOutputDto, error) {

	formattedDate, err := util_entity.GetDateFromString(*input.Birth)
	if err != nil {
		return nil, err
	}

	hashPassword, err := u.encrypter.Hash(*input.Password)
	if err != nil {
		return nil, err
	}

	user := user_entity.NewUser(*input.Name, hashPassword, *input.Email, formattedDate)
	user.FillUpdatableFieldsForCreate(loggedUserId)

	err = u.repository.Create(user)
	if err != nil {
		return nil, err
	}

	return user_entity.NewCreateUserOutputDto(user), nil
}
