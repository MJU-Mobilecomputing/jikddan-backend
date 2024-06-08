package interfaces

import (
	"context"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type IRepository interface {
	CreateDiaryDay(context.Context, pgtype.Date) (repository.DiaryDay, error)
	CreateDiaryMenu(context.Context, repository.CreateDiaryMenuParams) (repository.DiaryMenu, error)
	FindDiaryDayWithDate(context.Context, pgtype.Date) (repository.DiaryDay, error)
	FindDiaryDayWithMenus(context.Context, pgtype.Date) (repository.DiaryDayView, error)
}
