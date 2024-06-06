package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (app *Application) InitRoutes() {
	app.Handler.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	app.InitHomeRoutes()
}

func (app *Application) InitHomeRoutes() {
	e := app.Handler.Group("/home")
	e.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello")
	})
}
