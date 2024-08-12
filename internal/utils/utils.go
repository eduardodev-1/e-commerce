package utils

import (
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/encoder.v1"
	"gopkg.in/encoder.v1/types"
)

func VerifyEncode(data string, hash string) {
	// Using the default options
	// types.Argon2 types.Pbkdf2 types.Bcrypt types.Hkdf types.Scrypt
	encoding := encoder.New(types.Bcrypt) // or use encoder.NewBcryptEncoder()
	hash, err := encoding.Encode(data)
	fmt.Println(hash)
	verify, err := encoding.Verify("$2a$10$btdcIbJYWyZDbN85wSaYduZuOROvyjNee9azaI5/GEyGuO8HfHN3G", data)
	if err != nil {
		return
	}
	if verify {
		fmt.Println("match: " + strconv.FormatBool(verify))
	}
}

// JoinStrings une uma lista de strings em uma única string usando o separador especificado
func JoinStrings(strs []string) string {
	separator := ", "
	return strings.Join(strs, separator)
}
func SplitString(s string) []string {
	separator := ", "
	return strings.Split(s, separator)
}
