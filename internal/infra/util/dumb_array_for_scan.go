package infra_util

func NewDumbArrayForScan(fieldsToGet []string) []interface{} {
	result := make([]interface{}, len(fieldsToGet))

	for i := range fieldsToGet {
		var field string
		result[i] = &field
	}

	return result
}

func NewMapFromDumbArray(inputs []interface{}, fieldsToGet []string) map[string]string {
	var result = map[string]string{}

	for i, v := range fieldsToGet {
		result[v] = *(inputs[i].(*string))
	}

	return result
}
