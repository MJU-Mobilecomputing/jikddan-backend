package diarymenu

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"time"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/customerror"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/utils"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/pkg/interfaces"
)

type DiaryMenuService struct {
	Repository      interfaces.IRepository
	DiaryDayService interfaces.IDiaryDayService
	S3Service       interfaces.IS3Service
	GPTService      interfaces.IGPTService
}

func (d *DiaryMenuService) Create(file *multipart.FileHeader, body utils.CreateMenuRequest) (*repository.DiaryMenu, error) {
	var diaryDay *repository.DiaryDay
	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
	src, err := file.Open()
	if err != nil {
		return nil, customerror.InternalServerError(err)
	}
	fileUrl, err := d.S3Service.UploadFile(&src, filename)
	if err != nil {
		return nil, err
	}
	resp, err := d.GPTService.GetMenuNutrient(*fileUrl)
	if err != nil {
		return nil, err
	}

	diaryDay, err = d.DiaryDayService.FindOneByDate(body.Date)
	if err != nil {
		diaryDay, err = d.DiaryDayService.Create(body.Date)
		if err != nil {
			return nil, err
		}
	}
	log.Println(resp.Calrory)
	calrory := int32(resp.Calrory)
	salt := int32(resp.Salt)
	foodMoisture := int32(resp.FoodMoisture)
	score := int32(resp.Score)
	param := repository.CreateDiaryMenuParams{
		Summary:      &resp.Summary,
		TotalCal:     &calrory,
		Salt:         &salt,
		FoodMoisture: &foodMoisture,
		Img:          fileUrl,
		Date:         body.Date,
		MenuTime:     body.MenuTime,
		DiaryDayID:   &diaryDay.ID,
		Score:        &score,
	}
	diaryMenu, err := d.Repository.CreateDiaryMenu(context.Background(), param)
	if err != nil {
		return nil, customerror.DiaryMenuCreationFail(err)
	}
	return &diaryMenu, nil
}

func InitDiaryMenuService() *DiaryMenuService {
	return &DiaryMenuService{}
}

func (d DiaryMenuService) WithRepository(repo interfaces.IRepository) DiaryMenuService {
	d.Repository = repo
	return d
}

func (d DiaryMenuService) WithDiaryDayService(service interfaces.IDiaryDayService) DiaryMenuService {
	d.DiaryDayService = service
	return d
}

func (d DiaryMenuService) WithS3Service(service interfaces.IS3Service) DiaryMenuService {
	d.S3Service = service
	return d
}

func (d DiaryMenuService) WithGPTService(service interfaces.IGPTService) DiaryMenuService {
	d.GPTService = service
	return d
}
