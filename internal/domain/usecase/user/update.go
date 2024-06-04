package user_usecase

import (
	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	protocol_application "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol/application"
	protocol_data "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol/data"
)

type UpdateUserUseCase struct {
	repository protocol_data.UpdateRepository
	encrypter  protocol_application.Encrypter
}

func NewUpdateUserUsecase(repository protocol_data.UpdateRepository,
	encrypter protocol_application.Encrypter) *UpdateUserUseCase {
	return &UpdateUserUseCase{repository, encrypter}
}

func (u *UpdateUserUseCase) Execute(id string, user *user_entity.User) (*user_entity.UserDetailOutputDto, error) {

	user, err := u.repository.Update(id, user)
	if err != nil {
		return nil, err
	}

	output := user_entity.NewUserDetailOutputDto(user)
	return output, nil
}

func (u *UpdateUserUseCase) GetEditedUser(id string, payload *user_entity.UserUpdateInputDto) (*user_entity.User, error) {
	user, err := u.repository.GetForUpdate(id)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	if payload.Password != nil {
		hash, err := u.encrypter.Hash(*payload.Password)
		if err != nil {
			return nil, err
		}

		payload.Password = &hash
	}

	return user_entity.GetEditedUser(user, payload)
}
