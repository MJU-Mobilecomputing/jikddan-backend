package utils

import (
	"mime/multipart"

	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type MenuNeutrient struct {
	Calrory      int    `json:"calrory"`
	FoodMoisture int    `json:"food_moisture"`
	Salt         int    `json:"salt"`
	Carbon       int    `json:"carbon"`
	Fat          int    `json:"fat"`
	Protein      int    `json:"protein"`
	Summary      string `json:"summary"`
	Score        int    `json:"score"`
}

type CreateMenuRequest struct {
	Img      multipart.FileHeader    `form:"img"`
	Date     pgtype.Date             `form:"date"`
	MenuTime repository.NullMenuTime `form:"menu_time"`
}
