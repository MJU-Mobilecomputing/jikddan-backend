package interfaces

import (
	"mime/multipart"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/utils"
)

type IDiaryMenuService interface {
	Create(*multipart.FileHeader, utils.CreateMenuRequest) (*repository.DiaryMenu, error)
}
