package app

import (
	"fmt"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/middlewares"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/user"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/pkg/interfaces"
	"github.com/labstack/echo/v4/middleware"
)

var API_VER = "/api/v1%s"

func (app *Application) InitRoutes() {
	userService := user.InitUserService().WithRepository(&app.Repository).WithValidator(app.Validator)
	app.Handler.HTTPErrorHandler = CustomErrorHandler
	app.Handler.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	app.InitUserRoutes(&userService)
}

func (app *Application) InitUserRoutes(userService interfaces.IUserService) {
	e := app.Handler.Group(fmt.Sprintf(API_VER, "/user"))
	customMiddleware := middlewares.InitCustomMiddleware[repository.CreateUserParams]().WithValidator(app.Validator)
	userController := user.InitUserController().WithUserService(userService)
	e.POST("", userController.SignupUserController, customMiddleware.ValidateParam, customMiddleware.InterceptPassword)
}
