package diarymenu

import (
	"context"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/customerror"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/pkg/interfaces"
)

type DiaryMenuService struct {
	Repository      interfaces.IRepository
	DiaryDayService interfaces.IDiaryDayService
}

func (d *DiaryMenuService) Create(param repository.CreateDiaryMenuParams) (*repository.DiaryMenu, error) {
	diaryMenu, err := d.Repository.CreateDiaryMenu(context.Background(), param)
	if err != nil {
		return nil, customerror.DiaryMenuCreationFail(err)
	}
	_, err = d.DiaryDayService.FindOneByDate(param.Date)
	if err != nil {
		_, err := d.DiaryDayService.Create(param.Date)
		if err != nil {
			return nil, err
		}
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
