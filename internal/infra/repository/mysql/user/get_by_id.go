package user_mysql_repository

import (
	"database/sql"

	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	default_mysql_repository "github.com/mariojuniortrab/hauling-api/internal/infra/repository/mysql/default"
)

type getUserByIdRepository struct {
	UserRepositoryMysql
}

func NewGetUserByIdRepository(db *sql.DB) *getUserByIdRepository {
	repository := &getUserByIdRepository{}
	repository.DB = db
	return repository
}

func (r *getUserByIdRepository) GetById(id string) (*user_entity.User, error) {
	fieldsToGet := []string{"id", "name", "email", "active", "birth", "created_at",
		"updated_at", "created_id", "updated_id"}

	mappedResult, err := default_mysql_repository.GetById(r, fieldsToGet, id)
	if err != nil {
		return nil, err
	}
	if mappedResult == nil {
		return nil, nil
	}

	user, err := user_entity.NewUserFromMap(mappedResult)
	if err != nil {
		return nil, err
	}

	return user, nil
}
