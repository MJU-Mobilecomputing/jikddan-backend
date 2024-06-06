package interfaces

import (
	"context"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
)

type IRepository interface {
	CreateUser(ctx context.Context, arg repository.CreateUserParams) (repository.User, error)
	FindUserById(context.Context, int64) (repository.User, error)
}
