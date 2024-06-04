package user_handler

import (
	"net/http"

	protocol_application "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol/application"
	user_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/user"
	web_response_manager "github.com/mariojuniortrab/hauling-api/internal/presentation/web/response-manager"
)

type detailHandler struct {
	urlParser     protocol_application.URLParser
	detailUseCase *user_usecase.DetailUserUseCase
}

func NewDetailHandler(urlParser protocol_application.URLParser,
	detailUseCase *user_usecase.DetailUserUseCase) *detailHandler {
	return &detailHandler{urlParser, detailUseCase}
}

func (h *detailHandler) Handle(w http.ResponseWriter, r *http.Request) {
	id := h.urlParser.GetPathParamFromURL(r, "id")

	if id == "" {
		web_response_manager.RespondUiidIsRequired(w)
	}

	result, err := h.detailUseCase.Execute(id)
	if err != nil {
		web_response_manager.RespondInternalServerError(w, err)
		return
	}

	if result == nil {
		web_response_manager.RespondNotFound(w, "user")
		return
	}

	web_response_manager.RespondOk(w, "success", result)
}
