package infra_adapters

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	auth_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/auth"
)

type jwtAdapter struct{}

const secretKey = "abliblablubla"

func NewJwtAdapter() *jwtAdapter {
	return &jwtAdapter{}
}

func (j *jwtAdapter) GenerateToken(id string, email string, timeToExpire int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"id":    id,
		"exp":   time.Now().Add(time.Hour + time.Duration(timeToExpire)).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func (j *jwtAdapter) ParseToken(token string) (*auth_entity.TokenOutputDto, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return &auth_entity.TokenOutputDto{
			Email: claims["email"].(string),
			ID:    claims["id"].(string),
		},
		nil
}
