package interfaces

import "github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"

type IDiaryMenuService interface {
	Create(param repository.CreateDiaryMenuParams) (*repository.DiaryMenu, error)
}
