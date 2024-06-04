package default_mysql_repository

import (
	"strings"

	protocols_mysql_repository "github.com/mariojuniortrab/hauling-api/internal/infra/repository/mysql/protocols"
)

func Insert(r protocols_mysql_repository.DefaultRepositoryProtocol, mappedValues map[string]interface{}) error {
	fieldList := []string{}
	markList := []string{}
	valuesList := []interface{}{}

	for index, value := range mappedValues {
		fieldList = append(fieldList, index)
		markList = append(markList, "?")
		valuesList = append(valuesList, value)
	}

	if len(fieldList) == 0 || len(markList) == 0 || len(valuesList) == 0 {
		panic("error trying to build insert query")
	}

	query := strings.Replace(rawInsertQuery, "##table##", r.GetTableName(), 1)
	query = strings.Replace(query, "##fields##", strings.Join(fieldList, ","), 1)
	query = strings.Replace(query, "##values##", strings.Join(markList, ","), 1)

	_, err := r.GetDb().Exec(query, valuesList...)

	if err != nil {
		return err
	}

	return nil
}
