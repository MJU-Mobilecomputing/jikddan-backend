package app

import (
	"net/http"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/customerror"
	"github.com/labstack/echo/v4"
)

func CustomErrorHandler(err error, ctx echo.Context) {
	code := http.StatusInternalServerError
	if customError, ok := err.(*customerror.CustomError); ok {
		code = customError.StatusCode
		ctx.JSON(code, customError)
	} else if echoError, ok := err.(*echo.HTTPError); ok {
		switch echoError.Code {
		case http.StatusMethodNotAllowed:
			ctx.JSON(echoError.Code, customerror.MethodNotAllowed(echoError))
		case http.StatusNotFound:
			ctx.JSON(echoError.Code, customerror.PageNotFound(echoError))
		case http.StatusInternalServerError:
			ctx.JSON(echoError.Code, customerror.InternalServerError(echoError))
		}
	}
}
