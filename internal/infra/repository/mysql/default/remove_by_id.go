package default_mysql_repository

import (
	"strings"

	protocols_mysql_repository "github.com/mariojuniortrab/hauling-api/internal/infra/repository/mysql/protocols"
)

func RemoveById(r protocols_mysql_repository.DefaultRepositoryProtocol, id string) error {

	query := strings.Replace(rawDeleteQuery, "##table##", r.GetTableName(), 1)

	_, err := r.GetDb().Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
