package diaryday

import (
	"net/http"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/utils"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/pkg/interfaces"
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
	date, err := utils.ParseTime(param)
	if err != nil {
		return err
	}
	view, err := d.DiaryDayService.FindOneWithMenu(*date)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, view)
}

func (d *DiaryDayController) FindSummary(ctx echo.Context) error {
	param := ctx.Param("date")
	date, err := utils.ParseTime(param)
	if err != nil {
		return err
	}
	summary, err := d.DiaryDayService.FindSummary(*date)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, summary)
}

func (d DiaryDayController) WithDiaryDayService(service interfaces.IDiaryDayService) DiaryDayController {
	d.DiaryDayService = service
	return d
}
