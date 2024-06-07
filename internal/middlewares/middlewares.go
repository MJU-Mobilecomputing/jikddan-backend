package middlewares

import (
	"errors"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/constants"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/customerror"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type CustomMiddleware[T any] struct {
	param     T
	validator *validator.Validate
}

func (customMiddleware *CustomMiddleware[T]) ValidateParam(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var param T
		err := ctx.Bind(&param)
		if err != nil {
			return customerror.InvalidParamError(err)
		}
		err = customMiddleware.validator.Struct(param)
		if err != nil {
			return customerror.ValidateError(err)
		}
		ctx.Set(constants.PARAM_KEY, &param)
		return next(ctx)
	}
}

func (c *CustomMiddleware[T]) InterceptPassword(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		if param, ok := ctx.Get(constants.PARAM_KEY).(repository.CreateUserParams); ok {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(param.Password), bcrypt.DefaultCost)
			if err != nil {
				return customerror.InternalServerError(err)
			}
			param.Password = string(hashedPassword)
			ctx.Set(constants.USER_PARAM_KEY, &param)
			return next(ctx)
		} else {
			return customerror.InternalServerError(errors.New("userParam context does not exists"))
		}
	}
}

func InitCustomMiddleware[T any]() *CustomMiddleware[T] {
	return &CustomMiddleware[T]{}
}

func (c CustomMiddleware[T]) WithValidator(validator *validator.Validate) CustomMiddleware[T] {
	c.validator = validator
	return c
}
