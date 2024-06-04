package user_entity

import util_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/util"

type UserDetailOutputDto struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Birth     string `json:"birth"`
	Active    bool   `json:"active"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	CreatedId string `json:"createdId"`
	UpdatedId string `json:"updatedId"`
}

func NewUserDetailOutputDto(user *User) *UserDetailOutputDto {
	birth := util_entity.GetStringFromDate(user.Birth)
	createdAt := util_entity.GetStringFromDateTime(user.CreatedAt)
	updatedAt := util_entity.GetStringFromDateTime(user.UpdatedAt)

	return &UserDetailOutputDto{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Birth:     birth,
		Active:    user.Active,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		CreatedId: user.CreatedId,
		UpdatedId: user.UpdatedId,
	}
}
