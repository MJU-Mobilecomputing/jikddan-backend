package interfaces

import (
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type IDiaryDayService interface {
	Create(pgtype.Date) (*repository.DiaryDay, error)
	FindOneByDate(pgtype.Date) (*repository.DiaryDay, error)
	FindOneWithMenu(pgtype.Date) (*repository.DiaryDayView, error)
	FindSummary(pgtype.Date) (*repository.DiaryDailySummary, error)
}
