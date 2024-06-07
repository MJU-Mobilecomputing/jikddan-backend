package diaryday

import (
	"context"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/customerror"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/pkg/interfaces"
	"github.com/go-playground/validator/v10"
)

type DiaryDayService struct {
	Repository interfaces.IRepository
	Validator  *validator.Validate
}

func InitDiaryDayService() *DiaryDayService {
	return &DiaryDayService{}
}

func (d *DiaryDayService) Create(param repository.CreateDiaryDayParams) (*repository.DiaryDay, error) {
	err := d.Validator.Struct(param)
	if err != nil {
		return nil, customerror.ValidateError(err)
	}
	diaryDay, err := d.Repository.CreateDiaryDay(context.Background(), param)
	if err != nil {
		return nil, err
	}
	return &diaryDay, nil
}

func (d DiaryDayService) WithRepository(repo interfaces.IRepository) DiaryDayService {
	d.Repository = repo
	return d
}

func (d DiaryDayService) WithValidator(validator *validator.Validate) DiaryDayService {
	d.Validator = validator
	return d
}
