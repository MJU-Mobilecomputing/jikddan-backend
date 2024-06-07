package diarymenu

import (
	"net/http"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/utils"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/pkg/interfaces"
	"github.com/labstack/echo/v4"
)

type DiaryMenuController struct {
	DiaryMenuService interfaces.IDiaryMenuService
}

func InitController() *DiaryMenuController {
	return &DiaryMenuController{}
}

func (d *DiaryMenuController) CreateDiaryMenuController(ctx echo.Context) error {
	param, err := utils.GetParam[repository.CreateDiaryMenuParams](ctx)
	if err != nil {
		return err
	}
	if param.DiaryDayID == 0 {
		// create diary day
	}
	diaryMenu, err := d.DiaryMenuService.Create(*param)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, diaryMenu)
}

func (d DiaryMenuController) WithDiaryMenuService(service interfaces.IDiaryMenuService) DiaryMenuController {
	d.DiaryMenuService = service
	return d
}
