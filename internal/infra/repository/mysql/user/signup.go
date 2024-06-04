package user_mysql_repository

import (
	"database/sql"

	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	default_mysql_repository "github.com/mariojuniortrab/hauling-api/internal/infra/repository/mysql/default"
)

type signupRepository struct {
	UserRepositoryMysql
}

func NewSignupRepository(db *sql.DB) *signupRepository {
	repository := &signupRepository{}
	repository.DB = db
	return repository
}

func (r *signupRepository) Create(user *user_entity.User) error {
	return default_mysql_repository.Insert(r, user.Map(true))
}
