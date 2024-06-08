package gpt

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/customerror"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/utils"
	"github.com/MJU-Mobilecomputing/jjikdan-backend/pkg/interfaces"
)

type GPTService struct {
	Repository interfaces.IGPTRepository
}

func (g *GPTService) GetMenuNutrient(img string) (*utils.MenuNeutrient, error) {
	var result utils.MenuNeutrient
	resp, err := g.Repository.ChatWithImage(img, MENU_NEUTRIENT)
	if err != nil {
		return nil, customerror.InternalServerError(err)
	}
	jsonBody := utils.GetJsonString(*resp)
	err = json.Unmarshal([]byte(jsonBody), &result)
	if err != nil {
		return nil, customerror.InternalServerError(err)
	}
	log.Println(result)
	return &result, nil
}

func (g *GPTService) GetWeeklySolution(summary repository.FindWeeklySummaryRow) (*string, error) {
	resp, err := g.Repository.ChatRequest(fmt.Sprintf(WEEKLY_SOLUTION, summary.WeeklyProtein, summary.WeeklyFoodMoisture, summary.WeeklySalt, summary.WeeklyCarbon, summary.WeeklyFat))
	if err != nil {
		return nil, customerror.InternalServerError(err)
	}
	return resp, nil
}

func InitGPTService() *GPTService {
	return &GPTService{}
}

func (g GPTService) WithRepository(repo interfaces.IGPTRepository) GPTService {
	g.Repository = repo
	return g
}
