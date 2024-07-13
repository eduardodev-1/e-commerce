package utils

import (
	"e-commerce/internal/core/domain"
	"errors"
	"gopkg.in/encoder.v1"
)

func CheckPasswordRequest(passwordPair domain.PasswordPair) error {
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
