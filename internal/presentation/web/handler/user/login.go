package user_handler

import (
	"encoding/json"
	"net/http"

	auth_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/auth"
	auth_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/auth"
	user_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/user"
	web_response_manager "github.com/mariojuniortrab/hauling-api/internal/presentation/web/response-manager"
)

type loginHandler struct {
	loginValidation user_validation.LoginValidation
	loginUseCase    *auth_usecase.Login
}

func NewLoginHandle(loginValidation user_validation.LoginValidation,
	loginUseCase *auth_usecase.Login) *loginHandler {
	return &loginHandler{
		loginValidation,
		loginUseCase,
	}
}

func (h *loginHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input auth_entity.LoginInputDto

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		web_response_manager.RespondInternalServerError(w, err)
		return
	}

	validationErrs := h.loginValidation.Validate(&input)
	if validationErrs != nil {
		web_response_manager.RespondFieldErrorValidation(w, validationErrs)
		return
	}

	user, err := h.loginUseCase.GetUserDataForLogin(&input)
	if err != nil {
		web_response_manager.RespondInternalServerError(w, err)
		return
	}
	if user == nil {
		web_response_manager.RespondLoginInvalid(w)
		return
	}

	if h.loginValidation.IsCredentialInvalid(user, *input.Password) {
		web_response_manager.RespondLoginInvalid(w)
		return
	}

	output, err := h.loginUseCase.Execute(user)
	if err != nil {
		web_response_manager.RespondInternalServerError(w, err)
		return
	}

	web_response_manager.RespondOk(w, "login successful", output)
}
