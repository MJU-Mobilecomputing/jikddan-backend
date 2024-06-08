package diarymenu

import (
	"log"
	"net/http"
	"time"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/customerror"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/utils"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/pkg/interfaces"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

type DiaryMenuController struct {
	DiaryMenuService interfaces.IDiaryMenuService
	S3Service        interfaces.IS3Service
	GPTService       interfaces.IGPTService
}

func InitDiaryMenuController() *DiaryMenuController {
	return &DiaryMenuController{}
}

func (d *DiaryMenuController) CreateDiaryMenuController(ctx echo.Context) error {
	file, err := ctx.FormFile("img")
	if err != nil {
		log.Println(err)
		return customerror.InternalServerError(err)
	}
	dateValue := ctx.FormValue("date")
	menuTimeValue := ctx.FormValue("menu_time")

	parsedTime, err := time.Parse("2006-01-02", dateValue)
	if err != nil {
		return customerror.InternalServerError(err)
	}
	menuTime := repository.NullMenuTime{
		MenuTime: repository.MenuTime(menuTimeValue),
		Valid:    true,
	}
	date := pgtype.Date{
		Time:  parsedTime,
		Valid: true,
	}
	body := utils.CreateMenuRequest{
		MenuTime: menuTime,
		Date:     date,
		Img:      *file,
	}
	diaryMenu, err := d.DiaryMenuService.Create(file, body)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, diaryMenu)
}

/*
	func (d *DiaryMenuController) ImageUploadTest(ctx echo.Context) error {
		file, err := ctx.FormFile("image")
		if err != nil {
			return customerror.InternalServerError(err)
		}
		src, err := file.Open()
		if err != nil {
			return customerror.InternalServerError(err)
		}
		fileName := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
		fileUrl, err := d.S3Service.UploadFile(&src, fileName)
		if err != nil {
			log.Println(err)
			return customerror.InternalServerError(err)
		}
		resp, err := d.GPTService.GetMenuNutrient(*fileUrl)
		if err != nil {
			log.Println(err)
			return customerror.InternalServerError(err)
		}
		return ctx.JSON(http.StatusOK, resp)
	}
*/
func (d DiaryMenuController) WithDiaryMenuService(service interfaces.IDiaryMenuService) DiaryMenuController {
	d.DiaryMenuService = service
	return d
}
