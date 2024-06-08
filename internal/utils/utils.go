package utils

import (
	"errors"
	"regexp"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/constants"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/customerror"
	"github.com/labstack/echo/v4"
)

func GetParam[T any](ctx echo.Context) (*T, error) {
	if param, ok := ctx.Get(constants.PARAM_KEY).(*T); ok {
		return param, nil
	}
	return nil, customerror.InternalServerError(errors.New("cannot retrive param"))
}

func GetJsonString(str string) string {
	re := regexp.MustCompile(`\{[\s\S]*\}`)
	jsonString := re.FindString(str)
	return jsonString
}
