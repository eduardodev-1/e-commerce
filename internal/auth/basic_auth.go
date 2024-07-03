package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"math/big"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func CheckAppCredentials(c *fiber.Ctx) error {
	// Get authorization header
	auth := c.Get(fiber.HeaderAuthorization)

	// Check if the header contains content besides "basic".
	if len(auth) <= 6 || !utils.EqualFold(auth[:6], "basic ") {
		return errors.New("auth Method Fail")
	}

	// Decode the header contents
	raw, err := base64.StdEncoding.DecodeString(auth[6:])
	if err != nil {
		return errors.New("auth Check Fail")
	}
	// Get the credentials
	creds := utils.UnsafeString(raw)
	// Check if the credentials are in the correct form
	// which is "username:password".
	index := strings.Index(creds, ":")
	if index == -1 {
		return errors.New("auth Format Fail")
	}
	// Get the username and password
	username := creds[:index]
	password := creds[index+1:]
	err = checkAppCredentials(username, password)
	if err != nil {
		return err
	}
	return nil
}

func checkAppCredentials(username string, password string) error {
	var clientId = os.Getenv("CLIENT_ID")
	var clientSecret = os.Getenv("CLIENT_SECRET")
	if username != clientId || password != clientSecret {
		return errors.New("invalid App Credentials")
	}
	return nil
}
func generateSecurePassword(length int) (string, error) {
	// Caracteres permitidos na senha
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+"

	// Tamanho do conjunto de caracteres
	charsetLen := big.NewInt(int64(len(charset)))

	// Gerar uma senha aleatória usando criptografia segura
	var password string
	for i := 0; i < length; i++ {
		// Gerar um índice aleatório
		randomIndex, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return "", err
		}
		// Obter o caractere correspondente ao índice aleatório
		password += string(charset[randomIndex.Int64()])
	}

	return password, nil
	//exemplo para usar a função de gerar senha segura de 12 caracteres
	//// Gerar uma senha segura com 12 caracteres
	//password, err := generateSecurePassword(12)
	//if err != nil {
	//	panic(err)
	//}
	//println("Senha gerada:", password)
}
