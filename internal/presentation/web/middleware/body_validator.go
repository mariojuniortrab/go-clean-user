package web_middleware

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	protocol_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/protocol"
	web_response_manager "github.com/mariojuniortrab/hauling-api/internal/presentation/web/response-manager"
)

type bodyValidator struct {
	fieldsToValidate protocol_entity.Emptyable
}

func NewBodyValidator(fieldsToValidate protocol_entity.Emptyable) *bodyValidator {
	return &bodyValidator{
		fieldsToValidate: fieldsToValidate,
	}
}

func (m *bodyValidator) Middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		m.fieldsToValidate = m.fieldsToValidate.New()

		fakebody, err := getFakeBody(r)

		if err != nil {
			web_response_manager.RespondInternalServerError(w, err)
			return
		}

		fakebody = http.MaxBytesReader(w, fakebody, 1048576)

		dec := json.NewDecoder(fakebody)
		dec.DisallowUnknownFields()

		err = dec.Decode(m.fieldsToValidate)

		if err != nil {
			var syntaxError *json.SyntaxError
			var unmarshalTypeError *json.UnmarshalTypeError

			switch {

			case errors.As(err, &syntaxError):
				msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
				web_response_manager.RespondGenericError(w, http.StatusBadRequest, msg)
				return

			case errors.Is(err, io.ErrUnexpectedEOF):
				msg := "Request body contains badly-formed JSON"
				web_response_manager.RespondGenericError(w, http.StatusBadRequest, msg)
				return

			case errors.As(err, &unmarshalTypeError):
				msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
				web_response_manager.RespondGenericError(w, http.StatusBadRequest, msg)
				return

			case strings.HasPrefix(err.Error(), "json: unknown field "):
				fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
				msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
				web_response_manager.RespondGenericError(w, http.StatusBadRequest, msg)
				return

			case errors.Is(err, io.EOF):
				msg := "Request body must not be empty"
				web_response_manager.RespondGenericError(w, http.StatusBadRequest, msg)
				return

			case err.Error() == "http: request body too large":
				msg := "Request body must not be larger than 1MB"
				web_response_manager.RespondGenericError(w, http.StatusRequestEntityTooLarge, msg)
				return

			default:
				web_response_manager.RespondInternalServerError(w, err)
				return
			}
		}

		err = dec.Decode(&struct{}{})
		if !errors.Is(err, io.EOF) {
			msg := "Request body must only contain a single JSON object"
			web_response_manager.RespondGenericError(w, http.StatusBadRequest, msg)
			return
		}

		if m.fieldsToValidate.IsEmpty() {
			msg := "Request body must not be empty"
			web_response_manager.RespondGenericError(w, http.StatusBadRequest, msg)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func getFakeBody(r *http.Request) (io.ReadCloser, error) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	copyBody := io.NopCloser(bytes.NewBuffer(bodyBytes))

	return copyBody, nil
}
