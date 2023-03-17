package main

import (
	"log"
	"strings"

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
	
	keycloakConfig := keycloak.DefaultKeycloakConfig
	keycloakConfig.KeycloakURL = config.Authentication.BaseUrl
	keycloakConfig.KeycloakRealm = config.Authentication.Realm
	keycloakConfig.Skipper = func(c echo.Context) bool {
		log.Println(c.Request().RequestURI)
		if strings.HasPrefix(c.Request().RequestURI, "/public") {
			return true
		}
		return false
	}

	e.Use(keycloak.KeycloakWithConfig(keycloakConfig))
	e.Use(extractUsername)


	addKubernetesController()

	log.Println("Onyxia onboarding ...")
	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func extractUsername(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		if (user == nil) {
			log.Println("User is not logged in")
			return nil
		}
		token := user.(*jwt.Token)
		claims := token.Claims.(*jwt.MapClaims)
		username := (*claims)["preferred_username"]
		c.Set("username", username)
		return next(c)
	}
}



