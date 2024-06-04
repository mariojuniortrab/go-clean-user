package default_mysql_repository

import (
	"database/sql"
	"fmt"
	"strings"

	protocol_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/protocol"
	protocols_mysql_repository "github.com/mariojuniortrab/hauling-api/internal/infra/repository/mysql/protocols"
	infra_util "github.com/mariojuniortrab/hauling-api/internal/infra/util"
)

func List(list *protocol_entity.List,
	r protocols_mysql_repository.DefaultSelectRepositoryProtocol,
	fieldsToGet []string,
	mappedWhere map[string]interface{}) ([]map[string]string, int, error) {

	offset := (list.Page - 1) * list.Limit

	query := strings.Replace(rawSelectQuery, "##table##", r.GetTableName(), 1)

	query += mountListWhere(mappedWhere, list)
	total, err := getTotalRows(r, query)

	query = strings.Replace(query, "##fields##", strings.Join(fieldsToGet, ","), 1)

	if err != nil {
		return nil, 0, err
	}

	query += getOrderByForList(list.OrderBy, list.OrderType)
	query += fmt.Sprintf(" LIMIT %d ", list.Limit)
	query += fmt.Sprintf(" OFFSET %d ", offset)

	rows, err := r.GetDb().Query(query)
	if err != nil {
		return nil, 0, err
	}

	defer rows.Close()

	mappedList, err := infra_util.MapOrderedFieldsFromRows(fieldsToGet, rows)
	if err != nil {
		return nil, 0, err
	}

	return mappedList, total, err
}

func getOrderByForList(orderBy, orderType string) string {
	result := ""

	if orderType != "" {
		orderType = " " + orderType + " "
	}

	if orderBy != "" {
		result = fmt.Sprintf(" ORDER BY %s %s ", orderBy, orderType)
	}

	return result
}

func mountListWhere(cond map[string]interface{}, list *protocol_entity.List) string {
	result := ""
	where := []string{}

	for index, value := range cond {
		switch value.(type) {
		case string:
			where = append(where, fmt.Sprintf(" %s LIKE '%%%s%%' ", index, value))
		case bool:
			where = append(where, fmt.Sprintf(" %s = %t ", index, value))
		case int, float32, float64:
			where = append(where, fmt.Sprintf(" %s = %d ", index, value))
		default:
			panic("value type is unknown")
		}
	}

	if list.Q != "" {
		where = append(where, mountQWhere(cond, list.Q))
	}

	if len(where) > 0 {
		result = fmt.Sprintf(" WHERE %s ", strings.Join(where, " AND "))
	}

	return result
}

func mountQWhere(cond map[string]any, q string) string {
	result := ""
	qWhere := []string{}

	for index, value := range cond {
		switch value.(type) {
		case string:
			qWhere = append(qWhere, fmt.Sprintf(" %s LIKE '%%%s%%' ", index, q))
		case bool:
			qWhere = append(qWhere, fmt.Sprintf(" %s = '%s' ", index, q))
		case int, float32, float64:
			qWhere = append(qWhere, fmt.Sprintf(" %s = '%s' ", index, q))
		default:
			panic("value type is unknown")
		}
	}

	if len(qWhere) > 0 {
		result = fmt.Sprintf(" ( %s ) ", strings.Join(qWhere, " OR "))
	}

	return result
}

func getTotalRows(r protocols_mysql_repository.DefaultSelectRepositoryProtocol, rawQuery string) (int, error) {
	query := strings.Replace(rawQuery, "##fields##", "count(ID)", 1)
	var total int

	err := r.GetDb().QueryRow(query).
		Scan(&total)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}

		return 0, err
	}

	return total, nil
}
