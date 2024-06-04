package web_middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	protocol_application "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol/application"
	web_response_manager "github.com/mariojuniortrab/hauling-api/internal/presentation/web/response-manager"
)

type paginate struct {
	validator protocol_application.Validator
	urlParser protocol_application.URLParser
}

func NewPaginateMiddleware(validator protocol_application.Validator,
	urlParser protocol_application.URLParser) *paginate {
	return &paginate{validator, urlParser}
}

func (m *paginate) Middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		page := m.urlParser.GetQueryParamFromURL(r, "page")
		limit := m.urlParser.GetQueryParamFromURL(r, "limit")
		orderBy := m.urlParser.GetQueryParamFromURL(r, "orderBy")
		orderType := m.urlParser.GetQueryParamFromURL(r, "orderType")

		m.validate(page, "page")
		m.validate(limit, "limit")
		m.validateOrderFields(orderBy, orderType)

		if m.validator.HasErrors() {
			web_response_manager.RespondFieldErrorValidation(w, m.validator.GetErrorsAndClean())
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func (p *paginate) validate(input, fieldName string) {

	p.validator.
		ValidateRequiredField(input, fieldName).
		ValidateNumberField(input, fieldName)

	convertedField, err := strconv.Atoi(input)
	if err != nil {
		return
	}

	if convertedField == 0 {
		p.validator.AddError(fmt.Errorf("%s must be higher than 0", fieldName), fieldName)
	}
}

func (p *paginate) validateOrderFields(orderBy, orderType string) {
	if orderBy == "" && orderType == "" {
		return
	}

	if orderBy != "" && orderType == "" {
		p.validator.AddError(errors.New("orderType is required when orderBy is informed"), "orderType")
		return
	}

	if orderBy == "" && orderType != "" {
		p.validator.AddError(errors.New("orderBy is required when orderType is informed"), "orderBy")
		return
	}

	orderByItems := strings.Split(orderBy, ",")
	if len(orderByItems) > 1 {
		p.validator.AddError(errors.New("you can only order by 1 field at time"), "orderBy")
	}

	if strings.ToLower(orderType) != "asc" && strings.ToLower(orderType) != "desc" {
		p.validator.AddError(errors.New("orderType must be 'ASC' or 'DESC'"), "orderType")
	}
}
