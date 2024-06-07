package interfaces

import (
	"context"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
)

type IRepository interface {
	CreateUser(context.Context, repository.CreateUserParams) (repository.CreateUserRow, error)
	FindUserById(context.Context, int64) (repository.FindUserByIdRow, error)
	CreateDiaryDay(context.Context, repository.CreateDiaryDayParams) (repository.DiaryDay, error)
	CreateDiaryMenu(context.Context, repository.CreateDiaryMenuParams) (repository.DiaryMenu, error)
	FindOneDiaryWithMenu(context.Context, repository.FindOneDiaryWithMenuParams) (repository.DiaryDayView, error)
}
