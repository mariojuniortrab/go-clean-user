package user_mysql_repository

import (
	"database/sql"

	default_mysql_repository "github.com/mariojuniortrab/hauling-api/internal/infra/repository/mysql/default"
)

type removeUserRepository struct {
	UserRepositoryMysql
}

func NewRemoveUserRepository(db *sql.DB) *removeUserRepository {
	repository := &removeUserRepository{}
	repository.DB = db
	return repository
}

func (r *removeUserRepository) Remove(id string) error {
	return default_mysql_repository.RemoveById(r, id)
}
