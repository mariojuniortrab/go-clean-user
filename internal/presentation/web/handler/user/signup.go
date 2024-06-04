package user_handler

import (
	"encoding/json"
	"net/http"

	auth_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/auth"
	auth_usecase "github.com/mariojuniortrab/hauling-api/internal/domain/usecase/auth"
	user_validation "github.com/mariojuniortrab/hauling-api/internal/domain/validation/user"
	web_response_manager "github.com/mariojuniortrab/hauling-api/internal/presentation/web/response-manager"
)

type signupHandler struct {
	signUpValidation user_validation.SignupValidation
	signUp           *auth_usecase.Signup
}

func NewSignupHandler(signUpValidation user_validation.SignupValidation,
	signUp *auth_usecase.Signup) *signupHandler {
	return &signupHandler{
		signUpValidation: signUpValidation,
		signUp:           signUp,
	}
}

func (h *signupHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input auth_entity.SignupInputDto

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		web_response_manager.RespondInternalServerError(w, err)
		return
	}

	validationErrs := h.signUpValidation.Validate(&input)
	if validationErrs != nil {
		web_response_manager.RespondFieldErrorValidation(w, validationErrs)
		return
	}

	alreadyExistsErr, err := h.signUpValidation.AlreadyExists(*input.Email, "")
	if err != nil {
		web_response_manager.RespondInternalServerError(w, err)
		return
	}

	if alreadyExistsErr != nil {
		web_response_manager.RespondConflictError(w, alreadyExistsErr)
		return
	}

	output, err := h.signUp.Execute(&input)
	if err != nil {
		web_response_manager.RespondInternalServerError(w, err)
		return
	}

	web_response_manager.RespondCreated(w, "created", output)
}
