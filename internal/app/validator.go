package app

import (
	"github.com/MJU-Mobilecomputing/jjikdan-backend/internal/repository"
	"github.com/go-playground/validator/v10"
)

func isValidStatus(fl validator.FieldLevel) bool {
	if status, ok := fl.Field().Interface().(repository.NullStatus); ok {
		statusString := repository.Status(status.Status)
		switch statusString {
		case repository.StatusPending, repository.StatusComplete:
			return true
		}
		return false
	}
	return false
}

func isMenuTime(fl validator.FieldLevel) bool {
	if menuTime, ok := fl.Field().Interface().(repository.NullMenuTime); ok {
		menuTimeString := repository.MenuTime(menuTime.MenuTime)
		switch menuTimeString {
		case repository.MenuTimeLunch, repository.MenuTimeBreakfast, repository.MenuTimeDinner, repository.MenuTimeSnack:
			return true
		}
		return false
	}
	return false
}
