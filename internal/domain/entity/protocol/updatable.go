package protocol_entity

import (
	"time"

	util_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/util"
)

type Updatable struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedId string
	UpdatedId string
}

func (u *Updatable) FillUpdatableFieldsForCreate(id string) {
	now := time.Now()
	u.CreatedId = id
	u.UpdatedId = id
	u.CreatedAt = now
	u.UpdatedAt = now
}

func (u *Updatable) FillUpdatableFieldsForUpdate(id string) {
	u.UpdatedId = id
	u.UpdatedAt = time.Now()
}

func (u *Updatable) FillUpdatableFieldsFromMap(mappedResult map[string]string) error {
	if mappedResult["created_at"] != "" {
		parsedDate, err := util_entity.GetDateTimeFromString(mappedResult["created_at"])
		if err != nil {
			return err
		}
		u.CreatedAt = parsedDate
	}

	if mappedResult["updated_at"] != "" {
		parsedDate, err := util_entity.GetDateTimeFromString(mappedResult["updated_at"])
		if err != nil {
			return err
		}
		u.UpdatedAt = parsedDate
	}

	if mappedResult["created_id"] != "" {
		u.CreatedId = mappedResult["created_id"]
	}

	if mappedResult["updated_id"] != "" {
		u.UpdatedId = mappedResult["updated_id"]
	}

	return nil
}

func (u *Updatable) MapUpdatableFields(updatableMap map[string]interface{}) map[string]interface{} {
	updatableMap["created_at"] = u.CreatedAt
	updatableMap["updated_at"] = u.UpdatedAt
	updatableMap["created_id"] = u.CreatedId
	updatableMap["updated_id"] = u.UpdatedId

	return updatableMap
}
