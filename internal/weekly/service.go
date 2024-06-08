package weekly

import (
	"context"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/utils"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/pkg/interfaces"
)

type WeeklyService struct {
	Repository interfaces.IRepository
	GPTService interfaces.IGPTService
}

func InitWeeklyService() *WeeklyService {
	return &WeeklyService{}
}

func (w *WeeklyService) FindWeeklySummary(param repository.FindWeeklySummaryParams) (*utils.WeeklySummaryResponse, error) {
	weeklySummary, err := w.Repository.FindWeeklySummary(context.Background(), param)
	if err != nil {
		return nil, err
	}
	solution, err := w.GPTService.GetWeeklySolution(weeklySummary)
	if err != nil {
		return nil, err
	}
	response := utils.WeeklySummaryResponse{WeeklySummary: weeklySummary, Solution: *solution}
	return &response, nil
}

func (w WeeklyService) WithRepository(repo interfaces.IRepository) WeeklyService {
	w.Repository = repo
	return w
}

func (w WeeklyService) WithGPTService(service interfaces.IGPTService) WeeklyService {
	w.GPTService = service
	return w
}
