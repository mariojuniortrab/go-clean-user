package infra_util

import (
	"database/sql"
)

func MapOrderedFieldsFromRow(fieldsToGet []string, row *sql.Row) (map[string]string, error) {
	inputs := NewDumbArrayForScan(fieldsToGet)
	err := row.Scan(inputs...)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return NewMapFromDumbArray(inputs, fieldsToGet), nil
}
