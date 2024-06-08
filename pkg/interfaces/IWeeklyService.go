package interfaces

import (
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/utils"
)

type IWeeklyService interface {
	FindWeeklySummary(param repository.FindWeeklySummaryParams) (*utils.WeeklySummaryResponse, error)
}
