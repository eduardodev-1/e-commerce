package auth

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func ValidateRouteAuthority(username string, authorities []string, route *fiber.Route) {
	// Acessando os detalhes da rota
	method := route.Method
	path := route.Path

	// Printando na tela para debug
	fmt.Printf("Username: %s\n", username)
	fmt.Printf("Method: %s\n", method)
	fmt.Printf("Path: %s\n", path)
	fmt.Printf("Authorities: %v\n", authorities)
}
