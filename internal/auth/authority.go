package auth

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type Route struct {
	path   string
	method string
}

func ValidateRouteAuthority(username string, authorities []string, fiberRoute *fiber.Route, path string) error {
	// Acessando os detalhes da rota
	method := fiberRoute.Method
	route := Route{path: path, method: method}
	// Printando na tela para debug
	fmt.Printf("Username: %s\n", username)
	fmt.Printf("Method: %s\n", method)
	fmt.Printf("Path: %s\n", path)
	fmt.Printf("Authorities: %v\n", authorities)
	err := checkAuthority(authorities, route)
	if err != nil {
		return errors.New("access Denied")
	}
	return nil
}

func checkAuthority(authorities []string, route Route) error {

	return nil
}
