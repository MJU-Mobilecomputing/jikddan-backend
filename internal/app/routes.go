package app

import (
	"fmt"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/aws"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/diaryday"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/diarymenu"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/gpt"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/weekly"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/pkg/interfaces"
	"github.com/labstack/echo/v4/middleware"
)

var API_VER = "/api/v1%s"

func (app *Application) InitRoutes() {
	s3Service := aws.InitS3Service().WithClient(app.S3Client)
	openAIRepository := gpt.InitGPTRepository().WithClient(app.OpenAIClient)
	openAIService := gpt.InitGPTService().WithRepository(&openAIRepository)
	diaryDayService := diaryday.InitDiaryDayService().WithRepository(&app.Repository)
	diaryMenuService := diarymenu.InitDiaryMenuService().WithRepository(&app.Repository).WithDiaryDayService(&diaryDayService).WithS3Service(&s3Service).WithGPTService(&openAIService)

	weeklyService := weekly.InitWeeklyService().WithRepository(&app.Repository).WithGPTService(&openAIService)
	app.Handler.HTTPErrorHandler = CustomErrorHandler
	app.Handler.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	app.InitDiaryMenuRoutes(&diaryMenuService, &s3Service, &openAIService)
	app.InitDiaryDayRoutes(&diaryDayService)
	app.InitWeeklySummaryRoutes(&weeklyService)
}

func (app *Application) InitDiaryMenuRoutes(service interfaces.IDiaryMenuService, s3Service interfaces.IS3Service, gptService interfaces.IGPTService) {
	e := app.Handler.Group(fmt.Sprintf(API_VER, "/menu"))
	diaryMenuController := diarymenu.InitDiaryMenuController().WithDiaryMenuService(service)
	e.POST("", diaryMenuController.CreateDiaryMenuController)
}

func (app *Application) InitDiaryDayRoutes(service interfaces.IDiaryDayService) {
	e := app.Handler.Group(fmt.Sprintf(API_VER, "/diary"))
	diaryDayController := diaryday.InitDiaryDayController().WithDiaryDayService(service)
	e.GET("/:date", diaryDayController.FindDiaryDayWithMenu)
	e.GET("/:date/summary", diaryDayController.FindSummary)
}

func (app *Application) InitWeeklySummaryRoutes(service interfaces.IWeeklyService) {
	e := app.Handler.Group(fmt.Sprintf(API_VER, "/weekly"))
	weeklyController := weekly.InitWeeklyController().WithWeeklyService(service)
	e.GET("", weeklyController.GetWeeklySummaryController)
}
