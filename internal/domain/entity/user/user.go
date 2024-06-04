package user_entity

import (
	"strings"
	"time"

	"github.com/google/uuid"
	protocol_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/protocol"
	util_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/util"
)

type User struct {
	protocol_entity.Updatable
	ID       string
	Name     string
	Password string
	Active   bool
	Email    string
	Birth    time.Time
}

func NewUser(name, password, email string, birth time.Time) *User {
	userId := uuid.New().String()

	user := &User{
		ID:       userId,
		Name:     name,
		Password: password,
		Active:   true,
		Email:    email,
		Birth:    birth,
	}

	return user
}

func NewUserFromMap(mappedResult map[string]string) (*User, error) {
	user := &User{}

	if mappedResult["ID"] != "" {
		user.ID = mappedResult["ID"]
	}

	if mappedResult["id"] != "" {
		user.ID = mappedResult["id"]
	}

	if mappedResult["name"] != "" {
		user.Name = mappedResult["name"]
	}

	if mappedResult["email"] != "" {
		user.Email = mappedResult["email"]
	}

	if mappedResult["password"] != "" {
		user.Password = mappedResult["password"]
	}

	if mappedResult["active"] != "" {
		user.Active = strings.ToLower(mappedResult["active"]) == "true" || mappedResult["active"] == "1"
	}

	if mappedResult["birth"] != "" {
		parsedDate, err := util_entity.GetDateFromString(mappedResult["birth"])
		if err != nil {
			return nil, err
		}
		user.Birth = parsedDate
	}

	err := user.FillUpdatableFieldsFromMap(mappedResult)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Map(withId bool) map[string]interface{} {
	userMap := map[string]interface{}{
		"email":    u.Email,
		"name":     u.Name,
		"password": u.Password,
		"active":   u.Active,
		"birth":    u.Birth,
	}

	if withId {
		userMap["ID"] = u.ID
	}

	return u.MapUpdatableFields(userMap)
}
