package middlewares

import (
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/constants"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/customerror"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
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

func InitCustomMiddleware[T any]() *CustomMiddleware[T] {
	return &CustomMiddleware[T]{}
}

func (c CustomMiddleware[T]) WithValidator(validator *validator.Validate) CustomMiddleware[T] {
	c.validator = validator
	return c
}
