package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func addKubernetesController() {
	// Routes
	e.GET("/kubernetes", getPodsCount)
	e.GET("/kubernetes/namespace/:name", testNamespace)
}

func testNamespace(c echo.Context) error {
	namespaceName := c.Param("name")
	username := c.Get("username").(string)
	if (!hasPermissionOnNamespace(username,namespaceName)) {
		return unauthorized()
	}
	
	return c.JSON(http.StatusOK, &Result{ Result: namespaceExists(namespaceName)})
}

// Handler
func getPodsCount(c echo.Context) error {
	return c.JSON(http.StatusOK, &Result{Result: countPods()})
  }