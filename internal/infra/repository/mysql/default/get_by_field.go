package default_mysql_repository

import (
	"fmt"
	"strings"

	protocols_mysql_repository "github.com/mariojuniortrab/hauling-api/internal/infra/repository/mysql/protocols"
	infra_util "github.com/mariojuniortrab/hauling-api/internal/infra/util"
)

func GetByField(
	r protocols_mysql_repository.DefaultSelectRepositoryProtocol,
	fieldsToGet []string,
	mappedWhere map[string]interface{}) (map[string]string, error) {

	query := strings.Replace(rawSelectQuery, "##table##", r.GetTableName(), 1)
	query = strings.Replace(query, "##fields##", strings.Join(fieldsToGet, ","), 1)
	where, values := mountGetByFieldWhere(mappedWhere)

	if where == "" || len(values) == 0 {
		panic("cond invalid")
	}

	query += where

	row := r.GetDb().QueryRow(query, values...)

	return infra_util.MapOrderedFieldsFromRow(fieldsToGet, row)
}

func mountGetByFieldWhere(cond map[string]interface{}) (string, []interface{}) {
	result := ""
	where := []string{}
	var values []interface{}

	for index, value := range cond {
		if strings.ToLower(index) != "id" {
			where = append(where, fmt.Sprintf(" %s = ? ", index))
		} else {
			where = append(where, fmt.Sprintf(" %s <> ? ", index))
		}
		values = append(values, value)
	}

	if len(where) > 0 {
		result = fmt.Sprintf(" WHERE %s ", strings.Join(where, " AND "))
	}

	return result, values
}
