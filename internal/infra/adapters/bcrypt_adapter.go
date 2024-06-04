package infra_adapters

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptAdapter struct{}

func NewBcryptAdapter() *BcryptAdapter {
	return &BcryptAdapter{}
}

func (b *BcryptAdapter) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (b *BcryptAdapter) CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
