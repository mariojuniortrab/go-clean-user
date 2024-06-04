package user_mysql_repository

import (
	"database/sql"

	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	default_mysql_repository "github.com/mariojuniortrab/hauling-api/internal/infra/repository/mysql/default"
)

type createUserRepository struct {
	UserRepositoryMysql
}

func NewCreateUserRepository(db *sql.DB) *createUserRepository {
	repository := &createUserRepository{}
	repository.DB = db
	return repository
}

func (r *createUserRepository) Create(user *user_entity.User) error {
	return default_mysql_repository.Insert(r, user.Map(true))
}
