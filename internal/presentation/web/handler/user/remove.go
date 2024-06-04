package user_handler

import (
	"net/http"

	protocol_application "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol/application"
	user_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/user"
	web_response_manager "github.com/mariojuniortrab/hauling-api/internal/presentation/web/response-manager"
)

type removeHandler struct {
	urlParser     protocol_application.URLParser
	removeUseCase *user_usecase.RemoveUserUseCase
}

func NewRemoveHandler(urlParser protocol_application.URLParser,
	removeUseCase *user_usecase.RemoveUserUseCase) *removeHandler {
	return &removeHandler{urlParser, removeUseCase}
}

func (h *removeHandler) Handle(w http.ResponseWriter, r *http.Request) {
	id := h.urlParser.GetPathParamFromURL(r, "id")

	if id == "" {
		web_response_manager.RespondUiidIsRequired(w)
	}

	err, errNotFound := h.removeUseCase.Execute(id)
	if err != nil {
		web_response_manager.RespondInternalServerError(w, err)
		return
	}

	if errNotFound != nil {
		web_response_manager.RespondNotFound(w, "user")
		return
	}

	web_response_manager.RespondOk(w, "removed", nil)
}
