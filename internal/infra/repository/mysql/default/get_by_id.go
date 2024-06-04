package default_mysql_repository

import (
	"strings"

	protocols_mysql_repository "github.com/mariojuniortrab/hauling-api/internal/infra/repository/mysql/protocols"
	infra_util "github.com/mariojuniortrab/hauling-api/internal/infra/util"
)

func GetById(r protocols_mysql_repository.DefaultSelectRepositoryProtocol,
	fieldsToGet []string,
	id string) (map[string]string, error) {

	query := strings.Replace(rawSelectQuery, "##table##", r.GetTableName(), 1)
	query = strings.Replace(query, "##fields##", strings.Join(fieldsToGet, ","), 1)

	query += "WHERE id = ?"

	row := r.GetDb().QueryRow(query, id)

	return infra_util.MapOrderedFieldsFromRow(fieldsToGet, row)
}
