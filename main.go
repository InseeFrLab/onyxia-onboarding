package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	keycloak "github.com/baba2k/echo-keycloak"
	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

var config configuration

func main() {
	loadConfiguration()
	initKubernetesClient()

  
	
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	fmt.Println("Using authentication from "+config.Authentication.BaseUrl+" realm "+config.Authentication.Realm)
	e.Use(keycloak.Keycloak(config.Authentication.BaseUrl, config.Authentication.Realm))

	// Routes
	e.GET("/", hello)


	fmt.Println("Onyxia onboarding ...")
	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(*jwt.MapClaims)
		prettyJSONClaims, _ := json.MarshalIndent(claims, "", "   ")
		return c.String(http.StatusOK, fmt.Sprintf(
			fmt.Sprintf("There are %d pods in the cluster", countPods())+" Hello, User! Your claims are:\n%+v\n", string(prettyJSONClaims)))
  }

func loadConfiguration() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	viper.SetConfigName("config.local")
	viper.AddConfigPath(".")
	viper.MergeInConfig()

	err := viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
}