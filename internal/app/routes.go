package app

import (
	"fmt"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/diarymenu"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/middlewares"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/user"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/pkg/interfaces"
	"github.com/labstack/echo/v4/middleware"
)

var API_VER = "/api/v1%s"

func (app *Application) InitRoutes() {
	userService := user.InitUserService().WithRepository(&app.Repository).WithValidator(app.Validator)
	diaryMenuService := diarymenu.InitDiaryMenuService().WithRepository(&app.Repository).WithValidator(app.Validator)
	app.Handler.HTTPErrorHandler = CustomErrorHandler
	app.Handler.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	app.InitUserRoutes(&userService)
	app.InitDiaryMenuRoutes(&diaryMenuService)
}

func (app *Application) InitUserRoutes(userService interfaces.IUserService) {
	e := app.Handler.Group(fmt.Sprintf(API_VER, "/user"))
	customMiddleware := middlewares.InitCustomMiddleware[repository.CreateUserParams]().WithValidator(app.Validator)
	userController := user.InitUserController().WithUserService(userService)
	e.POST("", userController.SignupUserController, customMiddleware.ValidateParam, customMiddleware.InterceptPassword)
}

func (app *Application) InitDiaryMenuRoutes(diaryMenu interfaces.IDiaryMenuService) {
	e := app.Handler.Group(fmt.Sprintf(API_VER, "/menu"))
	customMiddleware := middlewares.InitCustomMiddleware[repository.CreateDiaryMenuParams]().WithValidator(app.Validator)
	diaryMenuController := diarymenu.InitController().WithDiaryMenuService(diaryMenu)
	e.POST("", diaryMenuController.CreateDiaryMenuController, customMiddleware.ValidateParam)
}
