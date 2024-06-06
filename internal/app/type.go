package app

import (
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/config"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type Application struct {
	Port       int
	DB         *pgxpool.Pool
	Handler    *echo.Echo
	Repository repository.Queries
	Validator  *validator.Validate
	Config     config.Config
}
