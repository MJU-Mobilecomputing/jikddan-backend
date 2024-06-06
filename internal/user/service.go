package user

import (
	"context"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/customerror"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/pkg/interfaces"
	"github.com/go-playground/validator/v10"
)

type UserService struct {
	Repository interfaces.IRepository
	Validator  *validator.Validate
}

func InitUserService() *UserService {
	return &UserService{}
}

func (u *UserService) Create(param repository.CreateUserParams) (*repository.CreateUserRow, error) {
	err := u.Validator.Struct(param)
	if err != nil {
		return nil, customerror.ValidateError(err)
	}
	user, err := u.Repository.CreateUser(context.Background(), param)
	if err != nil {
		return nil, customerror.UserCreationFail(err)
	}
	return &user, nil
}

func (u *UserService) FindOneById(id int64) (*repository.FindUserByIdRow, error) {
	user, err := u.Repository.FindUserById(context.Background(), id)
	if err != nil {
		return nil, customerror.UserNotFound(err)
	}
	return &user, nil
}

func (u UserService) WithRepository(repo interfaces.IRepository) UserService {
	u.Repository = repo
	return u
}

func (u UserService) WithValidator(validator *validator.Validate) UserService {
	u.Validator = validator
	return u
}
