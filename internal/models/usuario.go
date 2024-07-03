package models

import (
	"errors"
	"gopkg.in/encoder.v1"
)

type Usuario struct {
	Id              int    `json:"id"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	Enabled         bool   `json:"enabled"`
	AccountExpired  bool   `json:"account_expired" db:"account_expired"`
	AccountLocked   bool   `json:"account_locked" db:"account_locked"`
	PasswordExpired bool   `json:"password_expired" db:"password_expired"`
	ProjetoPadraoId int    `json:"projeto_padrao_id" db:"projeto_padrao_id"`
	TenantId        int    `json:"tenant_id" db:"tenant_id"`
}
type PasswordPair struct {
	Password       string
	HashedPassword string
}

func (passwordPair *PasswordPair) CheckPasswordRequest() error {
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

// AUTHENTICATION
type CredentialsToAuthenticate struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
type UsuarioToAuthenticate struct {
	Username string `db:"username"`
	Password string `db:"password"`
	TenantId string `db:"tenant_id"`
}

type AuthenticatedUser struct {
	Id       int    `json:"id"  form:"id"  db:"id"`
	Username string `json:"username"  form:"username"  db:"username"`
}
