package diaryday

import "github.com/jackc/pgx/v5/pgtype"

type FindDiaryDayWithMenuRequest struct {
	Date pgtype.Date `json:"date"`
}
