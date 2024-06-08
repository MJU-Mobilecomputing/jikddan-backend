package diaryday

import (
	"log"
	"net/http"
	"time"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/customerror"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/pkg/interfaces"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type DiaryDayController struct {
	DiaryDayService interfaces.IDiaryDayService
}

func InitDiaryDayController() *DiaryDayController {
	return &DiaryDayController{}
}

func (d *DiaryDayController) FindDiaryDayWithMenu(ctx echo.Context) error {
	param := ctx.Param("date")
	log.Println(param)
	var date pgtype.Date
	parsedTime, err := time.Parse("2006-01-02", param)
	log.Println(parsedTime)
	if err != nil {
		return customerror.InvalidParamError(err)
	}
	date.Time = parsedTime
	date.Valid = true
	log.Println(date)
	view, err := d.DiaryDayService.FindOneWithMenu(date)
	if err != nil {
		log.Println(err)
		return err
	}
	return ctx.JSON(http.StatusOK, view)
}

func (d DiaryDayController) WithDiaryDayService(service interfaces.IDiaryDayService) DiaryDayController {
	d.DiaryDayService = service
	return d
}
