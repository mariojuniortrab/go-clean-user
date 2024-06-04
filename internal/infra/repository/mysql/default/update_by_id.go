package default_mysql_repository

import (
	"fmt"
	"strings"

	protocols_mysql_repository "github.com/mariojuniortrab/hauling-api/internal/infra/repository/mysql/protocols"
)

func UpdateById(r protocols_mysql_repository.DefaultRepositoryProtocol, mappedValues map[string]interface{}, id string) error {
	setList := []string{}
	valuesList := []interface{}{}

	for index, value := range mappedValues {
		setList = append(setList, fmt.Sprintf(" %s = ?", index))
		valuesList = append(valuesList, value)
	}

	if len(setList) == 0 || len(valuesList) == 0 {
		panic("error trying to build update query")
	}

	valuesList = append(valuesList, id)

	query := strings.Replace(rawUpdateQuery, "##table##", r.GetTableName(), 1)
	query = strings.Replace(query, "##set##", strings.Join(setList, ","), 1)

	_, err := r.GetDb().Exec(query, valuesList...)

	if err != nil {
		return err
	}

	return nil
}
