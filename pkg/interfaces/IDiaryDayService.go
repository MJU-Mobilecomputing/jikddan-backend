package interfaces

import (
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type IDiaryDayService interface {
	Create(pgtype.Date) (*repository.DiaryDay, error)
	FindOneByDate(pgtype.Date) (*repository.DiaryDay, error)
}
