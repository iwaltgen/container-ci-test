package main // import "github.com/iwaltgen/gitlab-ci-test"

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	addMiddlewares(e)
	addRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Middlewares
func addMiddlewares(e *echo.Echo) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	return e
}

// Routes
func addRoutes(e *echo.Echo) *echo.Echo {
	e.GET("/", hello)
	e.GET("/machine", machine)
	return e
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// Handler
func machine(c echo.Context) error {
	name, err := os.Hostname()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, name)
}
