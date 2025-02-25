package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/aws"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/config"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/db"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
)

var (
	once sync.Once
	app  *Application
)

func InitApplication() (*Application, error) {
	if app == nil {
		once.Do(func() {
			err := setApplication()
			if err != nil {
				log.Fatal(err)
			}
		})
	} else {
		return nil, errors.New("app is already configured")
	}
	return app, nil
}

func (app *Application) Run() error {
	if err := app.Handler.Start(fmt.Sprintf(":%d", app.Port)); err != nil {
		return err
	}
	return nil
}

func setApplication() error {
	appConfig, err := config.InitConfig()
	if err != nil {
		return err
	}
	conn, err := db.InitDB(context.Background(), appConfig.DB)
	if err != nil {
		return err
	}
	handler := echo.New()
	s3client, err := aws.InitS3Client(appConfig.AWS.Bucket, appConfig.AWS.Key, appConfig.AWS.PrivateKey)
	if err != nil {
		return err
	}
	openaiClient := openai.NewClient(appConfig.OpenAI.Key)
	app = &Application{
		Port:         appConfig.App.Port,
		DB:           conn,
		Repository:   *repository.New(conn),
		Handler:      handler,
		Config:       *appConfig,
		Validator:    validator.New(),
		S3Client:     s3client,
		OpenAIClient: openaiClient,
	}

	app.Validator.RegisterValidation("status", isValidStatus)
	app.Validator.RegisterValidation("menutime", isMenuTime)
	app.InitRoutes()
	return nil
}
