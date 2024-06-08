package diarymenu

import (
	"net/http"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/customerror"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/pkg/interfaces"
	"github.com/labstack/echo/v4"
)

type DiaryMenuController struct {
	DiaryMenuService interfaces.IDiaryMenuService
}

func InitDiaryMenuController() *DiaryMenuController {
	return &DiaryMenuController{}
}

func (d *DiaryMenuController) CreateDiaryMenuController(ctx echo.Context) error {
	var param repository.CreateDiaryMenuParams
	err := ctx.Bind(&param)
	if err != nil {
		return customerror.InternalServerError(err)
	}
	diaryMenu, err := d.DiaryMenuService.Create(param)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, diaryMenu)
}

func (d DiaryMenuController) WithDiaryMenuService(service interfaces.IDiaryMenuService) DiaryMenuController {
	d.DiaryMenuService = service
	return d
}
