package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CustomErrorHandler(err error, ctx echo.Context) {
	ctx.JSON(http.StatusInternalServerError, err)
}
