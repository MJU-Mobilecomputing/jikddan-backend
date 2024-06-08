package interfaces

import "github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"

type IDiaryMenuService interface {
	Create(repository.CreateDiaryMenuParams) (*repository.DiaryMenu, error)
}
