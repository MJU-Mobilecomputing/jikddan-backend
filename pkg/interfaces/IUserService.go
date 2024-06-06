package interfaces

import "github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"

type IUserService interface {
	Create(repository.CreateUserParams) (*repository.CreateUserRow, error)
	FindOneById(int64) (*repository.FindUserByIdRow, error)
}
