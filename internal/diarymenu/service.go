package diarymenu

import (
	"context"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/customerror"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/pkg/interfaces"
	"github.com/go-playground/validator/v10"
)

type DiaryMenuService struct {
	Repository interfaces.IRepository
	Validator  *validator.Validate
}

func (d *DiaryMenuService) Create(param repository.CreateDiaryMenuParams) (*repository.DiaryMenu, error) {
	err := d.Validator.Struct(param)
	if err != nil {
		return nil, customerror.ValidateError(err)
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

func (d DiaryMenuService) WithValidator(validator *validator.Validate) DiaryMenuService {
	d.Validator = validator
	return d
}
