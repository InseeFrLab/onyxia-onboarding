package main

import (
	"log"

	keycloak "github.com/baba2k/echo-keycloak"
	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var e = echo.New()

func main() {
	loadConfiguration()
	initKubernetesClient()


	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	log.Println("Using authentication from "+config.Authentication.BaseUrl+" realm "+config.Authentication.Realm)
	e.Use(keycloak.Keycloak(config.Authentication.BaseUrl, config.Authentication.Realm))
	e.Use(extractUsername)


	addKubernetesController()

	log.Println("Onyxia onboarding ...")
	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func extractUsername(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(*jwt.MapClaims)
		username := (*claims)["preferred_username"]
		c.Set("username", username)
		return next(c)
	}
}



