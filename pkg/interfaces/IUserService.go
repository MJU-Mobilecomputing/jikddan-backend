package interfaces

import "github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"

type IUserService interface {
	Create(repository.CreateUserParams) (*repository.User, error)
	FindOneById(int64) (*repository.User, error)
}
