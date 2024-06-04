package web_middleware

import (
	"net/http"

	util_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/util"
	protocol_application "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol/application"
	web_response_manager "github.com/mariojuniortrab/hauling-api/internal/presentation/web/response-manager"
)

type uuidParser struct {
	urlParser protocol_application.URLParser
}

func NewUuidParser(urlParser protocol_application.URLParser) *uuidParser {
	return &uuidParser{urlParser}
}

func (m *uuidParser) Middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		uuid := m.urlParser.GetPathParamFromURL(r, "id")
		if !util_entity.IsUIID(uuid) {
			web_response_manager.RespondUiidInvalid(w)
			return
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
