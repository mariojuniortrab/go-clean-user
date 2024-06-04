package user_handler

import (
	"encoding/json"
	"net/http"

	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	protocol_application "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/protocol/application"
	user_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/user"
	user_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/user"
	web_response_manager "github.com/mariojuniortrab/hauling-api/internal/presentation/web/response-manager"
)

type updateHandler struct {
	urlParser        protocol_application.URLParser
	updateUseCase    *user_usecase.UpdateUserUseCase
	updateValidation user_validation.UpdateValidation
}

func NewUpdateHandler(urlParser protocol_application.URLParser,
	updateUseCase *user_usecase.UpdateUserUseCase,
	updateValidation user_validation.UpdateValidation) *updateHandler {
	return &updateHandler{urlParser, updateUseCase, updateValidation}
}

func (h *updateHandler) Handle(w http.ResponseWriter, r *http.Request) {
	id := h.urlParser.GetPathParamFromURL(r, "id")

	var payload user_entity.UserUpdateInputDto

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		web_response_manager.RespondInternalServerError(w, err)
		return
	}

	validationErrs := h.updateValidation.Validate(&payload)
	if validationErrs != nil {
		web_response_manager.RespondFieldErrorValidation(w, validationErrs)
		return
	}

	editedUser, err := h.updateUseCase.GetEditedUser(id, &payload)
	loggedUserId := r.Context().Value(protocol_application.UserIdKey).(string)
	editedUser.FillUpdatableFieldsForUpdate(loggedUserId)

	if err != nil {
		web_response_manager.RespondInternalServerError(w, err)
		return
	}
	if editedUser == nil {
		web_response_manager.RespondNotFound(w, "user")
		return
	}

	result, err := h.updateUseCase.Execute(id, editedUser)
	if err != nil {
		web_response_manager.RespondInternalServerError(w, err)
		return
	}

	web_response_manager.RespondOk(w, "success", result)
}
