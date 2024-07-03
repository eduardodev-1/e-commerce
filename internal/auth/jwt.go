package auth

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	jtoken "github.com/golang-jwt/jwt/v4"
)

type customClaim struct {
	UserName    string   `json:"username"`
	Authorities []string `json:"authorities"`
	jtoken.RegisteredClaims
}

const (
	duration = time.Hour * 24
)

func NewJWToken(usuarioId int, username string, authorities []string) (string, error) {
	// Create claims with multiple fields populated
	claims := customClaim{
		UserName:    username,
		Authorities: authorities,
		RegisteredClaims: jtoken.RegisteredClaims{
			ExpiresAt: jtoken.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jtoken.NewNumericDate(time.Now()),
			NotBefore: jtoken.NewNumericDate(time.Now()),
			ID:        strconv.Itoa(usuarioId),
		},
	}

	// Create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jtoken.Token, error) {
	var myClaim = new(customClaim)

	token, err := jtoken.ParseWithClaims(tokenString, myClaim, func(token *jtoken.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jtoken.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token signature %v", token.Header["alg"])
	}
	return token, nil
}

func ValidateAndExtractTokenData(token *jtoken.Token) (string, []string, error) {
	var username string
	var authorities []string
	if claims, ok := token.Claims.(*customClaim); !ok {
		return "", nil, errors.New("unknown claims type")
	} else {
		username = claims.UserName
		authorities = claims.Authorities
	}

	return username, authorities, nil
}
