package user_mysql_repository

import (
	"database/sql"

	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	default_mysql_repository "github.com/mariojuniortrab/hauling-api/internal/infra/repository/mysql/default"
)

type getUserByEmailRepository struct {
	UserRepositoryMysql
}

func NewGetUserByEmailRepository(db *sql.DB) *getUserByEmailRepository {
	repository := &getUserByEmailRepository{}
	repository.DB = db
	return repository
}

func (r *getUserByEmailRepository) GetByEmail(email, id string) (*user_entity.User, error) {
	whereMap := getWhereMap(email, id)
	fieldsToGet := []string{"id", "email", "name", "password", "active"}

	mappedResult, err := default_mysql_repository.GetByField(r, fieldsToGet, whereMap)
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

func getWhereMap(email, id string) map[string]interface{} {
	whereMap := map[string]interface{}{
		"email": email,
	}

	if id != "" {
		whereMap["id"] = id
	}

	return whereMap
}
