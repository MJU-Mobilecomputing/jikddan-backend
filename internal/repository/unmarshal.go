package repository

import (
	"encoding/json"
	"errors"
)

func (s *NullStatus) UnmarshalJSON(data []byte) error {
	var statusString string
	if err := json.Unmarshal(data, &statusString); err != nil {
		return err
	}
	status := Status(statusString)
	switch status {
	case StatusPending, StatusComplete:
		*s = NullStatus{Status: status, Valid: true}
		return nil
	}

	return errors.New("invalid status value")
}

func (s *NullMenuTime) UnmarshalJSON(data []byte) error {
	var menuString string
	if err := json.Unmarshal(data, &menuString); err != nil {
		return err
	}
	menuTime := MenuTime(menuString)
	switch menuTime {
	case MenuTimeBreakfast, MenuTimeLunch, MenuTimeDinner, MenuTimeSnack:
		*s = NullMenuTime{MenuTime: menuTime, Valid: true}
		return nil
	}

	return errors.New("invalid menutime value")
}
