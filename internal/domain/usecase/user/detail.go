package user_usecase

import (
	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	protocol_data "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol/data"
)

type DetailUserUseCase struct {
	repository protocol_data.GetUserByIdRepository
}

func NewDetailUserUsecase(repository protocol_data.GetUserByIdRepository) *DetailUserUseCase {
	return &DetailUserUseCase{repository}
}

func (u *DetailUserUseCase) Execute(id string) (*user_entity.UserDetailOutputDto, error) {
	user, err := u.repository.GetById(id)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	output := user_entity.NewUserDetailOutputDto(user)
	return output, nil
}
