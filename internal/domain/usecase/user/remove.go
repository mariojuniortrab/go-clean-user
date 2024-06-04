package user_usecase

import (
	"errors"

	protocol_data "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol/data"
)

type RemoveUserUseCase struct {
	getByIdRepostory protocol_data.GetUserByIdRepository
	removeRepository protocol_data.RemoveUserRepository
}

func NewRemoveUserUsecase(getByIdRepostory protocol_data.GetUserByIdRepository,
	removeRepository protocol_data.RemoveUserRepository) *RemoveUserUseCase {
	return &RemoveUserUseCase{getByIdRepostory, removeRepository}
}

func (u *RemoveUserUseCase) Execute(id string) (error, error) {
	user, err := u.getByIdRepostory.GetById(id)

	if err != nil {
		return err, nil
	}

	if user == nil {
		return nil, errors.New("not found")
	}

	err = u.removeRepository.Remove(id)
	if err != nil {
		return err, nil
	}

	return nil, nil
}
