package utils

import (
	"fmt"
	"reflect"
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

func ConvertHexToRGBA(hex string) (int, int, int, int, error) {
	hex = strings.TrimPrefix(hex, "#")

	var r, g, b, a int
	switch len(hex) {
	case 3:
		r = hexToInt(string(hex[0]) + string(hex[0]))
		g = hexToInt(string(hex[1]) + string(hex[1]))
		b = hexToInt(string(hex[2]) + string(hex[2]))
		a = 255
	case 4:
		r = hexToInt(string(hex[0]) + string(hex[0]))
		g = hexToInt(string(hex[1]) + string(hex[1]))
		b = hexToInt(string(hex[2]) + string(hex[2]))
		a = hexToInt(string(hex[3]) + string(hex[3]))
	case 6:
		r = hexToInt(hex[0:2])
		g = hexToInt(hex[2:4])
		b = hexToInt(hex[4:6])
		a = 255
	case 8:
		r = hexToInt(hex[0:2])
		g = hexToInt(hex[2:4])
		b = hexToInt(hex[4:6])
		a = hexToInt(hex[6:8])
	default:
		return 0, 0, 0, 0, fmt.Errorf("invalid hex color length")
	}

	return r, g, b, a, nil
}

func hexToInt(hex string) int {
	value, _ := strconv.ParseUint(hex, 16, 8)
	return int(value)
}

func PrintFieldsAndValues(i interface{}) {
	v := reflect.ValueOf(i)
	t := reflect.TypeOf(i)

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		fmt.Printf("%s: %v\n", field.Name, value.Interface())
	}
}
