package protocol_entity

import (
	"strconv"
	"strings"
)

type List struct {
	Page      int
	Limit     int
	OrderBy   string
	OrderType string
	Q         string
}

type ListInputDto struct {
	Limit     string
	Page      string
	OrderBy   string
	OrderType string
	Q         string
}

type ListOutputDto struct {
	Total int `json:"total"`
}

func FillListFromInput(input *ListInputDto, list *List) error {
	page, err := strconv.Atoi(input.Page)
	if err != nil {
		return err
	}

	limit, err := strconv.Atoi(input.Limit)
	if err != nil {
		return err
	}

	list.Page = page
	list.Limit = limit
	list.OrderBy = strings.ToLower(input.OrderBy)
	list.OrderType = strings.ToLower(input.OrderType)
	list.Q = input.Q

	return nil
}
