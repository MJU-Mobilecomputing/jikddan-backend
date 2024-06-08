package customerror

import "net/http"

var DIARYDAY_NOT_FOUND = NewError(http.StatusNotFound, "diary day not found", "404")

func DiaryDayNotFound(err error) error {
	diaryDayNotFound := DIARYDAY_NOT_FOUND
	diaryDayNotFound.Data = err
	return diaryDayNotFound
}
