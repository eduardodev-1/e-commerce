package utils

import (
	"errors"
	"github.com/eduardodev-1/e-commerce/internal/models"
	"gopkg.in/encoder.v1"
)

func CheckPasswordRequest(passwordPair models.PasswordPair) error {
	encoding := encoder.NewBcryptEncoder()
	verify, err := encoding.Verify(passwordPair.HashedPassword, passwordPair.Password)
	if err != nil || verify == false {
		return err
	}
	if !verify {
		return errors.New("password is incorrect")
	}
	return nil
}
