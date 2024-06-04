package user_mysql_repository

import (
	"database/sql"

	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	default_mysql_repository "github.com/mariojuniortrab/hauling-api/internal/infra/repository/mysql/default"
)

type listUserRepository struct {
	UserRepositoryMysql
}

func NewListUserRepository(db *sql.DB) *listUserRepository {
	repository := &listUserRepository{}
	repository.DB = db
	return repository
}

func (r *listUserRepository) List(input *user_entity.ListUserDto) ([]*user_entity.User, int, error) {
	var result []*user_entity.User

	mappedWhere := r.getWhereMap(input)
	fieldsToGet := []string{"id", "name", "email", "birth", "active"}

	mappedResult, total, err := default_mysql_repository.List(&input.List, r, fieldsToGet, mappedWhere)
	if err != nil {
		return nil, 0, err
	}

	for _, v := range mappedResult {
		user, err := user_entity.NewUserFromMap(v)
		if err != nil {
			return nil, 0, err
		}

		result = append(result, user)
	}

	return result, total, nil
}

func (r *listUserRepository) getWhereMap(input *user_entity.ListUserDto) map[string]interface{} {
	whereMap := make(map[string]interface{})

	if input.ID != "" {
		whereMap["ID"] = input.ID
	}

	if input.Email != "" {
		whereMap["email"] = input.Email
	}

	if input.Name != "" {
		whereMap["name"] = input.Name
	}

	if input.WillFilterActives {
		whereMap["active"] = input.Active
	}

	return whereMap
}
