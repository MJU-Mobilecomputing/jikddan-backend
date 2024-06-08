package weekly

import (
	"log"
	"net/http"
	"strconv"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/customerror"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/pkg/interfaces"
	"github.com/labstack/echo/v4"
)

type WeeklyController struct {
	WeeklyService interfaces.IWeeklyService
}

func InitWeeklyController() *WeeklyController {
	return &WeeklyController{}
}

func (w *WeeklyController) GetWeeklySummaryController(ctx echo.Context) error {
	month, err := strconv.Atoi(ctx.QueryParam("month"))
	if err != nil {
		return customerror.InvalidParamError(err)
	}
	weekOfMonth, err := strconv.Atoi(ctx.QueryParam("week_num"))
	if err != nil {
		return customerror.InvalidParamError(err)
	}
	param := repository.FindWeeklySummaryParams{
		Month:   int32(month),
		WeekNum: int32(weekOfMonth),
	}
	summary, err := w.WeeklyService.FindWeeklySummary(param)
	if err != nil {
		log.Println(err)
		return err
	}
	return ctx.JSON(http.StatusOK, summary)
}

func (w WeeklyController) WithWeeklyService(service interfaces.IWeeklyService) WeeklyController {
	w.WeeklyService = service
	return w
}
