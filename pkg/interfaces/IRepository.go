package interfaces

import (
	"context"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
)

type IRepository interface {
	CreateUser(context.Context, repository.CreateUserParams) (repository.CreateUserRow, error)
	FindUserById(context.Context, int64) (repository.FindUserByIdRow, error)
	CreateDiaryDay(context.Context, int64) (repository.DiaryDay, error)
}
