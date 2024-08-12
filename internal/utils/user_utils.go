package utils

import (
	"e-commerce/internal/core/domain/models"
	httpError "e-commerce/internal/httperror"
	"errors"
)

const (
	Seller = "seller"
	Client = "client"
)

func CheckUserType(userType string) *httpError.ErrorParams {
	errorParams := new(httpError.ErrorParams)
	if userType != Seller && userType != Client {
		errorParams.SetDefaultParams(errors.New("invalid User type"))
		return errorParams
	}
	return nil
}

func EncryptPassword(originalPassword string) (string, *httpError.ErrorParams) {
	var err error
	var errorParams = new(httpError.ErrorParams)
	var passwordPair = new(models.PasswordPair)
	hashedPassword, err := passwordPair.SetOriginalPasswordAndGetHashedPassword(originalPassword)
	if err != nil {
		errorParams.SetDefaultParams(err)
		return "", errorParams
	}
	return hashedPassword, nil
}

func CheckUsername(loggedUsername string, requestUsername string) *httpError.ErrorParams {
	var errorParams = new(httpError.ErrorParams)
	if loggedUsername != requestUsername {
		errorParams.SetDefaultParams(errors.New("invalid Username"))
		return errorParams
	}
	return nil
}
