package infra_util

import "database/sql"

func MapOrderedFieldsFromRows(fieldsToGet []string, rows *sql.Rows) ([]map[string]string, error) {
	var result []map[string]string

	for rows.Next() {
		inputs := NewDumbArrayForScan(fieldsToGet)

		err := rows.Scan(inputs...)
		if err != nil {
			return nil, err
		}

		result = append(result, NewMapFromDumbArray(inputs, fieldsToGet))
	}

	return result, nil
}
