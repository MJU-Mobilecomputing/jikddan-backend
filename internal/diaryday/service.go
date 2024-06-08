package diaryday

import (
	"context"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/pkg/interfaces"
	"github.com/jackc/pgx/v5/pgtype"
)

type DiaryDayService struct {
	Repository interfaces.IRepository
}

func (d *DiaryDayService) Create(date pgtype.Date) (*repository.DiaryDay, error) {
	diaryDay, err := d.Repository.CreateDiaryDay(context.Background(), date)
	if err != nil {
		return nil, err
	}
	return &diaryDay, nil
}

func (d *DiaryDayService) FindOneByDate(date pgtype.Date) (*repository.DiaryDay, error) {
	diaryDay, err := d.Repository.FindDiaryDayWithDate(context.Background(), date)
	if err != nil {
		return nil, err
	}
	return &diaryDay, nil
}

func InitDiaryDayService() *DiaryDayService {
	return &DiaryDayService{}
}

func (d DiaryDayService) WithRepository(repo interfaces.IRepository) DiaryDayService {
	d.Repository = repo
	return d
}
