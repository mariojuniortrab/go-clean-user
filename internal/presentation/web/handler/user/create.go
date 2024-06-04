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

type createUserandler struct {
	createUserValidation user_validation.CreateUserValidation
	createUser           *user_usecase.Create
}

func NewCreateUserHandler(createUserValidation user_validation.CreateUserValidation,
	createUser *user_usecase.Create) *createUserandler {
	return &createUserandler{
		createUserValidation: createUserValidation,
		createUser:           createUser,
	}
}

func (h *createUserandler) Handle(w http.ResponseWriter, r *http.Request) {
	var input user_entity.CreateUserInputDto

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		web_response_manager.RespondInternalServerError(w, err)
		return
	}

	validationErrs := h.createUserValidation.Validate(&input)
	if validationErrs != nil {
		web_response_manager.RespondFieldErrorValidation(w, validationErrs)
		return
	}

	alreadyExistsErr, err := h.createUserValidation.AlreadyExists(*input.Email, "")
	if err != nil {
		web_response_manager.RespondInternalServerError(w, err)
		return
	}

	if alreadyExistsErr != nil {
		web_response_manager.RespondConflictError(w, alreadyExistsErr)
		return
	}

	loggedUserId := r.Context().Value(protocol_application.UserIdKey).(string)
	output, err := h.createUser.Execute(&input, loggedUserId)
	if err != nil {
		web_response_manager.RespondInternalServerError(w, err)
		return
	}

	web_response_manager.RespondCreated(w, "created", output)
}
