package interfaces

import "github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"

type IDiaryDayService interface {
	Create(repository.CreateDiaryDayParams) (*repository.DiaryDay, error)
}
