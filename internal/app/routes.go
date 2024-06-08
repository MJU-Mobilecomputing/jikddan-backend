package app

import (
	"fmt"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/diaryday"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/diarymenu"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/pkg/interfaces"
	"github.com/labstack/echo/v4/middleware"
)

var API_VER = "/api/v1%s"

func (app *Application) InitRoutes() {
	diaryDayService := diaryday.InitDiaryDayService().WithRepository(&app.Repository)
	diaryMenuService := diarymenu.InitDiaryMenuService().WithRepository(&app.Repository).WithDiaryDayService(&diaryDayService)
	app.Handler.HTTPErrorHandler = CustomErrorHandler
	app.Handler.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	app.InitDiaryMenuRoutes(&diaryMenuService)
	app.InitDiaryDayRoutes(&diaryDayService)
}

func (app *Application) InitDiaryMenuRoutes(service interfaces.IDiaryMenuService) {
	e := app.Handler.Group(fmt.Sprintf(API_VER, "/menu"))
	diaryMenuController := diarymenu.InitDiaryMenuController().WithDiaryMenuService(service)
	e.POST("", diaryMenuController.CreateDiaryMenuController)
}

func (app *Application) InitDiaryDayRoutes(service interfaces.IDiaryDayService) {
	e := app.Handler.Group(fmt.Sprintf(API_VER, "/day"))
	diaryDayController := diaryday.InitDiaryDayController().WithDiaryDayService(service)
	e.GET("/:date", diaryDayController.FindDiaryDayWithMenu)
}
