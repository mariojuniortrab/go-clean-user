package user_mysql_repository

import (
	"database/sql"

	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	default_mysql_repository "github.com/mariojuniortrab/hauling-api/internal/infra/repository/mysql/default"
)

type updateUserRepository struct {
	UserRepositoryMysql
}

func NewUpdateUserRepository(db *sql.DB) *updateUserRepository {
	repository := &updateUserRepository{}
	repository.DB = db
	return repository
}

func (r *updateUserRepository) GetForUpdate(id string) (*user_entity.User, error) {
	fieldsToGet := []string{"id", "name", "password", "active", "birth", "email", "created_at",
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

func (r *updateUserRepository) Update(id string, editedUser *user_entity.User) (*user_entity.User, error) {
	err := default_mysql_repository.UpdateById(r, editedUser.Map(false), id)
	if err != nil {
		return nil, err
	}

	return r.getUpdatedUser(id)
}

func (r *updateUserRepository) getUpdatedUser(id string) (*user_entity.User, error) {
	fieldsToGet := []string{"id", "name", "email", "active", "birth", "created_at",
		"updated_at", "created_id", "updated_id"}

	mappedResult, err := default_mysql_repository.GetById(r, fieldsToGet, id)
	if err != nil {
		return nil, err
	}

	user, err := user_entity.NewUserFromMap(mappedResult)
	if err != nil {
		return nil, err
	}

	return user, nil
}
