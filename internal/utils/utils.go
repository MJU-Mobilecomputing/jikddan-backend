package utils

import (
	"errors"
	"regexp"
	"time"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/constants"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/customerror"
	"github.com/jackc/pgx/v5/pgtype"
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

func ParseTime(timeStr string) (*pgtype.Date, error) {
	var date pgtype.Date
	parsedTime, err := time.Parse("2006-01-02", timeStr)
	if err != nil {
		return nil, err
	}
	date.Time = parsedTime
	date.Valid = true
	return &date, nil
}
