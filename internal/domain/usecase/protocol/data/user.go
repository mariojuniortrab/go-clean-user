package protocol_data

import (
	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
)

type SignupRepository interface {
	Create(*user_entity.User) error
}

type CreateUserRepository interface {
	Create(*user_entity.User) error
}

type ListUserRepository interface {
	List(input *user_entity.ListUserDto) ([]*user_entity.User, int, error)
}

type GetUserByIdRepository interface {
	GetById(id string) (*user_entity.User, error)
}

type RemoveUserRepository interface {
	Remove(id string) error
}

type UpdateRepository interface {
	GetForUpdate(id string) (*user_entity.User, error)
	Update(id string, editedUser *user_entity.User) (*user_entity.User, error)
}
type GetUserByEmailRepository interface {
	GetByEmail(email string, id string) (*user_entity.User, error)
}
