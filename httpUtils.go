package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func unauthorized() *echo.HTTPError {
	return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
}